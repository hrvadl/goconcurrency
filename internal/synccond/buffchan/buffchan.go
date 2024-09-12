package buffchan

// My very simple implementation of buffered channels
// It's main purpose is to play with the sync.Cond struct
// to understand it more deeply.

import "sync"

func New[T any](size int) *BufferedChannel[T] {
	mu := &sync.Mutex{}
	return &BufferedChannel[T]{
		size: size,
		data: make([]T, 0, size),
		mu:   mu,
		cond: sync.NewCond(mu),
	}
}

type BufferedChannel[T any] struct {
	data []T
	size int
	mu   *sync.Mutex
	cond *sync.Cond
}

func (c *BufferedChannel[T]) Put(val T) {
	c.cond.L.Lock()
	for len(c.data) == c.size {
		c.cond.Wait()
	}
	c.cond.L.Unlock()

	c.data = append(c.data, val)
	c.cond.Signal()
}

func (c *BufferedChannel[T]) Read() T {
	c.cond.L.Lock()
	for len(c.data) == 0 {
		c.cond.Wait()
	}
	c.cond.L.Unlock()

	data := c.data[0]
	c.data = c.data[1:]
	c.cond.Signal()
	return data
}
