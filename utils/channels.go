package utils

import (
	"github.com/xwxb/routersim/netdev"
)

func Init() {

}

var (
	HostToRouterArpChan = make(chan netdev.ArpRequestPacket)
	RouterToHostArpChan = make(chan netdev.ArpResponsePacket)
)
