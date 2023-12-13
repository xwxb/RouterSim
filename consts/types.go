package consts

type MACAddress string
type IPAddress string

// SubNetMask 这里暂时直接使用子网掩码有多少一的整数字符串，对应CIDR表示法的子网掩码
// 真正要底层实现就是二进制了，干脆用最抽象的
type SubNetMask string

type NodeType string

const (
	HostType   = "host"
	RouterType = "router"
)

type RouteTable map[SubnetInfo]IPAddress
