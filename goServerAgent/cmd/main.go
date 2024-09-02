package main

import (
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/load"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
	"github.com/shirou/gopsutil/v4/process"
)

func main() {
	// Periodically collect and display metrics
	ticker := time.NewTicker(5 * time.Second) // Adjust the interval as needed
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			collectMetrics()
		}
	}
}

func collectMetrics() {
	fmt.Println("----- Server Metrics -----")

	// Memory Info
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("Error fetching virtual memory info: %v\n", err)
	} else {
		fmt.Printf("Memory Total: %.2f GB\n", float64(memInfo.Total)/1024/1024/1024)
		fmt.Printf("Memory Available: %.2f GB\n", float64(memInfo.Available)/1024/1024/1024)
		fmt.Printf("Memory Used: %.2f GB\n", float64(memInfo.Used)/1024/1024/1024)
		fmt.Printf("Memory Free: %.2f GB\n", float64(memInfo.Free)/1024/1024/1024)
		fmt.Printf("Memory Used Percent: %.2f%%\n", memInfo.UsedPercent)
	}

	// Swap Memory Info
	swapInfo, err := mem.SwapMemory()
	if err != nil {
		log.Printf("Error fetching swap memory info: %v\n", err)
	} else {
		fmt.Printf("Swap Total: %.2f GB\n", float64(swapInfo.Total)/1024/1024/1024)
		fmt.Printf("Swap Used: %.2f GB\n", float64(swapInfo.Used)/1024/1024/1024)
		fmt.Printf("Swap Free: %.2f GB\n", float64(swapInfo.Free)/1024/1024/1024)
		fmt.Printf("Swap Used Percent: %.2f%%\n", swapInfo.UsedPercent)
	}

	// CPU Info
	cpuInfo, err := cpu.Info()
	if err != nil {
		log.Printf("Error fetching CPU info: %v\n", err)
	} else {
		for _, cpu := range cpuInfo {
			fmt.Printf("CPU Model: %s\n", cpu.ModelName)
			fmt.Printf("CPU Cores: %d\n", cpu.Cores)
		}
	}

	// CPU Load Averages
	loadAvg, err := load.Avg()
	if err != nil {
		log.Printf("Error fetching load averages: %v\n", err)
	} else {
		fmt.Printf("CPU Load Average (1 min): %.2f\n", loadAvg.Load1)
		fmt.Printf("CPU Load Average (5 min): %.2f\n", loadAvg.Load5)
		fmt.Printf("CPU Load Average (15 min): %.2f\n", loadAvg.Load15)
	}

	// Core-wise CPU Usage
	percentPerCore, err := cpu.Percent(0, true) // Immediate percentage
	if err != nil {
		log.Printf("Error fetching CPU percentages: %v\n", err)
	} else {
		for i, percent := range percentPerCore {
			fmt.Printf("CPU Core %d Usage Percent: %.2f%%\n", i, percent)
		}
	}

	// CPU Times
	cpuTimes, err := cpu.Times(true)
	if err != nil {
		log.Printf("Error fetching CPU times: %v\n", err)
	} else {
		for i, times := range cpuTimes {
			fmt.Printf("CPU Core %d - User: %.2fs, System: %.2fs, Idle: %.2fs, Iowait: %.2fs\n",
				i, times.User, times.System, times.Idle, times.Iowait)
		}
	}

	// Disk Usage
	diskInfo, err := disk.Usage("/")
	if err != nil {
		log.Printf("Error fetching disk usage info: %v\n", err)
	} else {
		fmt.Printf("Disk Total: %.2f GB\n", float64(diskInfo.Total)/1024/1024/1024)
		fmt.Printf("Disk Free: %.2f GB\n", float64(diskInfo.Free)/1024/1024/1024)
		fmt.Printf("Disk Used: %.2f GB\n", float64(diskInfo.Used)/1024/1024/1024)
		fmt.Printf("Disk Used Percent: %.2f%%\n", diskInfo.UsedPercent)
	}

	// Disk I/O Statistics
	ioCounters, err := disk.IOCounters()
	if err != nil {
		log.Printf("Error fetching disk I/O counters: %v\n", err)
	} else {
		for name, stat := range ioCounters {
			fmt.Printf("Disk %s - Read Count: %d, Write Count: %d, Read Bytes: %.2f MB, Write Bytes: %.2f MB\n",
				name, stat.ReadCount, stat.WriteCount, float64(stat.ReadBytes)/1024/1024, float64(stat.WriteBytes)/1024/1024)
		}
	}

	// Network Info
	netInfo, err := net.IOCounters(true)
	if err != nil {
		log.Printf("Error fetching network I/O counters: %v\n", err)
	} else {
		for _, nic := range netInfo {
			fmt.Printf("Network Interface: %s\n", nic.Name)
			fmt.Printf("  Bytes Sent: %.2f MB\n", float64(nic.BytesSent)/1024/1024)
			fmt.Printf("  Bytes Received: %.2f MB\n", float64(nic.BytesRecv)/1024/1024)
			fmt.Printf("  Packets Sent: %d, Packets Received: %d\n", nic.PacketsSent, nic.PacketsRecv)
			fmt.Printf("  Errors In: %d, Errors Out: %d\n", nic.Errin, nic.Errout)
			fmt.Printf("  Drops In: %d, Drops Out: %d\n", nic.Dropin, nic.Dropout)
		}
	}

	// Process Info
	processes, err := process.Processes()
	if err != nil {
		log.Printf("Error fetching processes: %v\n", err)
	} else {
		fmt.Println("----- Top Processes by CPU Usage -----")
		topProcessesByCPU(processes, 5) // Top 5 processes
		fmt.Println("----- Top Processes by Memory Usage -----")
		topProcessesByMemory(processes, 5) // Top 5 processes
	}

	// System Uptime
	uptime, err := host.Uptime()
	if err != nil {
		log.Printf("Error fetching system uptime: %v\n", err)
	} else {
		fmt.Printf("System Uptime: %v hours\n", uptime/3600)
	}

	// Temperature Sensors (if supported)
	//sensorsInfo, err := host.SensorsTemperatures()
	//if err != nil {
	//	fmt.Println("Temperature monitoring not supported on this system or error fetching data.")
	//} else {
	//	if len(sensorsInfo) == 0 {
	//		fmt.Println("No temperature sensors found.")
	//	} else {
	//		for _, sensor := range sensorsInfo {
	//			fmt.Printf("Sensor %s Temperature: %.2f°C\n", sensor.SensorKey, sensor.Temperature)
	//		}
	//	}
	//}

	// System Info
	hostInfo, err := host.Info()
	if err != nil {
		log.Printf("Error fetching host info: %v\n", err)
	} else {
		fmt.Printf("Hostname: %s\n", hostInfo.Hostname)
		fmt.Printf("OS: %s %s\n", hostInfo.OS, hostInfo.PlatformVersion)
		fmt.Printf("Kernel Version: %s\n", hostInfo.KernelVersion)
		fmt.Printf("Boot Time: %s\n", time.Unix(int64(hostInfo.BootTime), 0).Format(time.RFC1123))
	}

	fmt.Println("---------------------------\n")
}

// Helper function to display top N processes by CPU usage
func topProcessesByCPU(processes []*process.Process, topN int) {
	type procInfo struct {
		Pid    int32
		Name   string
		CPU    float64
		Memory float32
		Exe    string
	}

	var procList []procInfo

	for _, proc := range processes {
		cpuPercent, err := proc.CPUPercent()
		if err != nil {
			continue
		}
		if cpuPercent < 0.1 { // Filter out negligible CPU usage
			continue
		}
		name, err := proc.Name()
		if err != nil {
			name = "Unknown"
		}
		procList = append(procList, procInfo{
			Pid:    proc.Pid,
			Name:   name,
			CPU:    cpuPercent,
			Memory: 0, // Placeholder, can be filled if needed
		})
	}

	// Sort the processes by CPU usage descending
	sort.Slice(procList, func(i, j int) bool {
		return procList[i].CPU > procList[j].CPU
	})

	// Display top N processes
	for i, proc := range procList {
		if i >= topN {
			break
		}
		fmt.Printf("PID: %d, Name: %s, CPU Usage: %.2f%%\n", proc.Pid, proc.Name, proc.CPU)
	}
}

// Helper function to display top N processes by Memory usage
func topProcessesByMemory(processes []*process.Process, topN int) {
	type procInfo struct {
		Pid    int32
		Name   string
		CPU    float64
		Memory float32
		Exe    string
	}

	var procList []procInfo

	for _, proc := range processes {
		memPercent, err := proc.MemoryPercent()
		if err != nil {
			continue
		}
		if memPercent < 0.1 { // Filter out negligible memory usage
			continue
		}
		name, err := proc.Name()
		if err != nil {
			name = "Unknown"
		}
		procList = append(procList, procInfo{
			Pid:    proc.Pid,
			Name:   name,
			CPU:    0, // Placeholder, can be filled if needed
			Memory: memPercent,
		})
	}

	// Sort the processes by Memory usage descending
	sort.Slice(procList, func(i, j int) bool {
		return procList[i].Memory > procList[j].Memory
	})

	// Display top N processes
	for i, proc := range procList {
		if i >= topN {
			break
		}
		fmt.Printf("PID: %d, Name: %s, Memory Usage: %.2f%%\n", proc.Pid, proc.Name, proc.Memory)
	}
}
