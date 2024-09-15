package main

import (
	"context"
	"math/rand/v2"
	"time"
)

const workInterval = time.Second * 3

func worker(ctx context.Context, pulseInterval time.Duration) (<-chan int, <-chan struct{}) {
	var (
		resCh       = make(chan int)
		pulse       = make(chan struct{})
		pulseTicker = time.Tick(pulseInterval)
		workTicker  = time.Tick(workInterval)
	)

	closeAll := func() {
		close(resCh)
		close(pulse)
	}

	sendPulse := func() {
		select {
		case pulse <- struct{}{}:
		default:
		}
	}

	sendWork := func() {
		select {
		case resCh <- rand.Int():
		case <-ctx.Done():
			return
		}
	}

	go func() {
		defer closeAll()
		for {
			select {
			case <-workTicker:
				sendWork()
			case <-pulseTicker:
				sendPulse()
			case <-ctx.Done():
				return
			}
		}
	}()

	return resCh, pulse
}
