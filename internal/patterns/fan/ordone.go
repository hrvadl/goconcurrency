package main

import (
	"context"
)

func orDone[T any](ctx context.Context, ch <-chan T) <-chan T {
	result := make(chan T)

	go func() {
		defer close(result)

		for {
			select {
			case val, ok := <-ch:
				if !ok {
					return
				}
				select {
				case result <- val:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return result
}
