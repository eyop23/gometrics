package exporter

import (
	"fmt"
	"net/http"
	
	"github.com/eyop23/gometrics/internal/collector"
)

// Exporter exposes metrics via HTTP
type Exporter struct {
	collector *collector.Collector
}

// New creates a new Exporter
func New(c *collector.Collector) *Exporter {
	return &Exporter{
		collector: c,
	}
}

// ServeHTTP handles HTTP requests to the /metrics endpoint
func (e *Exporter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	
	for name, counter := range e.collector.Counters() {
		fmt.Fprintf(w, "%s %d\n", name, counter.Value())
	}
	
	for name, gauge := range e.collector.Gauges() {
		fmt.Fprintf(w, "%s %d\n", name, gauge.Value())
	}
	
	for name, timer := range e.collector.Timers() {
		avgMs := timer.Avg().Milliseconds()
		fmt.Fprintf(w, "%s_avg_ms %d\n", name, avgMs)
		fmt.Fprintf(w, "%s_count %d\n", name, timer.Count())
	}
}

// Start starts the HTTP server on the given address
func (e *Exporter) Start(addr string) error {
	http.Handle("/metrics", e)
	fmt.Printf("Metrics server listening on %s/metrics\n", addr)
	return http.ListenAndServe(addr, nil)
}