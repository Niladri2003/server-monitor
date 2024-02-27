package main

import (
	"fmt"
	memory "github.com/Niladri2003/server-monitor/goServerAgent/mem"
	"os"
)

func main() {

	data, _ := memory.GetMemoryInfo()
	fmt.Println("Total:", data.Total, "Available:", data.Available, "Used:", data.Used, "Free:", data.Used, "Free:", data.Free, "Cached:", data.Cache, "Buffers:", data.Buffers, "UsedPercent:", data.UsedPercent)
	//fmt.Println(string(memInfo))
	cpu, _ := os.ReadFile("/proc/cpuinfo")
	fmt.Println(string(cpu))
	stat, _ := os.ReadFile("/proc/stat")
	fmt.Println("STAT:", string(stat))

}
