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
	result := queryInfluxDBConcurrent(c, client, query, "Memory")
	ch <- result
}

func querySwapMemoryUsage(c *fiber.Ctx, client influxdb2.Client, serverID, start, stop string, ch chan fiber.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	query := `from(bucket: "` + os.Getenv("INFLUXDB_BUCKET") + `")
              |> range(start: ` + start + `, stop: ` + stop + `)
              |> filter(fn: (r) => r["_measurement"] == "swap_memory")
              |> filter(fn: (r) => r["server_id"] == "` + serverID + `")
              |> filter(fn: (r) => r._field == "free_gb" or r._field == "total_gb" or r._field == "used_percent" or r._field == "used_gb")
              |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`
	result := queryInfluxDBConcurrent(c, client, query, "Swap")
	ch <- result
}

func queryCpuUsage(c *fiber.Ctx, client influxdb2.Client, serverID, start, stop string, ch chan fiber.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	query := `from(bucket: "` + os.Getenv("INFLUXDB_BUCKET") + `")
              |> range(start: ` + start + `, stop: ` + stop + `)
              |> filter(fn: (r) => r["_measurement"] == "cpu_metrics")
              |> filter(fn: (r) => r["server_id"] == "` + serverID + `")
              |> filter(fn: (r) => r._field == "core" or r._field == "idle_time_sec" or r._field == "iowait_time_sec" or r._field == "system_time_sec" or r._field == "usage_per_core_percent" or r._field == "user_time_sec" or r._field == "model")
              |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`
	result := queryInfluxDBConcurrent(c, client, query, "Cpu")
	ch <- result
}
func queryTop5ProcessByCpu(c *fiber.Ctx, client influxdb2.Client, serverID, start, stop string, ch chan fiber.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	query := `from(bucket: "` + os.Getenv("INFLUXDB_BUCKET") + `")
              |> range(start: ` + start + `, stop: ` + stop + `)
              |> filter(fn: (r) => r["_measurement"] == "cpu_metrics")
              |> filter(fn: (r) => r["server_id"] == "` + serverID + `")
              |> filter(fn: (r) => r._field == "core" or r._field == "idle_time_sec" or r._field == "iowait_time_sec" or r._field == "system_time_sec" or r._field == "usage_per_core_percent" or r._field == "user_time_sec" or r._field == "model")
              |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`
	result := queryInfluxDBConcurrent(c, client, query, "Cpu")
	ch <- result
}

func queryTop5ProcessByMemory(c *fiber.Ctx, client influxdb2.Client, serverID, start, stop string, ch chan fiber.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	query := `from(bucket: "` + os.Getenv("INFLUXDB_BUCKET") + `")
              |> range(start: ` + start + `, stop: ` + stop + `)
              |> filter(fn: (r) => r["_measurement"] == "cpu_metrics")
              |> filter(fn: (r) => r["server_id"] == "` + serverID + `")
              |> filter(fn: (r) => r._field == "core" or r._field == "idle_time_sec" or r._field == "iowait_time_sec" or r._field == "system_time_sec" or r._field == "usage_per_core_percent" or r._field == "user_time_sec" or r._field == "model")
              |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`
	result := queryInfluxDBConcurrent(c, client, query, "Cpu")
	ch <- result
}

func queryHostInfo(c *fiber.Ctx, client influxdb2.Client, serverID, start, stop string, ch chan fiber.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	query := `from(bucket: "` + os.Getenv("INFLUXDB_BUCKET") + `")
              |> range(start: ` + start + `, stop: ` + stop + `)
              |> filter(fn: (r) => r["_measurement"] == "cpu_metrics")
              |> filter(fn: (r) => r["server_id"] == "` + serverID + `")
              |> filter(fn: (r) => r._field == "core" or r._field == "idle_time_sec" or r._field == "iowait_time_sec" or r._field == "system_time_sec" or r._field == "usage_per_core_percent" or r._field == "user_time_sec" or r._field == "model")
              |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`
	result := queryInfluxDBConcurrent(c, client, query, "Cpu")
	ch <- result
}

func queryDiskUsage(c *fiber.Ctx, client influxdb2.Client, serverID, start, stop string, ch chan fiber.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	query := `from(bucket: "` + os.Getenv("INFLUXDB_BUCKET") + `")
              |> range(start: ` + start + `, stop: ` + stop + `)
              |> filter(fn: (r) => r["_measurement"] == "cpu_metrics")
              |> filter(fn: (r) => r["server_id"] == "` + serverID + `")
              |> filter(fn: (r) => r._field == "core" or r._field == "idle_time_sec" or r._field == "iowait_time_sec" or r._field == "system_time_sec" or r._field == "usage_per_core_percent" or r._field == "user_time_sec" or r._field == "model")
              |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`
	result := queryInfluxDBConcurrent(c, client, query, "Cpu")
	ch <- result
}
func queryNetworkStats(c *fiber.Ctx, client influxdb2.Client, serverID, start, stop string, ch chan fiber.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	query := `from(bucket: "` + os.Getenv("INFLUXDB_BUCKET") + `")
              |> range(start: ` + start + `, stop: ` + stop + `)
              |> filter(fn: (r) => r["_measurement"] == "cpu_metrics")
              |> filter(fn: (r) => r["server_id"] == "` + serverID + `")
              |> filter(fn: (r) => r._field == "core" or r._field == "idle_time_sec" or r._field == "iowait_time_sec" or r._field == "system_time_sec" or r._field == "usage_per_core_percent" or r._field == "user_time_sec" or r._field == "model")
              |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`
	result := queryInfluxDBConcurrent(c, client, query, "Cpu")
	ch <- result
}

func queryInfluxDBConcurrent(c *fiber.Ctx, client influxdb2.Client, query, measurement string) fiber.Map {
	// Initialize the QueryAPI for InfluxDB
	queryAPI := client.QueryAPI(os.Getenv("INFLUXDB_ORG"))

	// Execute the query
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		return fiber.Map{"error": fmt.Sprintf("Query error: %s", err.Error())}
	}
	defer result.Close()

	// Initialize a slice to hold the result data
	var data []map[string]interface{}

	// Process the query result based on the measurement type
	for result.Next() {
		record := result.Record()
		row := map[string]interface{}{"time": record.Time()} // Common field

		switch measurement {
		case "Disk", "Swap", "Memory":
			row["free_gb"] = record.ValueByKey("free_gb")
			row["total_gb"] = record.ValueByKey("total_gb")
			row["used_gb"] = record.ValueByKey("used_gb")
			row["used_pct"] = record.ValueByKey("used_percent")

		case "Network":
			row["bytes_recv_mb"] = record.ValueByKey("bytes_recv_mb")
			row["bytes_sent_mb"] = record.ValueByKey("bytes_sent_mb")
			row["drops_in"] = record.ValueByKey("drops_in")
			row["drops_out"] = record.ValueByKey("drops_out")
			row["errors_in"] = record.ValueByKey("errors_in")
			row["errors_out"] = record.ValueByKey("errors_out")
			row["interface_name"] = record.ValueByKey("interface_name")
			row["packets_recv"] = record.ValueByKey("packets_recv")
			row["packets_sent"] = record.ValueByKey("packets_sent")

		case "HostInfo":
			row["hostname"] = record.ValueByKey("hostname")
			row["kernel_version"] = record.ValueByKey("kernel_version")
			row["os"] = record.ValueByKey("os")
			row["platform_version"] = record.ValueByKey("platform_version")
			row["uptime_hours"] = record.ValueByKey("uptime_hours")

		case "Cpu":
			row["Model"] = record.ValueByKey("model")
			row["core"] = record.ValueByKey("core")
			row["usage_per_core_percent"] = record.ValueByKey("usage_per_core_percent")
			row["system_time_sec"] = record.ValueByKey("system_time_sec")
			row["iowait_time_sec"] = record.ValueByKey("iowait_time_sec")
			row["idle_time_sec"] = record.ValueByKey("idle_time_sec")
			row["user_time_sec"] = record.ValueByKey("user_time_sec")

		case "ProcessMemory", "ProcessCpu":
			row["name"] = record.ValueByKey("name")
			row["cpu_percent"] = record.ValueByKey("cpu_percent")
			row["memory_percent"] = record.ValueByKey("memory_percent")
			row["pid"] = record.ValueByKey("pid")
		}
		data = append(data, row)
	}

	// Check if there was an error in the query result
	if result.Err() != nil {
		return fiber.Map{"error": fmt.Sprintf("Query error: %s", result.Err().Error())}
	}

	// Return the data as a fiber.Map
	return fiber.Map{"data": data, "msg": "Data fetched successfully"}
}
