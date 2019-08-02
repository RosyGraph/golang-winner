package main

import "fmt"

type car struct {
	gasPedal      uint16
	brakePedal    uint16
	steeringWheel uint16
	topSpeedKMH   uint16
}

func main() {
	aCar := car{gasPedal: 22341,
		brakePedal:    0,
		steeringWheel: 13451,
		topSpeedKMH:   225.0,
	}
	fmt.Println(aCar.gasPedal)
}
