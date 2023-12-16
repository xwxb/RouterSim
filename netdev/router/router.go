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

// TODO 考虑抽象到 netdev 接口层
func (r *Router) Start() {
	for {
		select {
		case eFrame := <-utils.Host1ToRouterEFChan:
			if eFrame.PayloadType == consts.ARPType {
				log.Println("Router received ethernet frame from host1")

				var arpPacket netdev.ArpRequestPacket
				err := json.Unmarshal(eFrame.PayloadBytes, &arpPacket)
				if err != nil {
					log.Fatal(err)
					return
				}

				if arpPacket.DestIP == r.IPAddress {
					// 这里路由器发现不是发给自己的，继续广播。广播此处不做实现，直接发给主机2
					log.Println("dest ip is not router ip, continue broadcast")
					inChan, _ := utils.GetInAndOutChan(consts.RouterIPAddress, consts.Host2IPAddress)
					inChan <- eFrame
				}
			}
		default:
			log.Fatal("Router received unknown ethernet frame")
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
