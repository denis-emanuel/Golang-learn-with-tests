package sync

import "sync"

// Counter will increment a number.
type Counter struct {
	//A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.
	mu    sync.Mutex
	value int
}

// NewCounter returns a new Counter.
func NewCounter() *Counter {
	return &Counter{}
}

// If we would've passed the counter by value it would've tried to create a copy of the mutex which is wrong
func (c *Counter) Inc() {
	// Any goroutine calling Inc() will acquire the lock on the Counter if they are first. 
	c.mu.Lock()
	// All the other goroutines wait for the Unlock() to be called
	defer c.mu.Unlock()
	// Each goroutine has to wait before making a change
	c.value++
}

// Value returns the current count.
func (c *Counter) Value() int {
	return c.value
}