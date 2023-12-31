package router

import (
	"encoding/json"
	"github.com/xwxb/routersim/consts"
	"github.com/xwxb/routersim/netdev"
	"log"
)

func (r *Router) Receive() {
	select {
	case eFrame := <-netdev.Host1ToRouterEFChan:
		log.Println("Router received LAN ethernet frame from host1")
		// TODO else if 换成接口实现
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
				ch := netdev.GetDirChan(consts.RouterIPAddress, consts.Host2IPAddress)
				ch <- eFrame
			}
		} else if eFrame.PayloadType == consts.IPv4Type {
			log.Println("Payload type is IPv4")

			var ipv4Packet netdev.IPv4Packet
			err := json.Unmarshal(eFrame.PayloadBytes, &ipv4Packet)
			if err != nil {
				log.Fatal(err)
				return
			}

			// 从路由表中尝试获取下一跳
			ok := netdev.Host1SubnetInfo.Contains(ipv4Packet.DestinationIP)
			if !ok {
				// 没有找到下一跳，转发到默认网关，这里不做实现
				log.Println("Can't find next hop, go to default gateway route")
				return
			}
			v, _ := r.RouteTable[netdev.Host1SubnetInfo]
			r.SendOutEthernetFrame(eFrame, v)
		} else if eFrame.PayloadType == consts.ICMPType {
			log.Println("Payload type is ICMP")

			var icmpPacket netdev.ICMPPacket
			err := json.Unmarshal(eFrame.PayloadBytes, &icmpPacket)
			if err != nil {
				log.Fatal(err)
				return
			}

			// 如果是发给自己的，那就发出应答报文返回
			if icmpPacket.DestinationIP == r.IPAddress {
				log.Println("dest ip is router ip, return icmp response")
				icmpRespPacket := r.CreateICMPPacket(consts.Host1IPAddress, consts.ICMPTypeEchoReply)
				frame := r.CreateEthernetFrame(consts.Host1MACAddress, icmpRespPacket)
				r.SendOutEthernetFrame(frame, consts.Host1IPAddress)
			} else {
				// 如果不是发给自己的，那就转发
				log.Println("dest ip is not router ip, continue forward")
				// 暂时不维护
				//r.SendOutEthernetFrame(eFrame, icmpPacket.DestinationIP)
			}
		}
	case eFrame := <-netdev.Host2ToRouterEFChan:
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
				ch := netdev.GetDirChan(consts.RouterIPAddress, consts.Host1IPAddress)
				ch <- eFrame
			}
		}
	default:
		log.Println("Router received nothing")
	}
}
