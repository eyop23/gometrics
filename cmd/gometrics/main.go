package main

import (
	"fmt"
	"net/http"
	"time"
	
	"github.com/eyop23/gometrics/pkg/metrics"
)

// handler simulates a real API endpoint
func handler(w http.ResponseWriter, r *http.Request) {
	// Track this request
	metrics.Inc("http.requests")
	
	// Track active connections
	metrics.GaugeInc("active.connections")
	defer metrics.GaugeDec("active.connections")
	
	// Simulate some work and track how long it takes
	metrics.TrackFunc("api.response.time", func() {
		time.Sleep(time.Millisecond * 100) // simulate work
	})
	
	// Send response
	w.Write([]byte("Hello! This request was tracked.\n"))
}

// slowHandler simulates a slow endpoint
func slowHandler(w http.ResponseWriter, r *http.Request) {
	metrics.Inc("http.requests.slow")
	
	metrics.TrackFunc("slow.api.response.time", func() {
		time.Sleep(time.Second * 2) // simulate slow work
	})
	
	w.Write([]byte("This was a slow request.\n"))
}

func main() {
	fmt.Println("Starting real web server with metrics...")
	
	// Register application routes
	http.HandleFunc("/", handler)
	http.HandleFunc("/slow", slowHandler)
	
	// Start metrics server in background
	go func() {
		fmt.Println("Metrics available at: http://localhost:9090/metrics")
		if err := metrics.Serve(":9090"); err != nil {
			fmt.Printf("Metrics server error: %v\n", err)
		}
	}()
	
	// Start main application server
	fmt.Println("Application server running at: http://localhost:8086")
	fmt.Println("\nTry these:")
	fmt.Println("  curl http://localhost:8086/")
	fmt.Println("  curl http://localhost:8086/slow")
	fmt.Println("  curl http://localhost:9090/metrics")
	fmt.Println("\nPress Ctrl+C to stop\n")
	
	if err := http.ListenAndServe(":8086", nil); err != nil {
		fmt.Printf("Application server error: %v\n", err)
	}
}