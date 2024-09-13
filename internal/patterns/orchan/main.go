package main

import (
	"log/slog"
	"runtime"
	"time"
)

func main() {
	doneCh := or(
		time.After(time.Minute*5),
		time.After(time.Minute*1),
		time.After(time.Minute*5),
		time.After(time.Minute*6),
		time.After(time.Minute*2),
		time.After(time.Minute*5),
		time.After(time.Minute*6),
		time.After(time.Second*2),
		time.After(time.Minute*5),
		time.After(time.Minute*6),
		time.After(time.Minute*2),
	)
	slog.Info(
		"Waiting for one of the channels...",
		slog.Int("goroutines_num", runtime.NumGoroutine()),
	)

	<-doneCh
	slog.Info("Finished waiting!", slog.Int("goroutines_num", runtime.NumGoroutine()))
	slog.Info("Sleeping for 3 seconds...", slog.Int("goroutines_num", runtime.NumGoroutine()))
	time.Sleep(time.Second * 3)
	slog.Info("Slept! Now awake!!", slog.Int("goroutines_num", runtime.NumGoroutine()))
}

func or[T any](channels ...<-chan T) <-chan T {
	if len(channels) == 0 {
		return nil
	}

	if len(channels) == 1 {
		return channels[0]
	}

	orCh := make(chan T)

	go func() {
		slog.Info(
			"Waiting for one of the channels...",
			slog.Int("goroutines_num", runtime.NumGoroutine()),
		)
		defer close(orCh)
		select {
		case <-channels[0]:
		case <-channels[1]:
		case <-or(append(channels[2:], orCh)...):
		}
	}()

	return orCh
}
