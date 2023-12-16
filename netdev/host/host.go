package host

import (
	"github.com/xwxb/routersim/consts"
	"github.com/xwxb/routersim/netdev"
	"github.com/xwxb/routersim/utils"
	"log"
)

// 定义主机结构体
type Host struct {
	Type                 consts.NodeType
	netdev.NetDeviceBase // 路由表
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
			RouteTable: map[netdev.SubnetInfo]consts.IPAddress{},
		},
	}
}

func (h *Host) Start() {
	// 首先检查自己 ARP 缓存里面是否有目标主机的 MAC 地址
	log.Println("Host Start")

	// TODO 没想好要干嘛，不要强行做一些事情破坏设计
}

// 获取 ARP
func (h *Host) getArp(destIP consts.IPAddress, destMAC *consts.MACAddress) {
	log.Println("No ARP Cache, Start to Get ARP Resp")

	// 通过 ARP 协议获取目标主机的 MAC 地址
	arpRequestPacket := h.CreateARPRequestPacket(destIP)
	frame := h.createEthernetFrame(consts.BroadcastMACAddress, arpRequestPacket)
	h.sendToRouter(frame)

	// 接收 ARP 响应报文
	resp := <-utils.RouterToHost1EFChan
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
func (h *Host) sendIPv4Packet(ipAddress consts.IPAddress, payload string) {
	log.Println("Start to Find ARP Cache, curr table: ", h.ArpTable)
	destMAC, ok := h.ArpTable[consts.Host2IPAddress]
	if !ok {
		// 如果没有，则通过 ARP 协议获取目标主机的 MAC 地址
		h.getArp(consts.Host2IPAddress, &destMAC)
	}

	// 首先随机生成一个字符串作为 payload
	pl := utils.GetRandStr(10)
	packet := h.createIPv4Packet(consts.Host2IPAddress, pl)
	frame := h.createEthernetFrame(consts.Host2MACAddress, packet)
}

// 目前就考虑这种设计模式，确定主语，然后是把 channel 的实现尽量封装
func (h *Host) sendToRouter(frame *EthernetFrame) {
	utils.Host1ToRouterEFChan <- frame
}

//
