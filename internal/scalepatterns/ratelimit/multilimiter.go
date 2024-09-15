package main

import (
	"context"
	"slices"
	"time"

	"golang.org/x/time/rate"
)

func NewMultiLimiter(l ...Limiter) *MultiLimiter {
	slices.SortFunc(l, func(a Limiter, b Limiter) int {
		if diff := a.Limit() - b.Limit(); diff < 0 {
			return -1
		} else if diff > 0 {
			return 1
		}
		return 0
	})
	return &MultiLimiter{
		limiters: l,
	}
}

type Limiter interface {
	Wait(ctx context.Context) error
	Limit() rate.Limit
}

type MultiLimiter struct {
	limiters []Limiter
}

func (ml *MultiLimiter) Wait(ctx context.Context) error {
	for _, lim := range ml.limiters {
		if err := lim.Wait(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (ml *MultiLimiter) Limit() rate.Limit {
	for _, lim := range ml.limiters {
		return lim.Limit()
	}
	return 0
}

func Per(eventCount int, dur time.Duration) rate.Limit {
	return rate.Every(dur / time.Duration(eventCount))
}
