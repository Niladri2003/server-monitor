package memory

//paramerters to be measured 1.Total 2.Avilable 3.Used

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetMemoryInfo() (*Memorymet, error) {
	memInfo, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return nil, err
	}
	minfoReturn := &Memorymet{}
	memoryavailable := false
	//parse the content of /proc/meminfo file
	lines := strings.Split(string(memInfo), "\n")
	for _, line := range lines {
		fields := strings.Split(line, ":")
		if len(fields) != 2 {
			continue
		}
		key := strings.TrimSpace(fields[0])
		value := strings.TrimSpace(fields[1])
		value = strings.Replace(value, " kB", "", -1)
		fmt.Println(key, ":", value)
		switch key {
		//Total
		case "MemTotal":
			t, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return minfoReturn, err
			}
			minfoReturn.Total = t * 1024
			//Free
		case "MemFree":
			t, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return minfoReturn, err
			}
			minfoReturn.Free = t * 1024
			//Avilable
		case "MemAvailable":
			t, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return minfoReturn, err
			}
			memoryavailable = true
			minfoReturn.Available = t * 1024
			//Buffers
		case "Buffers":
			t, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return minfoReturn, err
			}
			minfoReturn.Buffers = t * 1024
		case "Cached":
			t, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return minfoReturn, err
			}
			minfoReturn.Cache = t * 1024
		}
	}
	if !memoryavailable {
		minfoReturn.Available = minfoReturn.Cache + minfoReturn.Free

	}
	minfoReturn.Used = minfoReturn.Total - minfoReturn.Buffers - minfoReturn.Cache
	minfoReturn.UsedPercent = float64(minfoReturn.Used) / float64(minfoReturn.Total) * 100
	//fmt.Println("Lines=>", lines)
	return minfoReturn, nil
}
