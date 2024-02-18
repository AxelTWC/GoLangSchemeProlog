package main

import (
	"fmt"
)

const (
	Idle    = 0
	Start   = 1
	Forward = 2
	Fast    = 3
	Reverse = -1
)

func statusPrint(state int8) {
	var m = make(map[int]string)
	m[Idle] = "Idle"
	m[Start] = "Start"
	m[Forward] = "Forward"
	m[Fast] = "Fast"
	m[Reverse] = "Reverse"

	switch state {
	case Idle:
		fmt.Printf("State is %s (%d)\n", m[Idle], Idle)
	case Start:
		fmt.Printf("State is %s (%d)\n", m[Start], Start)
	case Forward:
		fmt.Printf("State is %s (%d)\n", m[Forward], Forward)
	case Fast:
		fmt.Printf("State is %s (%d)\n", m[Fast], Fast)
	case Reverse:
		fmt.Printf("State is %s (%d)\n", m[Reverse], Reverse)
	default:
		fmt.Printf("Invalid state: %d\n", state)
	}
	return
}

func main() {
	var i int8
	for i = -1; i < 5; i++ {
		statusPrint(i)
	}
}
