package main

import "sync"

type CountUp interface {
	Inc()
	Value() int
}
type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.count += 1
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	return c.count
}
