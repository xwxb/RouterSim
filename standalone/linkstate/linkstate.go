package main

import "math/rand"

// LinkStateMessage is sent between routers to communicate current links
type LinkStateMessage struct {
	FromRouter int
	LinkStates []LinkState
}

// LinkState represents a link from router to neighbor and cost
type LinkState struct {
	NeighborRouter int
	Cost           int
}

// linkStatesForRouter generates link state update for given router
func linkStatesForRouter(router *Router) []LinkState {
	var states []LinkState
	for _, neighbor := range router.Neighbors {
		states = append(states, LinkState{
			NeighborRouter: neighbor,
			Cost:           rand.Intn(10) + 1,
		})
	}
	return states
}
