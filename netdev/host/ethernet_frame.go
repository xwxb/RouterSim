package host

import "github.com/xwxb/routersim/consts"

// 定义以太网帧结构体
type EthernetFrame struct {
	SourceMAC      consts.MACAddress
	DestinationMAC consts.MACAddress
	Payload        string
}

// 构造以太网帧的方法
func (h *Host) createEthernetFrame(destinationMAC consts.MACAddress, payload string) *EthernetFrame {
	return &EthernetFrame{
		SourceMAC:      h.MACAddress,
		DestinationMAC: destinationMAC,
		Payload:        payload,
	}
}
