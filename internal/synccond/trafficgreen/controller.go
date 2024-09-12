package main

import (
	"sync"
	"time"
)

func NewTrafficController(cars []Car, cond *sync.Cond, interval time.Duration) *TrafficController {
	return &TrafficController{
		cars:         cars,
		interval:     interval,
		cond:         cond,
		trafficLight: NewTrafficLight(),
	}
}

type TrafficController struct {
	cars         []Car
	trafficLight *TrafficLight
	interval     time.Duration
	cond         *sync.Cond
}

func (tc *TrafficController) Run() {
	go tc.changeTrafficLightsColor()
	var wg sync.WaitGroup
	wg.Add(len(tc.cars))

	for _, car := range tc.cars {
		car.SetIsAllowedToRunCheck(tc.trafficLight.IsAllowedToRide)
		go car.Ride()
	}
}

func (tc *TrafficController) changeTrafficLightsColor() {
	ticker := time.NewTicker(tc.interval)
	for range ticker.C {
		tc.cond.L.Lock()
		tc.trafficLight.NextColor()
		tc.cond.L.Unlock()
		tc.cond.Broadcast()
	}
}
