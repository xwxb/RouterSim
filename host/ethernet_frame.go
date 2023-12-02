package host

// 定义以太网帧结构体
type EthernetFrame struct {
	SourceMAC      string
	DestinationMAC string
	Payload        string
}

// 构造以太网帧的方法
func createEthernetFrame(sourceMAC, destinationMAC, payload string) *EthernetFrame {
	return &EthernetFrame{
		SourceMAC:      sourceMAC,
		DestinationMAC: destinationMAC,
		Payload:        payload,
	}
}
