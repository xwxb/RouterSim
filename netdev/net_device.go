package netdev

import "github.com/xwxb/routersim/consts"

type NetDevice interface {
	GetNextHop(consts.IPAddress) consts.IPAddress
	CreateArpResponsePacket() ArpResponsePacket
}

type NetDeviceAddrs struct {
	IPAddress  consts.IPAddress
	MACAddress consts.MACAddress
}

type NetDeviceBase struct {
	NetDeviceAddrs
	ArpTable   ArpTable
	RouteTable consts.RouteTable
}

func (n *NetDeviceBase) GetNextHop(ipAddress consts.IPAddress) (next consts.IPAddress, err error) {
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
func (n *NetDeviceBase) CreateARPRequestPacket() ArpRequestPacket {
	return ArpRequestPacket{
		IPAddress:  n.IPAddress,
		MACAddress: n.MACAddress,
		DestMAC:    consts.BroadcastMACAddress,
	}
}
