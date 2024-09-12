package main

import (
	"sync"
	"time"
)

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
