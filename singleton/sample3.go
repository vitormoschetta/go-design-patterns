package singleton

import (
	"sync"
)

var counterInstance *Counter

type Counter struct {
	value int
	mu    sync.Mutex
}

func NewCounter() *Counter {
	if counterInstance == nil {
		counterInstance = &Counter{}
	}
	return counterInstance
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}
