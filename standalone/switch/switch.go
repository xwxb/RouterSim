package main

import (
	"fmt"
	"math/rand"
	"time"
)

// SwitchTable stores MAC address to port mappings
type SwitchTable struct {
	Table map[string]int // MAC -> port
}

// EthernetFrame models a frame with dest/src MAC addresses
type EthernetFrame struct {
	DestMAC     string
	SrcMAC      string
	SwitchID    int
	IngressPort int
}

// Switch represents an Ethernet learning switch
type Switch struct {
	Id    int
	Table SwitchTable
	Ports []int // connected ports

	FrameChan chan EthernetFrame
}

// HandleFrame processes incoming Ethernet frames
func (s *Switch) HandleFrame(frame EthernetFrame) {
	// Learn source MAC address
	s.Table.Table[frame.SrcMAC] = frame.IngressPort

	// Look up destination port
	destPort, ok := s.Table.Table[frame.DestMAC]
	if !ok {
		// Flood if destination unknown
		s.Flood(frame)
		return
	}

	// Forward frame out destination port
	s.SendFrame(frame, destPort)
}

// Flood broadcasts frame on all ports except ingress
func (s *Switch) Flood(frame EthernetFrame) {
	for _, port := range s.Ports {
		if port != frame.IngressPort {
			go s.SendFrame(frame, port)
		}
	}
}

// SendFrame simulates forwarding frame out given port
func (s *Switch) SendFrame(frame EthernetFrame, port int) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	fmt.Printf("Forward frame from switch %d out port %d\n", s.Id, port)
}
