package main

import (
	"fmt"
	"time"
)

// global simulating
var routers []*Router

func initializeNetwork() {
	routers = make([]*Router, 3)

	for i := 0; i < 3; i++ {
		r := &Router{
			Id:        i + 1,
			Routes:    make([]Route, 0),
			Neighbors: []int{},
			Messages:  make(chan LinkStateMessage),
			Shutdown:  make(chan struct{}),
		}
		routers[i] = r
	}

	routers[0].Neighbors = []int{2}
	routers[1].Neighbors = []int{2}
	routers[2].Neighbors = []int{0, 1}
}

func main() {
	initializeNetwork()

	// start route update message handlers
	for _, router := range routers {
		go router.HandleMessages()
	}

	// initialize routing with empty link states
	for _, router := range routers {
		router.UpdateRoutes([]LinkState{})
	}

	// print initial routing tables
	printRoutingTables()

	// simulate link failure after 5s
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Link failure between router 0 and 2")

		// dynamic update
		routers[0].Neighbors = []int{1}
		routers[2].Neighbors = []int{1}
	}()

	// send initial route updates
	for i, router := range routers {
		router.Flood(LinkStateMessage{
			FromRouter: i + 1,
			LinkStates: linkStatesForRouter(router),
		})
	}

	// shutdown after 10s
	time.AfterFunc(10*time.Second, func() {
		for _, r := range routers {
			close(r.Shutdown)
		}
	})

	// wait for shutdown
	for _, r := range routers {
		<-r.Shutdown
	}
}

// printRoutingTables prints routing table for each router
func printRoutingTables() {
	fmt.Println("Routing tables:")
	for i, router := range routers {
		fmt.Printf(" Router %d\n", i+1)
		for _, route := range router.Routes {
			fmt.Printf("  %d via %d, cost %d\n",
				route.Destination, route.NextHop, route.Cost)
		}
		fmt.Println()
	}
}
