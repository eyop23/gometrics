# gometrics

A lightweight, zero-dependency metrics collection library for Go.

## Features

- **Thread-safe** — built for concurrent Go applications
- **Zero dependencies** — uses only Go standard library
- **Simple API** — one-line metric tracking
- **Prometheus compatible** — industry-standard format
- **Three metric types:**
  - **Counter** — tracks totals (only increases)
  - **Gauge** — tracks current values (increases/decreases)
  - **Timer** — measures durations

## Installation
```bash
go get github.com/eyop23/gometrics
```

## Quick Start
```go
package main

import (
    "net/http"
    "time"
    "github.com/eyop23/gometrics/pkg/metrics"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // Track requests
    metrics.Inc("http.requests")
    
    // Track active connections
    metrics.GaugeInc("active.connections")
    defer metrics.GaugeDec("active.connections")
    
    // Track response time
    metrics.TrackFunc("api.response.time", func() {
        // your business logic here
        time.Sleep(time.Millisecond * 100)
    })
    
    w.Write([]byte("Hello World"))
}

func main() {
    http.HandleFunc("/", handler)
    
    // Start metrics server
    go metrics.Serve(":9090")
    
    // Start application
    http.ListenAndServe(":8080", nil)
}
```

## Usage

### Counters (only increase)
```go
metrics.Inc("user.signups")           // increment by 1
metrics.IncBy("bytes.sent", 1024)     // increment by specific amount
```

### Gauges (increase or decrease)
```go
metrics.Set("queue.size", 42)         // set to specific value
metrics.GaugeInc("active.users")      // increment by 1
metrics.GaugeDec("active.users")      // decrement by 1
```

### Timers (measure duration)
```go
// Manual timing
start := time.Now()
doWork()
metrics.Track("work.duration", time.Since(start))

// Automatic timing
metrics.TrackFunc("database.query", func() {
    db.Query("SELECT * FROM users")
})
```

## Viewing Metrics

Start your application and visit:
```
http://localhost:9090/metrics
```

Example output:
```
http.requests 1543
active.connections 12
user.signups 89
api.response.time_avg_ms 120
api.response.time_count 1543
```

## Integration with Monitoring Tools

This library outputs metrics in Prometheus format. You can:
- Point **Prometheus** at the `/metrics` endpoint
- Create **Grafana** dashboards from the data
- Set up alerts based on thresholds

## License
