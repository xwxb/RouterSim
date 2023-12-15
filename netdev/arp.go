package netdev

import "github.com/xwxb/routersim/consts"

type ArpRequestPacket struct {
	// 自己的ip，mac，目标mac
	consts.IPAddress
	consts.MACAddress
	DestMAC consts.MACAddress
}

type ArpResponsePacket struct {
	consts.MACAddress
}

type ArpTable map[consts.IPAddress]consts.MACAddress
