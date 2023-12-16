package router

import (
	"encoding/json"
	"github.com/xwxb/routersim/consts"
	"github.com/xwxb/routersim/netdev"
	"github.com/xwxb/routersim/utils"
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
			RouteTable: map[netdev.SubnetInfo]consts.IPAddress{},
		},
	}
}

// TODO 考虑抽象到 netdev 接口层
func (r *Router) Start() {
	for {
		select {
		case eFrame := <-utils.Host1ToRouterEFChan:
			log.Println("Router received LAN ethernet frame from host1")
			if eFrame.PayloadType == consts.ARPType {
				log.Println("Payload type is ARP")

				var arpPacket netdev.ArpRequestPacket
				err := json.Unmarshal(eFrame.PayloadBytes, &arpPacket)
				if err != nil {
					log.Fatal(err)
					return
				}

				if arpPacket.DestIP != r.IPAddress {
					// 这里路由器发现不是发给自己的，继续用原包广播。广播此处不做实现，直接发给主机2
					log.Println("dest ip is not router ip, continue broadcast")
					ch := utils.GetDirChan(consts.RouterIPAddress, consts.Host2IPAddress)
					ch <- eFrame
				}
			}
		case eFrame := <-utils.Host2ToRouterEFChan:
			log.Println("Router received external ethernet frame from host2")
			if eFrame.PayloadType == consts.ARPType {
				log.Println("Payload type is ARP")

				var arpPacket netdev.ArpRequestPacket
				err := json.Unmarshal(eFrame.PayloadBytes, &arpPacket)
				if err != nil {
					log.Fatal(err)
					return
				}

				if arpPacket.DestIP != r.IPAddress {
					// 这里路由器发现不是发给自己的，而且已经标明 MAC 地址是发给主机1的，所以直接发给主机1
					// TODO 这里也是内部发送不做实现，暂时先这样
					log.Println("dest ip is not router ip, continue forward")
					ch := utils.GetDirChan(consts.RouterIPAddress, consts.Host1IPAddress)
					ch <- eFrame
				}
			}
		default:
			log.Fatal("Router received unknown ethernet frame")
		}
	}
}
