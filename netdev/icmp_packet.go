package netdev

import "github.com/xwxb/routersim/consts"

type ICMPPacket struct {
	IPv4Packet
	Type consts.ICMPPacketType
	// 省略代码、校验和、数据等字段
}

func (n *NetDeviceBase) CreateICMPPacket(destIP consts.IPAddress, typ consts.ICMPPacketType) *ICMPPacket {
	return &ICMPPacket{
		IPv4Packet: *n.CreateIPv4Packet(destIP, ""),
		Type:       typ,
	}
}
