package netdev

import (
	"github.com/xwxb/routersim/consts"
)

// 定义IPv4分组结构体
type IPv4Packet struct {
	SourceIP      consts.IPAddress
	DestinationIP consts.IPAddress
	Payload       string
}

// 构造IPv4分组的方法
func (n *NetDeviceBase) CreateIPv4Packet(destinationIP consts.IPAddress, payload string) *IPv4Packet {
	return &IPv4Packet{
		SourceIP:      n.IPAddress,
		DestinationIP: destinationIP,
		Payload:       payload,
	}
}
