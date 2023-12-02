package main

import (
	"github.com/xwxb/routersim/consts"
	"sync"

	"github.com/xwxb/routersim/host"
	"github.com/xwxb/routersim/router"
)

func main() {
	// 创建路由器
	router := router.NewRouter(consts.RouterMACAddress, consts.RouterIPAddress)

	// 创建等待组，用于等待所有goroutine完成
	var wg sync.WaitGroup

	// 启动路由器的goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		router.Start()
	}()

	// 创建主机节点A
	hostA := host.NewHost(consts.Host1MACAddress, consts.Host1IPAddress)
	wg.Add(1)
	go func() {
		defer wg.Done()
		hostA.Start()
	}()

	// 创建主机节点B
	hostB := host.NewHost(consts.Host2MACAddress, consts.Host2IPAddress)
	wg.Add(1)
	go func() {
		defer wg.Done()
		hostB.Start()
	}()

	// 等待所有goroutine完成
	wg.Wait()
}
