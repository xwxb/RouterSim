package host

import (
	"encoding/json"
	"github.com/xwxb/routersim/consts"
	"github.com/xwxb/routersim/netdev"
	"log"
)

// 定义以太网帧结构体
type EthernetFrame struct {
	SourceMAC      consts.MACAddress
	DestinationMAC consts.MACAddress
	// ? 我理解应该实际实现应该式有一个这样的标志位的，有时间看看，暂时这样实现
	PayloadType  consts.NetworkProtocolType
	PayloadBytes []byte
}

// 构造以太网帧的方法
// ? TODO 这里暂时不太明白，以太网帧这一层到底需不需要第一个入参；arp不需要，但是ipv4似乎又需要
func (h *Host) createEthernetFrame(destinationMAC consts.MACAddress, payload any) *EthernetFrame {
	b, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	var typ consts.NetworkProtocolType
	switch payload.(type) {
	case netdev.ArpRequestPacket:
		typ = consts.ARPType
	case IPv4Packet:
		typ = consts.IPv4Type
	}

	return &EthernetFrame{
		SourceMAC:      h.MACAddress,
		DestinationMAC: destinationMAC,
		PayloadType:    typ,
		PayloadBytes:   b,
	}
}
