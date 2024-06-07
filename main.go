package main

import (
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

// Monitor struct to store CPU and Memory statistics and a message
type monitor struct {
	CpuPercent []float64
	MemPercent *mem.VirtualMemoryStat
	Message    string
}

// Data variable to hold the monitor struct
var data monitor

// Thresholds for CPU and Memory usage
const (
	HighUsageThreshold = 50.0
)

// Index handler function for the root URL
func index(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	wg.Add(2)

	// Fetch CPU usage concurrently
	go func() {
		defer wg.Done()
		cpuPercent, err := cpu.Percent(0, false)
		if err != nil {
			log.Println("Error fetching CPU percent:", err)
			return
		}
		data.CpuPercent = cpuPercent
	}()

	// Fetch memory usage concurrently
	go func() {
		defer wg.Done()
		memPercent, err := mem.VirtualMemory()
		if err != nil {
			log.Println("Error fetching memory percent:", err)
			return
		}
		data.MemPercent = memPercent
	}()

	wg.Wait()

	// Determine the message based on the thresholds
	data.Message = ""
	if data.CpuPercent[0] > HighUsageThreshold && data.MemPercent.UsedPercent > HighUsageThreshold {
		data.Message = "Warning: High CPU and Memory usage detected. Optimize your application and scale resources if needed."
	} else if data.CpuPercent[0] > HighUsageThreshold {
		data.Message = "Warning: High CPU usage detected. Optimize your application for better performance."
	} else if data.MemPercent.UsedPercent > HighUsageThreshold {
		data.Message = "Warning: High Memory usage detected. Consider scaling your resources."
	} else {
		data.Message = "CPU and Memory usage are within acceptable limits. Keep up the good work!"
	}

	// Execute the template with the monitor data
	err := tpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Println("Error executing template:", err)
	}
}

// Main function to start the HTTP server
func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server running on port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
