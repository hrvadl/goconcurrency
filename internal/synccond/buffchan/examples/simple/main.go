package main

import (
	"fmt"
	"sync"

	"github.com/hrvadl/goconcurrency/internal/synccond/buffchan"
)

func main() {
	ch := buffchan.New[int](3)
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		for range 3 {
			fmt.Printf("Read value from channel: %v\n", ch.Read())
		}
	}()

	ch.Put(1)
	ch.Put(2)
	ch.Put(3)
	wg.Wait()
}
