package cpustats

import (
	"bufio"
	"fmt"
	"github.com/Niladri2003/server-monitor/goServerAgent/DataModel"
	"os"
	"strconv"
	"strings"
)

func GetCPUStats() (DataModel.CPUStats, error) {
	var totalStats DataModel.CPUStats
	totalStats.LogicalCPUs = make(map[string]DataModel.CPUStats)
	totalStats.SoftIrqDetails = make(map[string]uint64)
	file, err := os.Open("/proc/stat")
	if err != nil {
		return totalStats, fmt.Errorf("failed to open /proc/stat: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Collect stats for all CPUs (cpu, cpu0, cpu1, ...)
		if strings.HasPrefix(line, "cpu") {
			fields := strings.Fields(line)

			// Ensure there are enough fields for parsing
			if len(fields) < 11 {
				return totalStats, fmt.Errorf("invalid CPU line in /proc/stat: %s", line)
			}

			// Create a CPUStats struct for individual CPU stats
			var cpuStats DataModel.CPUStats
			var err error
			cpuStats.User, err = parseUint(fields[1])
			if err != nil {
				return totalStats, fmt.Errorf("failed to parse User field: %w", err)
			}
			cpuStats.Nice, err = parseUint(fields[2])
			if err != nil {
				return totalStats, err
			}
			cpuStats.System, err = parseUint(fields[3])
			if err != nil {
				return totalStats, err
			}
			cpuStats.Idle, err = parseUint(fields[4])
			if err != nil {
				return totalStats, err
			}
			cpuStats.Iowait, err = parseUint(fields[5])
			if err != nil {
				return totalStats, err
			}
			cpuStats.Irq, err = parseUint(fields[6])
			if err != nil {
				return totalStats, err
			}
			cpuStats.Softirq, err = parseUint(fields[7])
			if err != nil {
				return totalStats, err
			}
			cpuStats.Steal, err = parseUint(fields[8])
			if err != nil {
				return totalStats, err
			}
			cpuStats.Guest, err = parseUint(fields[9])
			if err != nil {
				return totalStats, err
			}
			cpuStats.GuestNice, err = parseUint(fields[10])
			if err != nil {
				return totalStats, err
			}
			fmt.Println(cpuStats)
			// Initialize the LogicalCPUs map for the individual CPUStats struct
			cpuStats.LogicalCPUs = make(map[string]DataModel.CPUStats)
			// Store individual CPU stats based on the prefix (cpu, cpu0, cpu1, ...)
			if len(fields[0]) > 3 { // Check if it's a specific CPU (like cpu0, cpu1, etc.)
				cpuID := fields[0]                       // e.g., cpu0, cpu1, etc.
				totalStats.LogicalCPUs[cpuID] = cpuStats // Store stats in the map
			} else {
				// If it's the total CPU line (just "cpu"), store it
				totalStats = cpuStats
			}
		}

		// Collect context switches (ctxt)
		if strings.HasPrefix(line, "ctxt") {
			fields := strings.Fields(line)
			if len(fields) < 2 {
				return totalStats, fmt.Errorf("invalid ctxt line in /proc/stat: %s", line)
			}
			totalStats.ContextSwitches, _ = parseUint(fields[1])
		}

		// Collect number of processes created (processes)
		if strings.HasPrefix(line, "processes") {
			fields := strings.Fields(line)
			if len(fields) < 2 {
				return totalStats, fmt.Errorf("invalid processes line in /proc/stat: %s", line)
			}
			totalStats.Processes, _ = parseUint(fields[1])
		}

		// Collect softirq statistics
		if strings.HasPrefix(line, "softirq") {
			fields := strings.Fields(line)
			if len(fields) < 11 {
				return totalStats, fmt.Errorf("invalid softirq line in /proc/stat: %s", line)
			}
			totalStats.SoftIrqTotal, _ = parseUint(fields[1])
			totalStats.SoftIrqDetails = make(map[string]uint64)
			softIrqTypes := []string{"hi", "timer", "net_tx", "net_rx", "block", "irq_poll", "tasklet", "sched", "hrtimer", "rcu"}
			for i, irqType := range softIrqTypes {
				totalStats.SoftIrqDetails[irqType], _ = parseUint(fields[i+2])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return totalStats, fmt.Errorf("error reading /proc/stat: %w", err)
	}

	return totalStats, nil
}

// parseUint safely parses a uint64 from a string, returning an error if the parsing fails.
func parseUint(s string) (uint64, error) {
	value, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse %q as uint64: %w", s, err)
	}
	return value, nil
}
