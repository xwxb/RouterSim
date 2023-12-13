package netdev

import "github.com/xwxb/routersim/consts"

type ArpRequestPacket struct {
	consts.MACAddress
}

type ArpResponsePacket struct {
	consts.MACAddress
}

type ArpTable map[consts.IPAddress]consts.MACAddress
