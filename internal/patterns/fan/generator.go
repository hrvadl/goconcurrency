package main

import (
	"context"
	"fmt"
	"time"
)

func generate[T any](ctx context.Context, val T) <-chan T {
	ch := make(chan T)

	go func() {
		defer close(ch)
		for {
			select {
			case ch <- val:
			case <-ctx.Done():
				return
			}
		}
	}()

	return ch
}

func take[T any](ctx context.Context, stream <-chan T, count int) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for range count {
			select {
			case val, ok := <-stream:
				if !ok {
					return
				}
				ch <- val
			case <-ctx.Done():
				return
			}
		}
	}()

	return ch
}

func sleep[T any](ctx context.Context, stream <-chan T, interval time.Duration) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for val := range orDone(ctx, stream) {
			time.Sleep(interval)
			ch <- val
		}
	}()
	return ch
}

func print[T any](ctx context.Context, stream <-chan T) {
	for val := range orDone(ctx, stream) {
		fmt.Printf("val: %v\n", val)
	}
}
