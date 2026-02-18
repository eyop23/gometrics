package metrics

import (
	"time"
	
	"github.com/eyop23/gometrics/internal/collector"
	"github.com/eyop23/gometrics/internal/exporter"
)

// defaultCollector is the global shared collector
// All metrics go here unless you create your own
var defaultCollector = collector.New()

// Inc increments a counter by 1
func Inc(name string) {
	defaultCollector.GetOrCreateCounter(name).Inc()
}

// IncBy increments a counter by a specific amount
func IncBy(name string, n int64) {
	counter := defaultCollector.GetOrCreateCounter(name)
	for i := int64(0); i < n; i++ {
		counter.Inc()
	}
}

// Set sets a gauge to a specific value
func Set(name string, value int64) {
	defaultCollector.GetOrCreateGauge(name).Set(value)
}

// GaugeInc increments a gauge by 1
func GaugeInc(name string) {
	defaultCollector.GetOrCreateGauge(name).Inc()
}

// GaugeDec decrements a gauge by 1
func GaugeDec(name string) {
	defaultCollector.GetOrCreateGauge(name).Dec()
}

// Track records a duration in a timer
func Track(name string, duration time.Duration) {
	defaultCollector.GetOrCreateTimer(name).Record(duration)
}

// TrackFunc runs a function and tracks how long it took
// Usage: metrics.TrackFunc("db.query", func() { runQuery() })
func TrackFunc(name string, fn func()) {
	start := time.Now()
	fn()
	duration := time.Since(start)
	Track(name, duration)
}

// Serve starts the metrics HTTP server
// Example: metrics.Serve(":9090")
func Serve(addr string) error {
	exp := exporter.New(defaultCollector)
	return exp.Start(addr)
}