package host

import "github.com/xwxb/routersim/consts"

// 定义IPv4分组结构体
type IPv4Packet struct {
	SourceIP      consts.IPAddress
	DestinationIP consts.IPAddress
	Payload       string
}

// 构造IPv4分组的方法
func (h *Host) createIPv4Packet(destinationIP consts.IPAddress, payload string) *IPv4Packet {
	return &IPv4Packet{
		SourceIP:      h.IPAddress,
		DestinationIP: destinationIP,
		Payload:       payload,
	}
}

// 暂时先用一个其实没实际意义的字符串化吧，TODO 后面关注一下具体干了什么
func (p *IPv4Packet) String() string {
	return string(p.SourceIP) + " " + string(p.DestinationIP) + " " + p.Payload
}
