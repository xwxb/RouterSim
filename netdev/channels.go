package netdev

import (
	"github.com/xwxb/routersim/consts"
	"log"
)

// 感觉有可能可以考虑合并成两个，就是路由器和路由器出，因为路由器理论上不知道是谁给他的。但是要考虑清楚并发的问题
// 上面这个想法保留，感觉没想清楚但是好像用得上
// 目前的核心思路是，MAC 和 IP 都可以用来唯一确定一个设备了，所以似乎根本就用不上这个 channel 了。但是这样太不解耦了，不考虑
// 所以暂时的思路还是由出入参来直接唯一获得两个 channel
var (
	// host1
	Host1ToRouterEFChan = make(chan *EthernetFrame, 3)
	RouterToHost1EFChan = make(chan *EthernetFrame, 3)

	// host2
	Host2ToRouterEFChan = make(chan *EthernetFrame, 3)
	RouterToHost2EFChan = make(chan *EthernetFrame, 3)
)

type AddrPair struct {
	Src, Dst any
}
type apToEFChanMap map[AddrPair]chan *EthernetFrame

var dirMap apToEFChanMap

func init() {
	// host1 and router
	regNagMap(consts.Host1IPAddress, consts.RouterIPAddress, Host1ToRouterEFChan)
	regNagMap(consts.RouterIPAddress, consts.Host1IPAddress, RouterToHost1EFChan)
	// host2 and router
	regNagMap(consts.Host2IPAddress, consts.RouterIPAddress, Host2ToRouterEFChan)
	regNagMap(consts.RouterIPAddress, consts.Host2IPAddress, RouterToHost2EFChan)

	log.Println("net channels init complete")
}

// 这里的设计思想还是 map 和 ip 在物理世界可以唯一确定一条路线，这里就模拟实现
// 实际目前这里只通过 IP 确定就行了，这里是考虑扩展性的设计
func regNagMap(from, to any, ch chan *EthernetFrame) {
	ipTour := AddrPair{from, to}
	if dirMap == nil {
		dirMap = make(apToEFChanMap)
	}
	dirMap[ipTour] = ch
}

func GetDirChan(from, to any) (ch chan *EthernetFrame) {
	ipTour := AddrPair{from, to}
	ch, ok := dirMap[ipTour]
	if !ok {
		log.Fatal("No such route ", "from ", from, " to ", to)
		return
	}
	return
}
