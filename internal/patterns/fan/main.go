package main

import (
	"context"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	intStream := generate(ctx, 1)
	threeInts := take(ctx, intStream, 3)
	fanOutInts := out(ctx, threeInts, 5)
	// TODO: sleep?
	print(ctx, in(ctx, fanOutInts...))
}
