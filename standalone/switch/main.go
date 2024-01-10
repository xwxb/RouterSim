package main

import (
	"time"
)

// Network with two connected switches
var switches []*Switch

func initializeSwitches() {
	switches = make([]*Switch, 2)

	for i := 0; i < 2; i++ {
		s := &Switch{
			Id:        i + 1,
			Table:     SwitchTable{make(map[string]int)},
			Ports:     []int{1, 2}, // simulate two ports
			FrameChan: make(chan EthernetFrame),
		}
		switches[i] = s

		// connect switches
		if i == 0 {
			s.Ports[1] = 3 // port 3 connects to other switch
		} else {
			s.Ports[0] = 3 // port 3 connects to other switch
		}
	}
}

func main() {
	initializeSwitches()

	// process frames
	for _, s := range switches {
		go func(switchId int) {
			for {
				frame := <-switches[switchId].FrameChan
				switches[switchId].HandleFrame(frame)
			}
		}(s.Id - 1)
	}

	// send test frames
	sendFrame("AA:AA:AA:AA:AA:AA", "BB:BB:BB:BB:BB:BB", 1, 1) // unknown dest
	sendFrame("AA:AA:AA:AA:AA:AA", "BB:BB:BB:BB:BB:BB", 1, 1) // flood
	sendFrame("BB:BB:BB:BB:BB:BB", "AA:AA:AA:AA:AA:AA", 2, 2) // known dest

	time.Sleep(3 * time.Second)
}

// sendFrame simulates sending a frame into the network
func sendFrame(destMAC string, srcMAC string, switchId, ingressPort int) {
	frame := EthernetFrame{
		DestMAC:     destMAC,
		SrcMAC:      srcMAC,
		SwitchID:    switchId,
		IngressPort: ingressPort,
	}
	switches[switchId-1].FrameChan <- frame
}
