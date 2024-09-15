package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/hrvadl/goconcurrency/internal/scalepatterns/heartbeat"
)

const (
	ctxTimeout      = time.Second * 55
	pulseInterval   = time.Second
	timeoutInterval = pulseInterval * 2
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()

	stream, hearbeat := heartbeat.Worker(ctx, pulseInterval)

	for {
		select {
		case <-ctx.Done():
			slog.Info("Timed out.")
			return
		case _, ok := <-hearbeat:
			if !ok {
				return
			}
			slog.Info("Got hearbeat from the worker")
		case val := <-stream:
			slog.Info("Got value from the stream", slog.Any("val", val))
		case <-time.After(timeoutInterval):
			slog.Info("Something wrong with the worker")
			return
		}
	}
}
