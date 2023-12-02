package host

// 定义主机结构体
type Host struct {
	MACAddress string
	IPAddress  string
}

func NewHost(macAddress, ipAddress string) *Host {
	return &Host{
		MACAddress: macAddress,
		IPAddress:  ipAddress,
	}
}

func (h *Host) Start() {
	// 模拟主机的启动逻辑...
}

// 随机构造
func createRandPacket() string {
	return "rand packet"
}
