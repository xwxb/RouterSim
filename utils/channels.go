package interfaces

import "github.com/xwxb/routersim/consts"

var (
	HostToRouterArpChan = make(chan consts.ArpRequestPacket)
	RouterToHostArpChan = make(chan consts.ArpResponsePacket)
)
