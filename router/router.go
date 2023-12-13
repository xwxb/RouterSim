package router

import (
	"github.com/xwxb/routersim/consts"
	"github.com/xwxb/routersim/utils"
)

type Router struct {
	Type consts.NodeType
	utils.NetDeviceBase
}

func NewRouter(macAddress, ipAddress string) *Router {
	return &Router{
		Type: consts.RouterType,
		NetDeviceBase: utils.NetDeviceBase{
			NetDeviceAddrs: utils.NetDeviceAddrs{
				IPAddress:  consts.IPAddress(ipAddress),
				MACAddress: consts.MACAddress(macAddress),
			},
			ArpTable:   make(consts.ArpTable),
			RouteTable: map[consts.SubnetInfo]consts.IPAddress{},
		},
	}
}

func (r *Router) Start() {
	for {
		select {
		case _ = <-utils.HostToRouterArpChan:
			arpResponsePacket := r.CreateArpResponsePacket()
			utils.RouterToHostArpChan <- arpResponsePacket
		}
	}
}

//func (router *Router) InsertARPTable(ipAddress, macAddress string) {
//	router.ARPTable[ipAddress] = macAddress
//}
//
//func (router *Router) InsertRoutingTable(destinationIP, nextHop string) {
//	router.RoutingTable[destinationIP] = nextHop
//}
