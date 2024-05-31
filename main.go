package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	cpuPercent, _ := cpu.Percent(0, false)
	memPercent, _ := mem.VirtualMemory()
	message := ""
	if cpuPercent[0] > 2 || memPercent.UsedPercent > 10 {
		message = "High CPU or Memory utilization. Please scale up!!"
	}
	fmt.Fprintf(w, "CPU util: %.2f%% and memory utl: %.2f%%\n%s", cpuPercent[0], memPercent.UsedPercent, message)
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server running on port :8080")
	http.ListenAndServe(":8080", nil)
}
