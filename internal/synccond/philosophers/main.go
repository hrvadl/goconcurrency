package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/exp/constraints"
)

// Exercise: Implementing a Dining Philosophers Problem
// Problem:

// Five philosophers are sitting at a round table, each with a plate of spaghetti in front of them.
// There is one fork between each pair of philosophers. To eat, a philosopher needs to use two forks:
// one to the left and one to the right.

// Explanation:

// Create forks and philosophers: Each philosopher is associated with two forks: one to the left and one to the right.
// Dine function:
// Think: The philosopher thinks.
// PickUpForks: The philosopher acquires both forks.
// Eat: The philosopher eats.
// PutDownForks: The philosopher releases both forks and signals the next philosopher.

const (
	philosophersCount = 5
	forksCount        = philosophersCount - 1
)

var (
	forks = []int{1, 2, 3, 4, 5}
	cond  = sync.NewCond(&sync.Mutex{})
)

func think() {
	time.Sleep(time.Second * 1)
}

func eat() {
	time.Sleep(time.Second * 1)
}

func dine(philosopherIndex int, wg *sync.WaitGroup) {
	defer wg.Done()
	think()
	fmt.Printf("Philosopher #%d thought...\n", philosopherIndex)
	pickUpForks(philosopherIndex)
	fmt.Printf("Philosopher #%d picked up forks...\n", philosopherIndex)
	eat()
	fmt.Printf("Philosopher #%d ate...\n", philosopherIndex)
	putBackForks(philosopherIndex)
	fmt.Printf("Philosopher #%d put back forks...\n", philosopherIndex)
}

func putBackForks(philosopher int) {
	cond.L.Lock()
	leftFork, rightFork := getNearestForks(philosopher)
	forks = addForks(forks, leftFork, rightFork)
	fmt.Printf(
		"Philosopher #%d put back: all forks: %v, left fork: %d, right fork: %d\n",
		philosopher,
		forks,
		leftFork,
		rightFork,
	)
	cond.L.Unlock()
	cond.Broadcast()
}

func pickUpForks(philosopher int) {
	cond.L.Lock()
	leftFork, rightFork := getNearestForks(philosopher)
	for !nearestForksIsAvailable(leftFork, rightFork) {
		fmt.Printf("philosopher: %d, nearest fork is not available\n", philosopher)
		cond.Wait()
	}

	fmt.Printf("philosopher: %d, nearest fork is available\n", philosopher)
	forks = removeForks(forks, leftFork, rightFork)
	fmt.Printf(
		"Philosopher #%d picked up forks: all forks: %v, left fork: %d, right fork: %d\n",
		philosopher,
		forks,
		leftFork,
		rightFork,
	)
	cond.L.Unlock()
}

func nearestForksIsAvailable(leftFork, rightFork int) bool {
	const totalAvailableForksPerPhilosopher = 2
	availableForks := make([]int, 0, totalAvailableForksPerPhilosopher)
	for _, fork := range forks {
		if fork == leftFork || fork == rightFork {
			availableForks = append(availableForks, fork)
		}
	}
	return len(availableForks) == totalAvailableForksPerPhilosopher
}

func getNearestForks(philosopher int) (int, int) {
	if isFirstPhilosopher(philosopher) {
		return philosophersCount, philosopher + 1
	}

	if isLastPhilosopher(philosopher) {
		return philosopher - 1, 1
	}

	return philosopher - 1, philosopher + 1
}

func isLastPhilosopher(philosopher int) bool {
	return philosopher == philosophersCount
}

func isFirstPhilosopher(philosopher int) bool {
	return philosopher == 1
}

func removeForks(forks []int, left, right int) []int {
	beforeLeft := filterLess(left, forks)
	afterLeft := filterMore(left, forks)
	withoutLeft := append(copied(beforeLeft), afterLeft...)

	beforeRight := filterLess(right, withoutLeft)
	afterRight := filterMore(right, withoutLeft)
	withoutRight := append(copied(beforeRight), afterRight...)

	return withoutRight
}

func copied[T any](slice []T) []T {
	copiedSlice := make([]T, len(slice))
	copy(copiedSlice, slice)
	return copiedSlice
}

func addForks(forks []int, left, right int) []int {
	beforeLeft := filterLess(left, forks)
	afterLeft := filterMore(left, forks)
	withLeft := append(append(beforeLeft, left), afterLeft...)

	beforeRight := filterLess(right, withLeft)
	afterRight := filterMore(right, withLeft)
	withRight := append(append(beforeRight, right), afterRight...)

	return withRight
}

func filterMore[T constraints.Ordered](than T, slice []T) []T {
	var result []T
	for _, el := range slice {
		if el > than {
			result = append(result, el)
		}
	}
	return result
}

func filterLess[T constraints.Ordered](than T, slice []T) []T {
	var result []T
	for _, el := range slice {
		if el < than {
			result = append(result, el)
		}
	}
	return result
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= philosophersCount; i++ {
		wg.Add(1)
		go dine(i, &wg)
	}
	wg.Wait()
}
