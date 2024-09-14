package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	intStream := generate(ctx, 1)
	threeInts := take(ctx, intStream, 3)
	fanOutInts := out(ctx, threeInts, sleepAdapter, 5)
	print(ctx, in(ctx, fanOutInts...))
	// print(ctx, sleep(ctx, threeInts, time.Second*5))
	fmt.Printf("Program took %vs to execute", time.Since(start))
}

func sleepAdapter[T any](ctx context.Context, stream <-chan T) <-chan T {
	return sleep(ctx, stream, time.Second*5)
}
