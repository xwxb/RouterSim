package consts

type MACAddress string
type IPAddress string

//type NetDeviceAddrs struct {
//	IPAddress  IPAddress
//	MACAddress MACAddress
//}

// SubNetMask 这里暂时直接使用子网掩码有多少一的整数字符串，对应CIDR表示法的子网掩码
// 真正要底层实现就是二进制了，干脆用最抽象的
type SubNetMask string

// 节点类型枚举
type NodeType string

const (
	HostType   = "host"
	RouterType = "router"
)

// 网络层协议类型枚举
type NetworkProtocolType string

const (
	ARPType  = "arp"
	IPv4Type = "ipv4"
	// 所以枚举类型命名的最佳实践应该是有前缀的，因为它没有go内置的限制。同时，枚举类型名也倾向于具体一点
	ICMPType = "icmp"
)

type ICMPPacketType string

const (
	ICMPTypeEchoReply   ICMPPacketType = "echo-reply"
	ICMPTypeEchoRequest ICMPPacketType = "echo-request"
)
