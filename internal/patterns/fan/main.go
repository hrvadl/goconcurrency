package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	start := time.Now()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream := getIntStream(ctx)
	print(ctx, stream)
	// print(ctx, stream)
	fmt.Printf("Program took %vs to execute", time.Since(start))
}

func getIntStream(ctx context.Context) <-chan int {
	intStream := generate(ctx, 1)
	threeInts := take(ctx, intStream, 3)
	fanOutInts := out(ctx, threeInts, sleepAdapter, runtime.NumCPU())
	return in(ctx, fanOutInts...)
}

func sleepAdapter[T any](ctx context.Context, stream <-chan T) <-chan T {
	return sleep(ctx, stream, time.Second*1)
}
