package netdev

import "github.com/xwxb/routersim/consts"

type ArpRequestPacket struct {
	// 自己的ip，mac，目标ip
	consts.IPAddress
	consts.MACAddress
	destIP consts.IPAddress
}

type ArpResponsePacket struct {
	consts.MACAddress
}

type ArpTable map[consts.IPAddress]consts.MACAddress
