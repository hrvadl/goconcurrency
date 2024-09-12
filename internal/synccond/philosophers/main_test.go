package main

import (
	"reflect"
	"testing"
)

func TestRemoveForks(t *testing.T) {
	t.Parallel()
	type args struct {
		forks []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Should remove forks correctly",
			args: args{
				forks: []int{1, 2, 3, 4, 5},
				left:  1,
				right: 2,
			},
			want: []int{3, 4, 5},
		},
		{
			name: "Should remove forks correctly",
			args: args{
				forks: []int{1, 2, 3, 4, 5},
				left:  5,
				right: 1,
			},
			want: []int{2, 3, 4},
		},
		{
			name: "Should remove forks correctly",
			args: args{
				forks: []int{1, 2, 3, 4, 5},
				left:  2,
				right: 3,
			},
			want: []int{1, 4, 5},
		},
		{
			name: "Should remove forks correctly",
			args: args{
				forks: []int{1, 2, 3, 4, 5},
				left:  3,
				right: 4,
			},
			want: []int{1, 2, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := removeForks(tt.args.forks, tt.args.left, tt.args.right); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("removeForks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddForks(t *testing.T) {
	t.Parallel()
	type args struct {
		forks []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Should add forks correctly",
			args: args{
				forks: []int{3, 4, 5},
				left:  1,
				right: 2,
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Should add forks correctly",
			args: args{
				forks: []int{2, 3, 4},
				left:  1,
				right: 5,
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Should add forks correctly",
			args: args{
				forks: []int{1, 2, 3},
				left:  4,
				right: 5,
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := addForks(tt.args.forks, tt.args.left, tt.args.right); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("addForks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNearestForks(t *testing.T) {
	type args struct {
		philosopherIndex int
	}
	tests := []struct {
		name      string
		args      args
		wantLeft  int
		wantRight int
	}{
		{
			name: "Should get nearest forks for the first philosopher correctly",
			args: args{
				philosopherIndex: 1,
			},
			wantLeft:  philosophersCount,
			wantRight: 2,
		},
		{
			name: "Should get nearest forks for the last philosopher correctly",
			args: args{
				philosopherIndex: philosophersCount,
			},
			wantLeft:  philosophersCount - 1,
			wantRight: 1,
		},
		{
			name: "Should get nearest forks for the philosopher in the middle correctly",
			args: args{
				philosopherIndex: 2,
			},
			wantLeft:  1,
			wantRight: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			left, right := getNearestForks(tt.args.philosopherIndex)
			if left != tt.wantLeft {
				t.Errorf("getNearestForks() left = %v, want %v", left, tt.wantLeft)
			}
			if right != tt.wantRight {
				t.Errorf("getNearestForks() right = %v, want %v", right, tt.wantRight)
			}
		})
	}
}

func TestNearestForksIsAvailable(t *testing.T) {
	type args struct {
		leftFork  int
		rightFork int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should calculate availability for the nearest forks correctly",
			args: args{
				leftFork:  1,
				rightFork: 3,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nearestForksIsAvailable(tt.args.leftFork, tt.args.rightFork); got != tt.want {
				t.Errorf("nearestForksIsAvailable() = %v, want %v", got, tt.want)
			}
		})
	}
}
