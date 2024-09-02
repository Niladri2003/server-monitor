package memory

//paramerters to be measured 1.Total 2.Avilable 3.Used

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// GetMemoryInfo captures the system memory metrics and returns them.
func GetMemoryInfo() (*Memorymet, error) {
	memInfo, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return nil, err
	}

	minfoReturn := &Memorymet{}
	lines := strings.Split(string(memInfo), "\n")

	for _, line := range lines {
		fields := strings.Split(line, ":")
		if len(fields) != 2 {
			continue
		}
		key := strings.TrimSpace(fields[0])
		value := strings.TrimSpace(fields[1])
		value = strings.Replace(value, " kB", "", -1)

		t, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return minfoReturn, err
		}

		switch key {
		case "MemTotal":
			minfoReturn.Total = t * 1024
		case "MemFree":
			minfoReturn.Free = t * 1024
		case "MemAvailable":
			minfoReturn.Available = t * 1024
		case "Buffers":
			minfoReturn.Buffers = t * 1024
		case "Cached":
			minfoReturn.Cached = t * 1024
		case "SwapTotal":
			minfoReturn.SwapTotal = t * 1024
		case "SwapFree":
			minfoReturn.SwapFree = t * 1024
		}
	}

	// Calculate used memory
	minfoReturn.Used = minfoReturn.Total - minfoReturn.Free - minfoReturn.Buffers - minfoReturn.Cached
	minfoReturn.UsedPercent = float64(minfoReturn.Used) / float64(minfoReturn.Total) * 100

	// Add a timestamp to the metrics
	minfoReturn.Timestamp = time.Now().Unix()

	return minfoReturn, nil
}
