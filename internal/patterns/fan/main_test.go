package main

import (
	"context"
	"reflect"
	"testing"
)

func TestGetIntStream(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Should get ints from the stream correctly",
			args: args{
				ctx: context.Background(),
			},
			want: []int{1, 1, 1},
		},
		{
			name: "Should cancel pipeline",
			args: args{
				ctx: newCancelledCtx(),
			},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			stream := getIntStream(tt.args.ctx)
			got := make([]int, 0, len(tt.want))
			for val := range stream {
				got = append(got, val)
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("Want: %v, got: %v", tt.want, got)
			}
		})
	}
}

func newCancelledCtx() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	return ctx
}
