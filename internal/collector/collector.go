package collector

import "sync"

// Collector holds all metrics in memory
type Collector struct {
	mu       sync.RWMutex        // protects all three maps
	counters map[string]*Counter // stores all counters by name
	gauges   map[string]*Gauge   // stores all gauges by name
	timers   map[string]*Timer   // stores all timers by name
}

// New creates and returns a new empty Collector
func New() *Collector {
	return &Collector{
		counters: make(map[string]*Counter),
		gauges:   make(map[string]*Gauge),
		timers:   make(map[string]*Timer),
	}
}

// GetOrCreateCounter returns existing counter or creates a new one
func (c *Collector) GetOrCreateCounter(name string) *Counter {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.counters[name]; !ok {
		c.counters[name] = &Counter{}
	}
	return c.counters[name]
}

// GetOrCreateGauge returns existing gauge or creates a new one
func (c *Collector) GetOrCreateGauge(name string) *Gauge {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.gauges[name]; !ok {
		c.gauges[name] = &Gauge{}
	}
	return c.gauges[name]
}

// GetOrCreateTimer returns existing timer or creates a new one
func (c *Collector) GetOrCreateTimer(name string) *Timer {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.timers[name]; !ok {
		c.timers[name] = &Timer{}
	}
	return c.timers[name]
}

// Counters returns all counters
func (c *Collector) Counters() map[string]*Counter {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.counters
}

// Gauges returns all gauges
func (c *Collector) Gauges() map[string]*Gauge {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.gauges
}

// Timers returns all timers
func (c *Collector) Timers() map[string]*Timer {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.timers
}