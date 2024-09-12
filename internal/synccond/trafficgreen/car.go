package main

import (
	"fmt"
	"sync"
	"time"
)

const runInterval = time.Second * 2

func NewCar(color string, cond *sync.Cond) *Car {
	return &Car{
		color: color,
		cond:  cond,
	}
}

type Car struct {
	color          string
	cond           *sync.Cond
	isAllowedToRun func() bool
}

func (c *Car) Ride() {
	for {
		c.cond.L.Lock()
		for !c.isAllowedToRun() {
			fmt.Printf("%s car is waiting...!\n", c.color)
			c.cond.Wait()
		}
		fmt.Printf("%s car is running!\n", c.color)
		c.cond.L.Unlock()
		time.Sleep(runInterval)
	}
}

func (c *Car) SetIsAllowedToRunCheck(check func() bool) {
	c.isAllowedToRun = check
}
