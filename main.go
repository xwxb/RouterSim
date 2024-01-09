package main

import (
	"github.com/xwxb/routersim/consts"
	"github.com/xwxb/routersim/netdev"
	"github.com/xwxb/routersim/netdev/host"
	"github.com/xwxb/routersim/netdev/router"
	"github.com/xwxb/routersim/utils"
	"sync"
)

func main() {
	// 创建路由器
	rt := router.NewRouter(consts.RouterMACAddress, consts.RouterIPAddress)

	// 创建等待组，用于等待所有goroutine完成
	var wg sync.WaitGroup

	// 启动路由器的goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		rt.Start()
	}()

	// 创建主机节点1
	host1 := host.NewHost(consts.Host1MACAddress, consts.Host1IPAddress)
	wg.Add(1)
	go func() {
		defer wg.Done()
		host1.Start()
	}()

	// 创建主机节点2
	host2 := host.NewHost(consts.Host2MACAddress, consts.Host2IPAddress)
	wg.Add(1)
	go func() {
		defer wg.Done()
		host2.Start()
	}()

	rt.ConfigRouteTable(netdev.Host1SubnetInfo, consts.Host2IPAddress)
	host1.SendIPv4Packet(consts.Host2IPAddress, utils.GetRandStr(10))
	host1.SendICMPRequestPacket(consts.Host2IPAddress, consts.ICMPTypeEchoRequest)

	// 等待所有goroutine完成
	wg.Wait()
}
