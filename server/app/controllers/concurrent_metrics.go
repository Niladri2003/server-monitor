package controllers

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"net/http"
	"os"
	"sync"
)

func GetAllMetrics(c *fiber.Ctx, client influxdb2.Client) error {
	serverID := c.Query("server_id")
	start := c.Query("start", "-1h") // Default to -1h if not provided
	stop := c.Query("stop", "now()") // Default to now() if not provided

	// Ensure that serverID is provided
	if serverID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "server_id must be provided",
		})
	}

	var wg sync.WaitGroup
	wg.Add(9)

	// Channels for each metric
	memChan := make(chan fiber.Map)
	swapChan := make(chan fiber.Map)
	cpuChan := make(chan fiber.Map)
	top5CpuChan := make(chan fiber.Map)
	top5MemChan := make(chan fiber.Map)
	hostInfoChan := make(chan fiber.Map)
	diskChan := make(chan fiber.Map)
	networkChan := make(chan fiber.Map)

	// Query functions running in parallel
	go queryMemoryUsage(c, client, serverID, start, stop, memChan, &wg)
	go querySwapMemoryUsage(c, client, serverID, start, stop, swapChan, &wg)
	go queryCpuUsage(c, client, serverID, start, stop, cpuChan, &wg)
	go queryTop5ProcessByCpu(c, client, serverID, start, stop, top5CpuChan, &wg)
	go queryTop5ProcessByMemory(c, client, serverID, start, stop, top5MemChan, &wg)
	go queryHostInfo(c, client, serverID, start, stop, hostInfoChan, &wg)
	go queryDiskUsage(c, client, serverID, start, stop, diskChan, &wg)
	go queryNetworkStats(c, client, serverID, start, stop, networkChan, &wg)

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(memChan)
		close(swapChan)
		close(cpuChan)
		close(top5CpuChan)
		close(top5MemChan)
		close(hostInfoChan)
		close(diskChan)
		close(networkChan)
	}()

	// Combine the results from all channels
	response := fiber.Map{
		"memory_usage":        <-memChan,
		"swap_memory_usage":   <-swapChan,
		"cpu_usage":           <-cpuChan,
		"top_5_cpu_processes": <-top5CpuChan,
		"top_5_mem_processes": <-top5MemChan,
		"host_info":           <-hostInfoChan,
		"disk_usage":          <-diskChan,
		"network_stats":       <-networkChan,
	}

	return c.Status(http.StatusOK).JSON(response)
}

// Individual query functions
func queryMemoryUsage(c *fiber.Ctx, client influxdb2.Client, serverID, start, stop string, ch chan fiber.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	query := `from(bucket: "` + os.Getenv("INFLUXDB_BUCKET") + `")
              |> range(start: ` + start + `, stop: ` + stop + `)
              |> filter(fn: (r) => r["_measurement"] == "memory")
              |> filter(fn: (r) => r["server_id"] == "` + serverID + `")
              |> filter(fn: (r) => r._field == "free_gb" or r._field == "total_gb" or r._field == "used_percent" or r._field == "used_gb")
              |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`
	queryInfluxDBConcurrent(c, client, query, "Memory", ch)

}

func querySwapMemoryUsage(c *fiber.Ctx, client influxdb2.Client, serverID, start, stop string, ch chan fiber.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	query := `from(bucket: "` + os.Getenv("INFLUXDB_BUCKET") + `")
              |> range(start: ` + start + `, stop: ` + stop + `)
              |> filter(fn: (r) => r["_measurement"] == "swap_memory")
              |> filter(fn: (r) => r["server_id"] == "` + serverID + `")
              |> filter(fn: (r) => r._field == "free_gb" or r._field == "total_gb" or r._field == "used_percent" or r._field == "used_gb")
              |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`
	queryInfluxDBConcurrent(c, client, query, "Swap", ch)

}

func queryCpuUsage(c *fiber.Ctx, client influxdb2.Client, serverID, start, stop string, ch chan fiber.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	query := `from(bucket: "` + os.Getenv("INFLUXDB_BUCKET") + `")
              |> range(start: ` + start + `, stop: ` + stop + `)
              |> filter(fn: (r) => r["_measurement"] == "cpu_metrics")
              |> filter(fn: (r) => r["server_id"] == "` + serverID + `")
              |> filter(fn: (r) => r._field == "core" or r._field == "idle_time_sec" or r._field == "iowait_time_sec" or r._field == "system_time_sec" or r._field == "usage_per_core_percent" or r._field == "user_time_sec" or r._field == "model")
              |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`
	queryInfluxDBConcurrent(c, client, query, "Cpu", ch)

}
func queryTop5ProcessByCpu(c *fiber.Ctx, client influxdb2.Client, serverID, start, stop string, ch chan fiber.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	query := `from(bucket: "` + os.Getenv("INFLUXDB_BUCKET") + `")
              |> range(start: ` + start + `, stop: ` + stop + `)
              |> filter(fn: (r) => r["_measurement"] == "top_processes_by_cpu")
              |> filter(fn: (r) => r["server_id"] == "` + serverID + `")
              |> filter(fn: (r) => r._field == "name" or r._field == "cpu_percent" or r._field == "memory_percent" or r._field == "pid")
              |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`
	queryInfluxDBConcurrent(c, client, query, "ProcessCpu", ch)

}

func queryTop5ProcessByMemory(c *fiber.Ctx, client influxdb2.Client, serverID, start, stop string, ch chan fiber.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	query := `from(bucket: "` + os.Getenv("INFLUXDB_BUCKET") + `")
              |> range(start: ` + start + `, stop: ` + stop + `)
              |> filter(fn: (r) => r["_measurement"] == "top_processes_by_memory")
              |> filter(fn: (r) => r["server_id"] == "` + serverID + `")
              |> filter(fn: (r) => r._field == "name" or r._field == "cpu_percent" or r._field == "memory_percent" or r._field == "pid")
              |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`
	queryInfluxDBConcurrent(c, client, query, "ProcessMemory", ch)

}

func queryHostInfo(c *fiber.Ctx, client influxdb2.Client, serverID, start, stop string, ch chan fiber.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	query := `from(bucket: "` + os.Getenv("INFLUXDB_BUCKET") + `")
              |> range(start: ` + start + `, stop: ` + stop + `)
              |> filter(fn: (r) => r["_measurement"] == "host_info")
              |> filter(fn: (r) => r["server_id"] == "` + serverID + `")
              |> filter(fn: (r) => r._field == "boot_time" or r._field == "hostname" 
	          or r._field == "kernel_version" or r._field == "os" or r._field == "platform_version" or r._field == "uptime_hours")
              |> last()
              |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`
	queryInfluxDBConcurrent(c, client, query, "HostInfo", ch)

}

func queryDiskUsage(c *fiber.Ctx, client influxdb2.Client, serverID, start, stop string, ch chan fiber.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	query := `from(bucket: "` + os.Getenv("INFLUXDB_BUCKET") + `")
              |> range(start: ` + start + `, stop: ` + stop + `)
              |> filter(fn: (r) => r["_measurement"] == "disk")
              |> filter(fn: (r) => r["server_id"] == "` + serverID + `")
             |> filter(fn: (r) => r._field == "free_gb" or r._field == "total_gb" or r._field == "used_percent" or r._field == "used_gb")
              |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`
	queryInfluxDBConcurrent(c, client, query, "Disk", ch)
}
func queryNetworkStats(c *fiber.Ctx, client influxdb2.Client, serverID, start, stop string, ch chan fiber.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	query := `from(bucket: "` + os.Getenv("INFLUXDB_BUCKET") + `")
              |> range(start: ` + start + `, stop: ` + stop + `)
              |> filter(fn: (r) => r["_measurement"] == "network")
              |> filter(fn: (r) => r["server_id"] == "` + serverID + `")
              |> filter(fn: (r) => r._field == "bytes_recv_mb" or r._field == "bytes_sent_mb" 
	                           or r._field == "drops_in" or r._field == "drops_out" 
	                           or r._field == "errors_in" or r._field == "errors_out" 
	                           or r._field == "interface_name" or r._field == "packets_recv" 
	                           or r._field == "packets_sent")
              |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`
	queryInfluxDBConcurrent(c, client, query, "Network", ch)

}

func queryInfluxDBConcurrent(c *fiber.Ctx, client influxdb2.Client, query, measurement string, ch chan fiber.Map) {
	// Initialize the QueryAPI for InfluxDB
	queryAPI := client.QueryAPI(os.Getenv("INFLUXDB_ORG"))
	fmt.Println(measurement, query)

	// Execute the query
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		ch <- fiber.Map{"error": fmt.Sprintf("Query error: %s", err.Error())}
		return
	}
	defer result.Close()

	// Initialize a slice to hold the result data
	var data []map[string]interface{}

	// Process the query result based on the measurement type
	switch measurement {

	case "Disk":
		fmt.Println("Disk")
		for result.Next() {
			record := result.Record()
			fmt.Println(record)
			// The fields are now columns, so we'll extract them individually
			row := map[string]interface{}{
				"time":     record.Time(),
				"free_gb":  record.ValueByKey("free_gb"),
				"total_gb": record.ValueByKey("total_gb"),
				"used_gb":  record.ValueByKey("used_gb"),
				"used_pct": record.ValueByKey("used_percent"),
			}
			data = append(data, row)
		}
	case "Swap":
		for result.Next() {
			record := result.Record()
			//fmt.Println(record)
			// The fields are now columns, so we'll extract them individually
			row := map[string]interface{}{
				"time":     record.Time(),
				"free_gb":  record.ValueByKey("free_gb"),
				"total_gb": record.ValueByKey("total_gb"),
				"used_gb":  record.ValueByKey("used_gb"),
				"used_pct": record.ValueByKey("used_percent"),
			}
			data = append(data, row)
		}
	case "Memory":
		for result.Next() {
			record := result.Record()
			//fmt.Println(record)
			// The fields are now columns, so we'll extract them individually
			row := map[string]interface{}{
				"time":     record.Time(),
				"free_gb":  record.ValueByKey("free_gb"),
				"total_gb": record.ValueByKey("total_gb"),
				"used_gb":  record.ValueByKey("used_gb"),
				"used_pct": record.ValueByKey("used_percent"),
			}
			data = append(data, row)
		}
	case "Network":
		for result.Next() {
			record := result.Record()
			fmt.Println("data", record)
			// The fields are now columns, so we'll extract them individually
			row := map[string]interface{}{
				"time":           record.Time(),
				"bytes_recv_mb":  record.ValueByKey("bytes_recv_mb"),
				"bytes_sent_mb":  record.ValueByKey("bytes_sent_mb"),
				"drops_in":       record.ValueByKey("drops_in"),
				"drops_out":      record.ValueByKey("drops_out"),
				"errors_in":      record.ValueByKey("errors_in"),
				"errors_out":     record.ValueByKey("errors_out"),
				"interface_name": record.ValueByKey("interface_name"),
				"packets_recv":   record.ValueByKey("packets_recv"),
				"packets_sent":   record.ValueByKey("packets_sent"),
			}
			data = append(data, row)
		}
	case "HostInfo":
		for result.Next() {
			record := result.Record()
			//fmt.Println("data", record)
			// The fields are now columns, so we'll extract them individually
			row := map[string]interface{}{
				"time":             record.Time(),
				"hostname":         record.ValueByKey("hostname"),
				"kernel_version":   record.ValueByKey("kernel_version"),
				"os":               record.ValueByKey("os"),
				"platform_version": record.ValueByKey("platform_version"),
				"uptime_hours":     record.ValueByKey("uptime_hours"),
			}
			data = append(data, row)
		}
	case "Cpu":
		for result.Next() {
			record := result.Record()
			//fmt.Println("data", record)
			// The fields are now columns, so we'll extract them individually
			row := map[string]interface{}{
				"time":                   record.Time(),
				"Model":                  record.ValueByKey("model"),
				"core":                   record.ValueByKey("core"),
				"usage_per_core_percent": record.ValueByKey("usage_per_core_percent"),
				"system_time_sec":        record.ValueByKey("system_time_sec"),
				"iowait_time_sec":        record.ValueByKey("iowait_time_sec"),
				"idle_time_sec":          record.ValueByKey("idle_time_sec"),
				"user_time_sec":          record.ValueByKey("user_time_sec"),
			}
			data = append(data, row)
		}
	case "ProcessMemory":
		for result.Next() {
			record := result.Record()
			//fmt.Println("data", record)
			// The fields are now columns, so we'll extract them individually
			row := map[string]interface{}{
				"time":           record.Time(),
				"name":           record.ValueByKey("name"),
				"cpu_percent":    record.ValueByKey("cpu_percent"),
				"memory_percent": record.ValueByKey("memory_percent"),
				"pid":            record.ValueByKey("pid"),
			}
			data = append(data, row)
		}
	case "ProcessCpu":
		for result.Next() {
			record := result.Record()
			//fmt.Println("data", record)
			// The fields are now columns, so we'll extract them individually
			row := map[string]interface{}{
				"time":           record.Time(),
				"name":           record.ValueByKey("name"),
				"cpu_percent":    record.ValueByKey("cpu_percent"),
				"memory_percent": record.ValueByKey("memory_percent"),
				"pid":            record.ValueByKey("pid"),
			}
			data = append(data, row)
		}
	}

	if result.Err() != nil {
		ch <- fiber.Map{"error": fmt.Sprintf("Query processing error: %s", result.Err().Error())}
		return
	}

	// Send the final result through the channel
	ch <- fiber.Map{"data": data}
}
