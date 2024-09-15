package main

import "context"

func tee[T any](ctx context.Context, ch <-chan T) (<-chan T, <-chan T) {
	var (
		out1 = make(chan T)
		out2 = make(chan T)
	)

	go func() {
		defer close(out1)
		for val := range orDone(ctx, ch) {
			out1, out2 := out1, out2
			for range 2 {
				select {
				case out1 <- val:
					out1 = nil
				case out2 <- val:
					out2 = nil
				case <-ctx.Done():
					return
				}
			}
		}
	}()

	return out1, out2
}

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
