package host

import (
	"encoding/json"
	"github.com/xwxb/routersim/consts"
	"github.com/xwxb/routersim/netdev"
	"log"
)

func (h *Host) Receive() {
	if h.IPAddress == consts.Host2IPAddress {
		select {
		case eFrame := <-netdev.RouterToHost2EFChan:
			log.Println("Host received ethernet frame from router")
			if eFrame.PayloadType == consts.ARPType {
				log.Println("Payload type is ARP")

				var arpPacket netdev.ArpRequestPacket
				err := json.Unmarshal(eFrame.PayloadBytes, &arpPacket)
				if err != nil {
					log.Fatal(err)
					return
				}
				// 问题在这里，被h1收到了
				if arpPacket.DestIP == h.IPAddress {
					// 此时主机2发现是发给自己的，所以创建 ARP 响应报文
					log.Println("dest ip is host2-self ip, return arp response")
					arpRespPacket := h.CreateArpResponsePacket()
					frame := h.CreateEthernetFrame(consts.Host1MACAddress, arpRespPacket)
					// todo  这里ip应该是主机1，但是暂时没有内部发送的实现，所以先发给路由器
					h.SendOutEthernetFrame(frame, consts.RouterIPAddress)
				}
			} else if eFrame.PayloadType == consts.IPv4Type {
				log.Println("Payload type is IPv4")

				var ipv4Packet netdev.IPv4Packet
				err := json.Unmarshal(eFrame.PayloadBytes, &ipv4Packet)
				if err != nil {
					log.Fatal(err)
					return
				}

				// 接受到 IPv4 数据包，直接打印示意
				log.Println("Receive IPv4 Packet, src ip:", ipv4Packet.SourceIP, "payload:", ipv4Packet.Payload)
			}
		case eFrame := <-netdev.RouterToHost1EFChan:
			log.Println("Host received ethernet frame from router")
			if eFrame.PayloadType == consts.ICMPType {
				log.Println("Payload type is ICMP")

				var icmpPacket netdev.ICMPPacket
				err := json.Unmarshal(eFrame.PayloadBytes, &icmpPacket)
				if err != nil {
					log.Fatal(err)
					return
				}

				// 接受到 ICMP 数据包，如果类型是应答报文，直接打印示意
				if icmpPacket.Type == consts.ICMPTypeEchoReply {
					log.Println("Receive ICMP Reply, src ip:", icmpPacket.SourceIP)
				}
			}
		}

	}
}
