package main

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"html/template"
	"log"
	"net/http"
	"sync"
)

// Template variable to hold the parsed HTML template
var tpl *template.Template

// Initialize the template
func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

// Monitor struct to store CPU and Memory statistics for JSON response
type Monitor struct {
	CpuPercent float64 `json:"cpuPercent"`
	MemPercent float64 `json:"memPercent"`
}

// Constants for CPU and Memory usage thresholds
const (
	HighUsageThreshold = 50.0
)

// index handler function for the root URL
func index(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Println("Error executing template:", err)
	}
}

// usage handler function for serving CPU and Memory usage data in JSON format
func usage(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	var cpuPercent []float64
	var memStats *mem.VirtualMemoryStat
	var cpuErr, memErr error

	// Use WaitGroup to handle concurrent fetching of CPU and memory usage
	wg.Add(2)

	// Fetch CPU usage concurrently
	go func() {
		defer wg.Done()
		cpuPercent, cpuErr = cpu.Percent(0, false)
	}()

	// Fetch memory usage concurrently
	go func() {
		defer wg.Done()
		memStats, memErr = mem.VirtualMemory()
	}()

	wg.Wait()

	// Check for errors and handle them
	if cpuErr != nil {
		http.Error(w, "Error fetching CPU usage: "+cpuErr.Error(), http.StatusInternalServerError)
		return
	}
	if memErr != nil {
		http.Error(w, "Error fetching memory usage: "+memErr.Error(), http.StatusInternalServerError)
		return
	}

	// Create a Monitor instance with the fetched data
	monitorData := Monitor{
		CpuPercent: cpuPercent[0],
		MemPercent: memStats.UsedPercent,
	}

	// Convert the Monitor instance to JSON format
	jsonResponse, err := json.MarshalIndent(monitorData, "", "  ")
	if err != nil {
		http.Error(w, "Error generating JSON response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header and write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// main function to start the HTTP server
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/usage", usage)
	fmt.Println("Server running on port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
