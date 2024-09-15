package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/hrvadl/goconcurrency/internal/scalepatterns/heartbeat"
)

const (
	ctxTimeout    = time.Second * 55
	pulseInterval = time.Second * 1
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()

	monitor := heartbeat.NewMonitor(heartbeat.WorkerAdapter)
	pulse := monitor(ctx, pulseInterval)

	for {
		select {
		case <-ctx.Done():
			slog.Info("Timed out.")
			return
		case _, ok := <-pulse:
			if !ok {
				slog.Info("Seems like monitor is not healthy!")
				return
			}
		}
	}
}
