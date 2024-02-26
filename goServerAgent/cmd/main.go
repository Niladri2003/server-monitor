package main

import (
	"fmt"
	"github.com/Niladri2003/server-monitor/goServerAgent/mem"
)

func main() {

	data, _ := memory.GetMemoryInfo()
	fmt.Println("Total:", data.Total, "Available:", data.Available, "Used:", data.Used, "Free:", data.Used, "Free:", data.Free, "Cached:", data.Cache, "Buffers:", data.Buffers, "UsedPercent:", data.UsedPercent)
	//fmt.Println(string(memInfo))
}
