package host

import (
	"github.com/xwxb/routersim/consts"
	"github.com/xwxb/routersim/netdev"
	"github.com/xwxb/routersim/utils"
	"log"
	"time"
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
	log.Println("Start to Find ARP Cache, curr table: ", h.ArpTable)
	destMAC, ok := h.ArpTable[consts.Host2IPAddress]

	if !ok {
		// 如果没有，则通过 ARP 协议获取目标主机的 MAC 地址
		log.Println("No ARP Cache, Start to Get ARP Resp")

		// 通过 ARP 协议获取目标主机的 MAC 地址
		arpRequestPacket := h.CreateARPRequestPacket() // 构造 ARP 请求报文
		h.broadcastARPRequestPacket(arpRequestPacket)  // 广播 ARP 请求报文

		// 接收 ARP 响应报文
		resp := <-utils.RouterToHostArpChan
		respMAC := resp.MACAddress
		log.Println("Get ARP Resp, MAC Address:", respMAC)

		// 更新自己的 ARP 缓存
		log.Println("Update ARP Cache using", consts.Host2IPAddress, "and", respMAC, "to update ARP Cache")
		if h.ArpTable == nil {
			h.ArpTable = make(netdev.ArpTable)
		}
		h.ArpTable[consts.Host2IPAddress] = resp.MACAddress

		// 更新目标主机的 MAC 地址
		destMAC = respMAC
	}

	// 循环每隔 10s 构造、封装、发送一个随机的 IPv4 数据包
	for range time.Tick(10 * time.Second) {
		// 首先随机生成一个字符串作为 payload
		pl := utils.GetRandStr(10)
		packet := h.createIPv4Packet(consts.Host2IPAddress, pl)
		frame := h.createEthernetFrame(consts.Host2MACAddress, packet.String())
	}

}

// 广播 ARP 请求报文
func (h *Host) broadcastARPRequestPacket(arpRequestPacket netdev.ArpRequestPacket) {
	// 广播 ARP 请求报文，这里只是模拟，实际上是通过以太网发送
	// 通过 channel 发送请求到其他所有 goroutine，如果节点类型是 Router 则往回发送 ARP 响应报文，包含自己的 MAC 地址
	utils.HostToRouterArpChan <- arpRequestPacket
}
