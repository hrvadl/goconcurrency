package main

import (
	"context"
	"log/slog"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan int)
	defer close(ch)

	ch1, ch2 := tee(ctx, ch)
	ch <- 1
	printTee(ctx, ch1, ch2)
	ch <- 2
	printTee(ctx, ch1, ch2)
	ch <- 3
	printTee(ctx, ch1, ch2)
}

func printTee[T any](ctx context.Context, ch1, ch2 <-chan T) {
	for range 2 {
		select {
		case val := <-ch1:
			slog.Info("Got value from chan #1", slog.Any("val", val))
		case val := <-ch2:
			slog.Info("Got value from chan #2", slog.Any("val", val))
		case <-ctx.Done():
			return
		}
	}
}
