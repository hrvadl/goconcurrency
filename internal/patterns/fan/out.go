package main

import (
	"context"
)

func out[T any](
	ctx context.Context,
	stream <-chan T,
	amount int,
) []<-chan T {
	outed := make([]chan T, amount)
	for i := range amount {
		outed[i] = make(chan T)
	}

	go func() {
		defer func() {
			for _, ch := range outed {
				close(ch)
			}
		}()

		for {
			for _, ch := range outed {
				select {
				case <-ctx.Done():
					return
				case val, ok := <-stream:
					if !ok {
						return
					}

					ch <- val
				}
			}
		}
	}()

	readOnlyOuted := make([]<-chan T, amount)
	for i := range outed {
		readOnlyOuted[i] = outed[i]
	}

	return readOnlyOuted
}
