package collector

import (
	"sync"
	"time"
)

// Timer records how long things take.
// Example use: measuring database query duration.
type Timer struct {
	mu     sync.Mutex     // protects the values slice
	values []time.Duration // stores all recorded durations
}

// Record adds a new duration to the timer
func (t *Timer) Record(d time.Duration) {
	t.mu.Lock()         // lock before writing
	defer t.mu.Unlock() // unlock when function finishes
	t.values = append(t.values, d)
}

// Avg returns the average duration across all recorded values
func (t *Timer) Avg() time.Duration {
	t.mu.Lock()
	defer t.mu.Unlock()
	if len(t.values) == 0 {
		return 0
	}
	var total time.Duration
	for _, v := range t.values {
		total += v
	}
	return total / time.Duration(len(t.values))
}

// Count returns how many durations have been recorded
func (t *Timer) Count() int {
	t.mu.Lock()
	defer t.mu.Unlock()
	return len(t.values)
}