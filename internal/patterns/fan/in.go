package main

import (
	"context"
	"sync"
)

func in[T any](ctx context.Context, streams ...<-chan T) <-chan T {
	aggregated := make(chan T)
	var wg sync.WaitGroup

	multiplex := func(stream <-chan T) {
		defer wg.Done()
		for value := range orDone(ctx, stream) {
			aggregated <- value
		}
	}

	wg.Add(len(streams))
	for _, stream := range streams {
		go multiplex(stream)
	}

	go func() {
		defer close(aggregated)
		wg.Wait()
	}()

	return aggregated
}
