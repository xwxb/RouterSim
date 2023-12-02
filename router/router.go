package router

type Router struct {
	MACAddress   string
	IPAddress    string
	ARPTable     map[string]string // ARP缓存表
	RoutingTable map[string]string // 路由表
}

func NewRouter(macAddress, ipAddress string) *Router {
	return &Router{
		MACAddress:   macAddress,
		IPAddress:    ipAddress,
		ARPTable:     make(map[string]string),
		RoutingTable: make(map[string]string),
	}
}

func (router *Router) InsertARPTable(ipAddress, macAddress string) {
	router.ARPTable[ipAddress] = macAddress
}

func (router *Router) InsertRoutingTable(destinationIP, nextHop string) {
	router.RoutingTable[destinationIP] = nextHop
}
