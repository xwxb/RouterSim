package host

import (
	"encoding/json"
	"github.com/xwxb/routersim/consts"
	"log"
)

// 定义以太网帧结构体
type EthernetFrame struct {
	SourceMAC      consts.MACAddress
	DestinationMAC consts.MACAddress
	PayloadBytes   []byte
}

// 构造以太网帧的方法
// ? TODO 这里暂时不太明白，以太网帧这一层到底需不需要第一个入参；arp不需要，但是ipv4似乎又需要
func (h *Host) createEthernetFrame(destinationMAC consts.MACAddress, payload any) *EthernetFrame {
	b, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	return &EthernetFrame{
		SourceMAC:      h.MACAddress,
		DestinationMAC: destinationMAC,
		PayloadBytes:   b,
	}
}
