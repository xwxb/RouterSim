package host

import (
	"encoding/json"
	"github.com/xwxb/routersim/consts"
	"github.com/xwxb/routersim/netdev"
	"log"
)

// 定义主机结构体
type Host struct {
	Type                 consts.NodeType // TODO 这个应该抽到 netdev 层
	netdev.NetDeviceBase                 // 路由表
}

func NewHost(macAddress, ipAddress string) *Host {
	return &Host{
		Type: consts.HostType,
		NetDeviceBase: netdev.NetDeviceBase{
			NetDeviceAddrs: netdev.NetDeviceAddrs{
				IPAddress:  consts.IPAddress(ipAddress),
				MACAddress: consts.MACAddress(macAddress),
			},
			ArpTable:   make(netdev.ArpTable),
			RouteTable: map[*netdev.SubnetInfo]consts.IPAddress{},
		},
	}
}

func (h *Host) Start() {
	// 首先检查自己 ARP 缓存里面是否有目标主机的 MAC 地址
	log.Println("Host Start")

	// Start 这里应该主要是做一些循环监听式的事情
	// TODO 重构成 delayTask 式的
	for range consts.HostRcvTickerChan {
		h.Receive()
	}
}

func (h *Host) Receive() {
	if h.IPAddress == consts.Host2IPAddress {
		select {
		case eFrame := <-netdev.RouterToHost2EFChan:
			log.Println("Host received ethernet frame from router")
			if eFrame.PayloadType == consts.ARPType {
				log.Println("Payload type is ARP")

				var arpPacket netdev.ArpRequestPacket
				err := json.Unmarshal(eFrame.PayloadBytes, &arpPacket)
				if err != nil {
					log.Fatal(err)
					return
				}
				// 问题在这里，被h1收到了
				if arpPacket.DestIP == h.IPAddress {
					// 此时主机2发现是发给自己的，所以创建 ARP 响应报文
					log.Println("dest ip is host2-self ip, return arp response")
					arpRespPacket := h.CreateArpResponsePacket()
					frame := h.CreateEthernetFrame(consts.Host1MACAddress, arpRespPacket)
					// todo  这里ip应该是主机1，但是暂时没有内部发送的实现，所以先发给路由器
					h.SendOutEthernetFrame(frame, consts.RouterIPAddress)
				}
			} else if eFrame.PayloadType == consts.IPv4Type {
				log.Println("Payload type is IPv4")

				var ipv4Packet netdev.IPv4Packet
				err := json.Unmarshal(eFrame.PayloadBytes, &ipv4Packet)
				if err != nil {
					log.Fatal(err)
					return
				}

				// 接受到 IPv4 数据包，直接打印示意
				log.Println("Receive IPv4 Packet, src ip:", ipv4Packet.SourceIP, "payload:", ipv4Packet.Payload)
			}

		}

	}
}

// 获取 ARP
// hotTODO 应该是私有的，公有测试用
func (h *Host) getArp(destIP consts.IPAddress, destMAC *consts.MACAddress) {
	log.Println("No ARP Cache, Start to Get ARP Resp")

	// 通过 ARP 协议获取目标主机的 MAC 地址
	arpRequestPacket := h.CreateARPRequestPacket(destIP)
	frame := h.CreateEthernetFrame(consts.RouterMACAddress, arpRequestPacket)
	h.sendToRouter(frame)

	// 接收 ARP 响应报文
	// TODO 目前这里暂时这样写，感觉这里好像一定程度上是异步的，应该在监听处
	resp := <-netdev.RouterToHost1EFChan
	respMAC := resp.DestinationMAC
	log.Println("Get ARP Resp, MAC Address:", respMAC)

	// 更新自己的 ARP 缓存
	log.Println("Update ARP Cache using", consts.Host2IPAddress, "and", respMAC, "to update ARP Cache")
	if h.ArpTable == nil {
		h.ArpTable = make(netdev.ArpTable)
	}
	h.ArpTable[consts.Host2IPAddress] = resp.DestinationMAC

	// 更新目标主机的 MAC 地址
	destMAC = &respMAC
}

// 发送ipv4数据包
func (h *Host) SendIPv4Packet(ipAddress consts.IPAddress, payload string) {
	log.Println("Start to Send IPv4 Packet, dest ip:", ipAddress, "payload:", payload)

	log.Println("Start to Find ARP Cache, curr table: ", h.ArpTable)
	destMAC, ok := h.ArpTable[ipAddress]
	if !ok {
		// 如果没有，则通过 ARP 协议获取目标主机的 MAC 地址
		h.getArp(consts.Host2IPAddress, &destMAC)
	}

	packet := h.CreateIPv4Packet(consts.Host2IPAddress, payload)
	frame := h.CreateEthernetFrame(consts.Host2MACAddress, packet)
	h.SendOutEthernetFrame(frame, consts.RouterIPAddress)
}

// 目前就考虑这种设计模式，确定主语，然后是把 channel 的实现尽量封装
// TODO 这里不应该特化，清楚路由器和广播的关系; 这里理论上应该是一个 SendInner 的抽象 ，现在懒得实现了
func (h *Host) sendToRouter(frame *netdev.EthernetFrame) {
	netdev.Host1ToRouterEFChan <- frame
}

//
