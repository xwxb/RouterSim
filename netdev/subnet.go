package netdev

import (
	"github.com/xwxb/routersim/consts"
	"net"
)

type SubnetInfo struct {
	consts.IPAddress
	consts.SubNetMask
}

type RouteTable map[*SubnetInfo]consts.IPAddress

// 这里相当于直接套了一层底层的子网实现了。。
func (s *SubnetInfo) Contains(ipAddress consts.IPAddress) (ok bool) {
	_, subnet, _ := net.ParseCIDR(string(s.IPAddress) + "/" + string(s.SubNetMask))
	ip := net.ParseIP(string(ipAddress))
	return subnet.Contains(ip)
}
