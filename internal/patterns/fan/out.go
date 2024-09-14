package main

import (
	"context"
)

type generatorFn[T any] func(ctx context.Context, stream <-chan T) <-chan T

func out[T any](
	ctx context.Context,
	stream <-chan T,
	fn generatorFn[T],
	amount int,
) []<-chan T {
	outed := make([]<-chan T, amount)
	for i := range amount {
		outed[i] = fn(ctx, stream)
	}
	return outed
}
