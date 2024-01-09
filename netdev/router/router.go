package router

import (
	"github.com/xwxb/routersim/consts"
	"github.com/xwxb/routersim/netdev"
	"log"
)

// TODO 单例重构
type Router struct {
	Type consts.NodeType
	netdev.NetDeviceBase
}

func NewRouter(macAddress, ipAddress string) *Router {
	return &Router{
		Type: consts.RouterType,
		NetDeviceBase: netdev.NetDeviceBase{
			NetDeviceAddrs: netdev.NetDeviceAddrs{
				IPAddress:  consts.IPAddress(ipAddress),
				MACAddress: consts.MACAddress(macAddress),
			},
			ArpTable:   make(netdev.ArpTable),
			RouteTable: map[*netdev.SubnetInfo]consts.IPAddress{},
		},
	}
}

// TODO 考虑抽象到 netdev 接口层
func (r *Router) Start() {
	log.Println("Router Start")

	for range consts.RouterRcvTickerChan {
		r.Receive()
	}
}

// 手动配置路由表
func (r *Router) ConfigRouteTable(subnetInfo *netdev.SubnetInfo, nextHop consts.IPAddress) {
	r.RouteTable[subnetInfo] = nextHop
}
