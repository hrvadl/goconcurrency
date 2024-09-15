package main

import (
	"fmt"
	"testing"
)

func BenchmarkPrime(b *testing.B) {
	tc := []struct {
		num int
	}{
		{num: 100},
		{num: 1000},
		{num: 10000},
		{num: 100000},
	}

	for _, tt := range tc {
		b.Run(fmt.Sprintf("input size %d", tt.num), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				prime(tt.num)
			}
		})
	}
}
