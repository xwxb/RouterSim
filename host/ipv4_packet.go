package host

// 定义IPv4分组结构体
type IPv4Packet struct {
	SourceIP      string
	DestinationIP string
	Payload       string
}

// 构造IPv4分组的方法
func createIPv4Packet(sourceIP, destinationIP, payload string) *IPv4Packet {
	return &IPv4Packet{
		SourceIP:      sourceIP,
		DestinationIP: destinationIP,
		Payload:       payload,
	}
}
