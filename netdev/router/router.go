package router

import (
	"encoding/json"
	"github.com/xwxb/routersim/consts"
	"github.com/xwxb/routersim/netdev"
	"github.com/xwxb/routersim/utils"
	"log"
)

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
			RouteTable: map[netdev.SubnetInfo]consts.IPAddress{},
		},
	}
}

func (r *Router) Start() {
	for {
		select {
		case eFrame := <-utils.Host1ToRouterEFChan:
			log.Println("Router received ethernet frame from host1")

			var arpPacket netdev.ArpRequestPacket
			err := json.Unmarshal(eFrame.PayloadBytes, &arpPacket)
			if err != nil {
				log.Fatal(err)
				return
			}

			arpResponsePacket := r.CreateArpResponsePacket()
			utils.RouterToHost1EFChan <- arpResponsePacket
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
