package DataModel

import "time"

type SystemMetrics struct {
	Timestamp    time.Time      `json:"timestamp"`
	Memory       MemoryInfo     `json:"memory"`
	Swap         SwapMemoryInfo `json:"swap"`
	CPU          CPUInfo        `json:"cpustats"`
	LoadAverages LoadAverages   `json:"load_averages"`
	Disk         DiskInfo       `json:"disk"`
	Network      []NetworkInfo  `json:"network"`
	TopProcesses []ProcessInfo  `json:"top_processes"`
	SystemInfo   HostInfo       `json:"system_info"`
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
