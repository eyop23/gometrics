package collector

import "sync/atomic"

// Counter is a metric that only ever increases.
// Example use: counting total HTTP requests.
type Counter struct {
	value int64 // the current count, uses int64 for atomic operations
}

// Inc increments the counter by 1 safely across goroutines
func (c *Counter) Inc() {
	atomic.AddInt64(&c.value, 1)
}

// Value returns the current count
func (c *Counter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}