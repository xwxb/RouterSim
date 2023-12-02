package router

import (
	"encoding/hex"
	"fmt"
)

// 处理以太网帧的方法
func (router *Router) processEthernetFrame(frame *EthernetFrame) {
	// 解析以太网帧中的IPv4分组
	ipv4PacketBytes, err := hex.DecodeString(frame.Payload)
	if err != nil {
		fmt.Println("Error decoding IPv4 packet:", err)
		return
	}

	// 构造IPv4分组
	ipv4Packet := &IPv4Packet{}
	if err := ipv4Packet.FromBytes(ipv4PacketBytes); err != nil {
		fmt.Println("Error creating IPv4 packet:", err)
		return
	}

	// 输出日志
	fmt.Printf("Router received Ethernet Frame: %+v\n", frame)
	fmt.Printf("Decoded IPv4 Packet: %+v\n", ipv4Packet)

	// 根据IPv4分组的目标IP地址进行处理
	router.processIPv4Packet(ipv4Packet)
}

// 处理ARP分组的方法
func (router *Router) processARPPacket(arpPacket *ARPPacket) {
	// 更新ARP缓存表
	router.ARPTable[arpPacket.SenderIP] = arpPacket.SenderMAC

	// 输出日志
	fmt.Printf("Router received ARP Packet: %+v\n", arpPacket)
}

// 处理IPv4分组的方法
func (router *Router) processIPv4Packet(ipv4Packet *IPv4Packet) {
	// 根据目标IP地址查找路由表
	nextHop, found := router.RoutingTable[ipv4Packet.DestinationIP]
	if !found {
		fmt.Printf("Router could not find a route for destination IP: %s\n", ipv4Packet.DestinationIP)
		return
	}

	// 查找下一跳的MAC地址
	nextHopMAC, found := router.ARPTable[nextHop]
	if !found {
		fmt.Printf("Router could not find MAC address for next hop: %s\n", nextHop)
		return
	}

	// 构造以太网帧
	ethernetFrame := createEthernetFrame(router.MACAddress, nextHopMAC, ipv4Packet.ToBytes())

	// 输出日志
	fmt.Printf("Router sending Ethernet Frame: %+v\n", ethernetFrame)
}

// 处理ICMP分组的方法
func (router *Router) processICMPPacket(icmpPacket *ICMPPacket) {
	// 处理ICMP分组的逻辑...

	// 输出日志
	fmt.Printf("Router received ICMP Packet: %+v\n", icmpPacket)
}
