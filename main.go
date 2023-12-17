package main

import (
	"fmt"
	"github.com/xwxb/routersim/consts"
	"github.com/xwxb/routersim/netdev/host"
	"github.com/xwxb/routersim/netdev/router"
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

	var destMAC consts.MACAddress
	host1.GetArp(consts.Host2IPAddress, &destMAC)
	fmt.Println(destMAC)

	// 等待所有goroutine完成
	wg.Wait()
}
