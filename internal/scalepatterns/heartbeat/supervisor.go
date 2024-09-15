package heartbeat

import (
	"context"
	"log/slog"
	"time"
)

type Monitorable func(ctx context.Context, pulseInteval time.Duration) <-chan struct{}

func NewMonitor(fn Monitorable) Monitorable {
	return func(ctx context.Context, pulseInteval time.Duration) <-chan struct{} {
		var (
			timeout     = pulseInteval * 2
			pulse       = make(chan struct{})
			pulseTicker = time.Tick(pulseInteval)
			workerPulse <-chan struct{}
		)

		go func() {
			defer close(pulse)

			startWard := func() {
				workerPulse = fn(ctx, pulseInteval)
			}

			sendPulse := func() {
				select {
				case pulse <- struct{}{}:
				default:
				}
			}

			startWard()

			for {
				select {
				case <-pulseTicker:
					sendPulse()
				case _, ok := <-workerPulse:
					if !ok {
						slog.Info("Oh... seems like worker is not healthy... Heartbeat failed!")
						startWard()
					}
					slog.Info("Worker is healthy!")
				case <-time.After(timeout):
					slog.Warn("Oh... seems like worker is not healthy... Timed out!")
					startWard()
				case <-ctx.Done():
					return
				}
			}
		}()

		return pulse
	}
}
