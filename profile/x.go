package profile

import (
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// type Counter struct {
// 	count atomic.Int64
// }

// func (c *Counter) Inc() {
// 	c.count.Add(1)
// }
