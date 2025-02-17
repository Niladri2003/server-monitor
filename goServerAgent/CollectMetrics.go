package goServerAgent

import (
	"fmt"
	"github.com/Niladri2003/server-monitor/goServerAgent/DataModel"
	cpu2 "github.com/Niladri2003/server-monitor/goServerAgent/cpustats"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/load"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
	"github.com/shirou/gopsutil/v4/process"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

// SystemMetrics struct holds all the system metrics to be sent via Kafka
type SystemMetrics struct {
	Timestamp    time.Time          `json:"timestamp"`
	Memory       MemoryInfo         `json:"memory"`
	Swap         SwapMemoryInfo     `json:"swap"`
	CPU          CPUInfo            `json:"cpustats"`
	LoadAverages LoadAverages       `json:"load_averages"`
	Disk         DiskInfo           `json:"disk"`
	Network      []NetworkInfo      `json:"network"`
	TopProcesses []ProcessInfo      `json:"top_processes"`
	SystemInfo   HostInfo           `json:"system_info"`
	CpuStats     DataModel.CPUStats `json:"cpu_stats"`
}

// MemoryInfo holds the memory details
type MemoryInfo struct {
	Total       float64 `json:"total_gb"`
	Available   float64 `json:"available_gb"`
	Used        float64 `json:"used_gb"`
	Free        float64 `json:"free_gb"`
	UsedPercent float64 `json:"used_percent"`
}

// SwapMemoryInfo holds the swap memory details
type SwapMemoryInfo struct {
	Total       float64 `json:"total_gb"`
	Used        float64 `json:"used_gb"`
	Free        float64 `json:"free_gb"`
	UsedPercent float64 `json:"used_percent"`
}

// CPUInfo holds CPU details
type CPUInfo struct {
	Model        string             `json:"model"`
	Cores        int                `json:"cores"`
	UsagePerCore []PerCoreUsage     `json:"usage_per_core_percent"`
	Times        []CPUTimesInfo     `json:"times"`
	Temperatures []CPUTemperature   `json:"temperatures"`     // New field for temperatures
	Context      CPUContextSwitches `json:"context_switches"` // New field for context switches
	Interrupts   CPUInterrupts      `json:"interrupts"`
}
type CPUTemperature struct {
	Core        int     `json:"core"`
	Temperature float64 `json:"temperature_celsius"`
}

type CPUContextSwitches struct {
	Voluntary   uint64 `json:"voluntary"`
	Involuntary uint64 `json:"involuntary"`
}

type CPUInterrupts struct {
	Interrupts uint64 `json:"interrupts"`
	SoftIRQs   uint64 `json:"soft_irqs"`
}

type PerCoreUsage struct {
	Core      int     `json:"core"`
	CoreID    string  `json:"coreID"`
	CacheSize int32   `json:"cacheSize"`
	Usage     float64 `json:"usage"`
}

// CPUTimesInfo holds individual CPU core times
type CPUTimesInfo struct {
	User   float64 `json:"user_time_sec"`
	System float64 `json:"system_time_sec"`
	Idle   float64 `json:"idle_time_sec"`
	Iowait float64 `json:"iowait_time_sec"`
	Steal  float64 `json:"steal_time_sec"`
}

// LoadAverages holds CPU load averages
type LoadAverages struct {
	Load1  float64 `json:"load1"`
	Load5  float64 `json:"load5"`
	Load15 float64 `json:"load15"`
}

// DiskInfo holds the disk usage details
type DiskInfo struct {
	Total       float64       `json:"total_gb"`
	Free        float64       `json:"free_gb"`
	Used        float64       `json:"used_gb"`
	UsedPercent float64       `json:"used_percent"`
	IOStats     []DiskIOStats `json:"io_stats"`
}

// DiskIOStats holds disk I/O statistics
type DiskIOStats struct {
	Name       string  `json:"name"`
	ReadBytes  float64 `json:"read_bytes_mb"`
	WriteBytes float64 `json:"write_bytes_mb"`
}

// NetworkInfo holds network interface details
type NetworkInfo struct {
	Name        string  `json:"name"`
	BytesSent   float64 `json:"bytes_sent_mb"`
	BytesRecv   float64 `json:"bytes_recv_mb"`
	PacketsSent uint64  `json:"packets_sent"`
	PacketsRecv uint64  `json:"packets_recv"`
	ErrorsIn    uint64  `json:"errors_in"`
	ErrorsOut   uint64  `json:"errors_out"`
	DropsIn     uint64  `json:"drops_in"`
	DropsOut    uint64  `json:"drops_out"`
}

// ProcessInfo holds details of top processes by CPU and memory
type ProcessInfo struct {
	Pid    int32   `json:"pid"`
	Name   string  `json:"name"`
	CPU    float64 `json:"cpu_percent"`
	Memory float32 `json:"memory_percent"`
}

// HostInfo holds basic system information
type HostInfo struct {
	Hostname        string `json:"hostname"`
	OS              string `json:"os"`
	PlatformVersion string `json:"platform_version"`
	KernelVersion   string `json:"kernel_version"`
	UptimeHours     uint64 `json:"uptime_hours"`
	BootTime        string `json:"boot_time"`
}

func CollectMetrics() SystemMetrics {
	metrics := SystemMetrics{
		Timestamp: time.Now(),
	}

	// Memory Info
	memInfo, err := mem.VirtualMemory()
	if err == nil {
		metrics.Memory = MemoryInfo{
			Total:       float64(memInfo.Total) / 1024 / 1024 / 1024,
			Available:   float64(memInfo.Available) / 1024 / 1024 / 1024,
			Used:        float64(memInfo.Used) / 1024 / 1024 / 1024,
			Free:        float64(memInfo.Free) / 1024 / 1024 / 1024,
			UsedPercent: memInfo.UsedPercent,
		}
	}

	// Swap Memory Info
	swapInfo, err := mem.SwapMemory()
	if err == nil {
		metrics.Swap = SwapMemoryInfo{
			Total:       float64(swapInfo.Total) / 1024 / 1024 / 1024,
			Used:        float64(swapInfo.Used) / 1024 / 1024 / 1024,
			Free:        float64(swapInfo.Free) / 1024 / 1024 / 1024,
			UsedPercent: swapInfo.UsedPercent,
		}
	}
	data, err := cpu2.GetCPUStats()
	fmt.Println(data)
	metrics.CpuStats = data
	// Collect CPU Info (includes static frequency)
	cpuInfo, err := cpu.Info()
	if err != nil {
		log.Fatalf("Failed to get CPU info: %v", err)
	}
	var totalcore int
	for _, cpu := range cpuInfo {
		totalcore += int(cpu.Cores)
	}
	fmt.Println("CORES", totalcore)
	//fmt.Println("CPU Core", cpuInfo[0], cpuInfo[0].CacheSize)
	if len(cpuInfo) > 0 {
		metrics.CPU.Model = cpuInfo[0].ModelName
		metrics.CPU.Cores = totalcore

		// Set static frequency for each core (logical CPU)
		for i := 0; i < len(cpuInfo); i++ {
			metrics.CPU.UsagePerCore = append(metrics.CPU.UsagePerCore, PerCoreUsage{
				Core:      i,
				CoreID:    cpuInfo[i].CoreID,
				CacheSize: cpuInfo[i].CacheSize,
				Usage:     cpuInfo[i].Mhz, // Base frequency in MHz
			})
		}
	}

	// Collect CPU Usage per Core (logical CPUs)
	percentPerCore, err := cpu.Percent(0, true)
	if err != nil {
		log.Fatalf("Failed to get CPU usage per core: %v", err)
	}

	// Update usage percentage for each core
	for i, percent := range percentPerCore {
		if i < len(metrics.CPU.UsagePerCore) {
			metrics.CPU.UsagePerCore[i].Usage = percent
		}
	}

	// Collect CPU Times (per core)
	cpuTimes, err := cpu.Times(true)
	if err != nil {
		log.Fatalf("Failed to get CPU times: %v", err)
	}
	for _, times := range cpuTimes {
		metrics.CPU.Times = append(metrics.CPU.Times, CPUTimesInfo{
			User:   times.User,
			System: times.System,
			Idle:   times.Idle,
			Iowait: times.Iowait,
			Steal:  times.Steal, // Collect steal time
		})
	}

	// Collect CPU Temperature (if available, gopsutil doesn't provide cross-platform support for this directly)
	// You can get temperature using other OS specific tools or libraries like lm-sensors on Linux.

	// Collect context switches (can be done via host.Info())
	hostInfo, err := host.Info()
	if err != nil {
		log.Fatalf("Failed to get host info: %v", err)
	}
	metrics.CPU.Context = CPUContextSwitches{
		Voluntary:   hostInfo.Uptime, // Example, replace with actual data if available
		Involuntary: hostInfo.Uptime, // Example, replace with actual data if available
	}
	//fmt.Println(hostInfo)

	// Interrupts collection might need other packages as gopsutil doesn't have native support.
	// Replace these lines if you find the right tool/library for your platform.
	metrics.CPU.Interrupts = CPUInterrupts{
		Interrupts: 0, // Placeholder, replace with actual data
		SoftIRQs:   0, // Placeholder, replace with actual data
	}

	// Display the final collected metrics
	//fmt.Printf("Collected Metrics: %+v\n", metrics)
	// CPU Load Averages
	loadAvg, err := load.Avg()
	if err == nil {
		metrics.LoadAverages = LoadAverages{
			Load1:  loadAvg.Load1,
			Load5:  loadAvg.Load5,
			Load15: loadAvg.Load15,
		}
	}

	// Disk Usage
	diskInfo, err := disk.Usage("/")
	if err == nil {
		metrics.Disk = DiskInfo{
			Total:       float64(diskInfo.Total) / 1024 / 1024 / 1024,
			Free:        float64(diskInfo.Free) / 1024 / 1024 / 1024,
			Used:        float64(diskInfo.Used) / 1024 / 1024 / 1024,
			UsedPercent: diskInfo.UsedPercent,
		}
	}

	// Network Info
	netInfo, err := net.IOCounters(true)
	if err == nil {
		for _, nic := range netInfo {
			metrics.Network = append(metrics.Network, NetworkInfo{
				Name:        nic.Name,
				BytesSent:   float64(nic.BytesSent) / 1024 / 1024,
				BytesRecv:   float64(nic.BytesRecv) / 1024 / 1024,
				PacketsSent: nic.PacketsSent,
				PacketsRecv: nic.PacketsRecv,
				ErrorsIn:    nic.Errin,
				ErrorsOut:   nic.Errout,
				DropsIn:     nic.Dropin,
				DropsOut:    nic.Dropout,
			})
		}
	}

	// System Info
	hostInfo, err = host.Info()
	if err == nil {
		metrics.SystemInfo = HostInfo{
			Hostname:        hostInfo.Hostname,
			OS:              hostInfo.OS,
			PlatformVersion: hostInfo.PlatformVersion,
			KernelVersion:   hostInfo.KernelVersion,
			UptimeHours:     hostInfo.Uptime / 3600,
			BootTime:        time.Unix(int64(hostInfo.BootTime), 0).Format(time.RFC1123),
		}
	}
	//fmt.Println(metrics)
	return metrics
}

// Helper function to display top N processes by CPU usage
func TopProcessesByCPU(processes []*process.Process, topN int) []ProcessInfo {

	var procList []ProcessInfo

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
		procList = append(procList, ProcessInfo{
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

	// If there are fewer than topN processes, return all
	if len(procList) < topN {
		return procList
	}

	// Return the top N processes
	return procList[:topN]
}

// Helper function to display top N processes by Memory usage
func TopProcessesByMemory(processes []*process.Process, topN int) []ProcessInfo {

	var procList []ProcessInfo

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
		procList = append(procList, ProcessInfo{
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

	if len(procList) < topN {
		return procList
	}
	return procList[:topN]
}

func readSystemTemperatures() {
	thermalBasePath := "/sys/class/thermal/"

	thermalZones, err := filepath.Glob(thermalBasePath + "thermal_zone*")
	if err != nil {
		fmt.Printf("Error reading thermal zones: %v\n", err)
		return
	}

	for _, zonePath := range thermalZones {
		tempFile := filepath.Join(zonePath, "temp")
		typeFile := filepath.Join(zonePath, "type")

		// Read the temperature
		tempContent, err := ioutil.ReadFile(tempFile)
		if err != nil {
			fmt.Printf("Error reading temperature for %s: %v\n", zonePath, err)
			continue
		}
		temp, err := strconv.ParseFloat(strings.TrimSpace(string(tempContent)), 64)
		if err != nil {
			fmt.Printf("Error parsing temperature for %s: %v\n", zonePath, err)
			continue
		}

		// Read the type
		typeContent, err := ioutil.ReadFile(typeFile)
		if err != nil {
			fmt.Printf("Error reading type for %s: %v\n", zonePath, err)
			continue
		}
		sensorType := strings.TrimSpace(string(typeContent))

		fmt.Printf("Temperature for %s (%s): %.2f°C\n", zonePath, sensorType, temp/1000)
	}
}
