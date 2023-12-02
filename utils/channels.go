package utils

import "github.com/xwxb/routersim/consts"

func Init() {

}

var (
	HostToRouterArpChan = make(chan consts.ArpRequestPacket)
	RouterToHostArpChan = make(chan consts.ArpResponsePacket)
)
