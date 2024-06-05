package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

type monitor struct {
	CpuPercent []float64
	MemPercent *mem.VirtualMemoryStat
	Message    string
}

var data monitor

func index(w http.ResponseWriter, r *http.Request) {

	data.CpuPercent, _ = cpu.Percent(0, false)
	data.MemPercent, _ = mem.VirtualMemory()
	data.Message = ""

	if data.CpuPercent[0] > 50 && data.MemPercent.UsedPercent > 50 {
		data.Message = "Warning: High CPU and Memory usage detected. Optimize your application and scale resources if needed."
	} else if data.CpuPercent[0] > 50 {
		data.Message = "Warning: High CPU usage detected. Optimize your application for better performance."
	} else if data.MemPercent.UsedPercent > 50 {
		data.Message = "Warning: High Memory usage detected. Consider scaling your resources."
	} else {
		data.Message = "CPU and Memory usage are within acceptable limits. Keep up the good work!"
	}

	err := tpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

	//fmt.Fprintf(w, "CPU util: %.2f%% and memory utl: %.2f%%\n%s", data.cpuPercent[0], data.memPercent.UsedPercent, message)
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server running on port :8080")
	http.ListenAndServe(":8080", nil)
}
