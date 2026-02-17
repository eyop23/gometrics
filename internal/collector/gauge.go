package collector

import "sync/atomic"

// Gauge is a metric that can go up or down.
// Example use: tracking current active connections.
type Gauge struct {
	value int64
}

// Inc increases the gauge by 1
func (g *Gauge) Inc() {
	atomic.AddInt64(&g.value, 1)
}

// Dec decreases the gauge by 1
func (g *Gauge) Dec() {
	atomic.AddInt64(&g.value, -1)
}

// Set sets the gauge to a specific value
func (g *Gauge) Set(v int64) {
	atomic.StoreInt64(&g.value, v)
}

// Value returns the current value of the gauge
func (g *Gauge) Value() int64 {
	return atomic.LoadInt64(&g.value)
}