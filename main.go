package main

import (
	"sync"

	"github.com/xwxb/routersim/host"
	"github.com/xwxb/routersim/router"
)

func main() {
	// 创建路由器
	router := router.NewRouter("AA:AA:AA:AA:AA:AA", "192.168.1.1")

	// 创建等待组，用于等待所有goroutine完成
	var wg sync.WaitGroup

	// 启动路由器的goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		routerStart(router)
	}()

	// 创建主机节点A
	hostA := host.NewHost("BB:BB:BB:BB:BB:BB", "192.168.1.2")
	wg.Add(1)
	go func() {
		defer wg.Done()
		hostStart(hostA, router)
	}()

	// 创建主机节点B
	hostB := host.NewHost("CC:CC:CC:CC:CC:CC", "192.168.1.3")
	wg.Add(1)
	go func() {
		defer wg.Done()
		hostStart(hostB, router)
	}()

	// 等待所有goroutine完成
	wg.Wait()
}

// 路由器启动逻辑
func routerStart(router *Router) {
	// 模拟路由器的其他启动逻辑...
}
