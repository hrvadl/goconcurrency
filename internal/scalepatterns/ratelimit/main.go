package main

import (
	"context"
	"log/slog"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	lim := NewMultiLimiter(
		rate.NewLimiter(Per(1, time.Second*1), 1),
		rate.NewLimiter(Per(10, time.Minute*1), 10),
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for {
		if err := lim.Wait(ctx); err != nil {
			panic(err)
		}
		printHello()
	}
}

func printHello() {
	slog.Info("hello!")
}
