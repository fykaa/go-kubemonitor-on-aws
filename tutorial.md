
## 1 . main file for core functionality
Creating main go file
- Make sure you have the github.com/shirou/gopsutil/cpu and github.com/shirou/gopsutil/mem packages installed to run this code. You can install them using go get:
go get github.com/shirou/gopsutil/cpu
go get github.com/shirou/gopsutil/mem

- This code sets up a simple HTTP server that listens on port 8080, and the / route handler function fetches CPU and memory utilization information using the gopsutil package, then returns the data along with a message indicating high utilization if applicable.

