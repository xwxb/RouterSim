package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Router maintains routing table and participates in routing updates
type Router struct {
	Id        int
	Routes    []Route
	Neighbors []int // router ids
	Messages  chan LinkStateMessage
	Shutdown  chan struct{}
}

// Route represents a route to destination and next hop
type Route struct {
	Destination int
	NextHop     int
	Cost        int
}

// Flood sends route update to all neighbors
func (r *Router) Flood(msg LinkStateMessage) {
	for _, neighbor := range r.Neighbors {
		go func(n int) {
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) // simulate delay
			routers[n].Messages <- msg
		}(neighbor)
	}
}

// HandleMessages processes incoming route updates
func (r *Router) HandleMessages() {
	for {
		select {
		case msg := <-r.Messages:
			r.ProcessMessage(msg)
		case <-r.Shutdown:
			close(r.Messages)
			return
		}
	}
}

// ProcessMessage handles incoming route update
func (r *Router) ProcessMessage(msg LinkStateMessage) {
	fmt.Printf("Router %d received update from %d\n", r.Id, msg.FromRouter)

	// Update routing table based on new link states
	r.UpdateRoutes(msg.LinkStates)

	// Flood update to neighbors
	r.Flood(msg)
}

// UpdateRoutes recalculates routes based on Dijkstra's algorithm
func (r *Router) UpdateRoutes(updates []LinkState) {
	//// Initialize the priority queue with the updates
	//pq := make(PriorityQueue, len(updates))
	//for i, update := range updates {
	//	pq[i] = &update
	//}
	//heap.Init(&pq)
	//
	//// Initialize the routes map
	//r.Routes = make(map[int]int)
	//for _, update := range updates {
	//	r.Routes[update.NeighborRouter] = math.MaxInt32
	//}
	//
	//// Dijkstra's algorithm
	//for pq.Len() > 0 {
	//	u := heap.Pop(&pq).(*LinkState)
	//	for _, update := range updates {
	//		if update.NeighborRouter == u.NeighborRouter {
	//			alt := r.Routes[u.NeighborRouter] + update.Cost
	//			if alt < r.Routes[update.NeighborRouter] {
	//				r.Routes[update.NeighborRouter] = alt
	//				heap.Push(&pq, &LinkState{update.NeighborRouter, alt})
	//			}
	//		}
	//	}
	//}
}
