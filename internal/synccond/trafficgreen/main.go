package main

import (
	"sync"
	"time"
)

// Exercise: Implementing a TrafficLight program
// Problem:

// There're limited amount of cars which can ride. Also, there's some traffich lights
// on the road, which can force cars to stop, if they don't show a green light.
// Cars should ride only when traffic light shows green, otherwise they should wait
// for the traffict light to start showing green.

const changeTrafficLightInterval = time.Second * 5

func main() {
	done := make(chan struct{})
	cond := sync.NewCond(&sync.Mutex{})
	cars := []Car{
		*NewCar("red", cond),
		*NewCar("yellow", cond),
		*NewCar("green", cond),
		*NewCar("black", cond),
		*NewCar("white", cond),
	}

	ctrl := NewTrafficController(cars, cond, changeTrafficLightInterval)
	ctrl.Run()
	<-done
}
