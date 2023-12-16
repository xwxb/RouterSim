package netdev

import (
	"github.com/xwxb/routersim/consts"
	"github.com/xwxb/routersim/netdev/host"
	"github.com/xwxb/routersim/utils"
)

type NetDevice interface {
	GetNextHop(consts.IPAddress) consts.IPAddress
	CreateArpResponsePacket() ArpResponsePacket
	SendOutEthernetFrame(ef *host.EthernetFrame, ip consts.IPAddress)
}

// 这个要重构到 addrs.go 中，暂时没找到最佳实践
type NetDeviceAddrs struct {
	IPAddress  consts.IPAddress
	MACAddress consts.MACAddress
}

type NetDeviceBase struct {
	NetDeviceAddrs
	ArpTable   ArpTable
	RouteTable consts.RouteTable
}

func (n *NetDeviceBase) GetNextHop(ipAddress consts.IPAddress) (next consts.IPAddress) {
	// 正常路由器的匹配规则，首先遍历 RouteTable 中所有 key（子网），然后匹配传输的参数 ipAddress 是否在 子网内，如果在，则返回 value（下一跳）
	for subnetInfo, nextHop := range n.RouteTable {
		if subnetInfo.Contains(ipAddress) {
			next = nextHop
			return
		}
	}
	return
}

func (n *NetDeviceBase) CreateArpResponsePacket() ArpResponsePacket {
	return ArpResponsePacket{
		MACAddress: n.MACAddress,
	}
}

// 构造 ARP 请求报文
func (n *NetDeviceBase) CreateARPRequestPacket(destIPAddress consts.IPAddress) ArpRequestPacket {
	return ArpRequestPacket{
		IPAddress:  n.IPAddress,
		MACAddress: n.MACAddress,
		DestIP:     destIPAddress,
	}
}

func (n *NetDeviceBase) SendOutEthernetFrame(ef *host.EthernetFrame, ip consts.IPAddress) {
	ch := utils.GetDirChan(n.IPAddress, ip)
	ch <- ef
}
