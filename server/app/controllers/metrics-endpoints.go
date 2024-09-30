package controllers

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"net/http"
	"os"
)

func CpuLoadAvg(c *fiber.Ctx, client influxdb2.Client) error {
	// Parse the time range from the query parameters
	start := c.Query("start", "-1h") // Default to -6h if not provided
	stop := c.Query("stop", "now()") // Default to now() if not provided
	serverID := c.Query("server_id")
	// Ensure that  serverID  is provided
	if serverID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Either server_id or api_key must be provided",
		})
	}

	query := `from(bucket: "` + os.Getenv("INFLUXDB_BUCKET") + `")
              |> range(start: ` + start + `, stop: ` + stop + `)
              |> filter(fn: (r) => r["_measurement"] == "load_averages")
              |> filter(fn: (r) => r["server_id"] == "` + serverID + `")
              |> filter(fn: (r) => r._field == "load1" or r._field == "load5" or r._field == "load15")
              |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`
	return queryInfluxDB(c, client, query, "CpuLoad")

}
func GetMemoryUsage(c *fiber.Ctx, client influxdb2.Client) error {
	// Parse the time range from the query parameters
	start := c.Query("start", "-1h") // Default to -6h if not provided
	stop := c.Query("stop", "now()") // Default to now() if not provided
	serverID := c.Query("server_id")

	// Ensure that  serverID  is provided
	if serverID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Either server_id or api_key must be provided",
		})
	}
	// Start building the query string
	query := `from(bucket: "` + os.Getenv("INFLUXDB_BUCKET") + `")
              |> range(start: ` + start + `, stop: ` + stop + `)
              |> filter(fn: (r) => r["_measurement"] == "memory")`
	// Add filters for server_id or api_key if provided

	query += ` |> filter(fn: (r) => r["server_id"] == "` + serverID + `")`

	query += `|> filter(fn: (r) => r._field == "free_gb" or r._field == "total_gb"
				or r._field == "used_percent" or r._field == "used_gb")
  	          |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`

	return queryInfluxDB(c, client, query, "Memory")
}
func GetSwapMemoryUsage(c *fiber.Ctx, client influxdb2.Client) error {
	// Parse the time range from the query parameters
	start := c.Query("start", "-1h") // Default to -6h if not provided
	stop := c.Query("stop", "now()") // Default to now() if not provided
	serverID := c.Query("server_id")

	// Ensure that  serverID  is provided
	if serverID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Either server_id or api_key must be provided",
		})
	}
	// Start building the query string
	query := `from(bucket: "` + os.Getenv("INFLUXDB_BUCKET") + `")
              |> range(start: ` + start + `, stop: ` + stop + `)
              |> filter(fn: (r) => r["_measurement"] == "swap_memory")`
	// Add filters for server_id or api_key if provided

	query += ` |> filter(fn: (r) => r["server_id"] == "` + serverID + `")`

	query += `|> filter(fn: (r) => r._field == "free_gb" or r._field == "total_gb" 
				or r._field == "used_percent" or r._field == "used_gb")
  	          |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`

	return queryInfluxDB(c, client, query, "Swap")
}
func GetCpuUsage(c *fiber.Ctx, client influxdb2.Client) error {
	// Parse the time range from the query parameters
	start := c.Query("start", "-1h") // Default to -6h if not provided
	stop := c.Query("stop", "now()") // Default to now() if not provided
	serverID := c.Query("server_id")

	// Ensure that  serverID  is provided
	if serverID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Either server_id or api_key must be provided",
		})
	}
	// Start building the query string
	query := `from(bucket: "` + os.Getenv("INFLUXDB_BUCKET") + `")
              |> range(start: ` + start + `, stop: ` + stop + `)
              |> filter(fn: (r) => r["_measurement"] == "cpu_metrics")`
	// Add filters for server_id or api_key if provided

	query += ` |> filter(fn: (r) => r["server_id"] == "` + serverID + `")`

	query += `|> filter(fn: (r) => r._field == "core" or r._field == "idle_time_sec" 
		      or r._field == "iowait_time_sec" or r._field == "system_time_sec" or 
              r._field == "usage_per_core_percent" or r._field == "user_time_sec" or r._field == "model")
  	          |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`

	return queryInfluxDB(c, client, query, "Cpu")
}
func GetTop5ProcessByCpu(c *fiber.Ctx, client influxdb2.Client) error {
	// Parse the time range from the query parameters
	start := c.Query("start", "-1h") // Default to -6h if not provided
	stop := c.Query("stop", "now()") // Default to now() if not provided
	serverID := c.Query("server_id")

	// Ensure that  serverID  is provided
	if serverID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Either server_id or api_key must be provided",
		})
	}
	// Start building the query string
	query := `from(bucket: "` + os.Getenv("INFLUXDB_BUCKET") + `")
              |> range(start: ` + start + `, stop: ` + stop + `)
              |> filter(fn: (r) => r["_measurement"] == "top_processes_by_cpu")`
	// Add filters for server_id or api_key if provided

	query += ` |> filter(fn: (r) => r["server_id"] == "` + serverID + `")`

	query += `|> filter(fn: (r) => r._field == "name" or r._field == "cpu_percent" or r._field == "memory_percent" or r._field == "pid" )
  	          |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`

	return queryInfluxDB(c, client, query, "ProcessCpu")
}
func GetTop5ProcessByMemory(c *fiber.Ctx, client influxdb2.Client) error {
	// Parse the time range from the query parameters
	start := c.Query("start", "-1h") // Default to -6h if not provided
	stop := c.Query("stop", "now()") // Default to now() if not provided
	serverID := c.Query("server_id")

	// Ensure that  serverID  is provided
	if serverID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Either server_id or api_key must be provided",
		})
	}
	// Start building the query string
	query := `from(bucket: "` + os.Getenv("INFLUXDB_BUCKET") + `")
              |> range(start: ` + start + `, stop: ` + stop + `)
              |> filter(fn: (r) => r["_measurement"] == "top_processes_by_memory")`
	// Add filters for server_id or api_key if provided

	query += ` |> filter(fn: (r) => r["server_id"] == "` + serverID + `")`

	query += `|> filter(fn: (r) => r._field == "name" or r._field == "cpu_percent" or r._field == "memory_percent" or r._field == "pid" )
  	          |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`

	return queryInfluxDB(c, client, query, "ProcessMemory")
}
func GetHostInfo(c *fiber.Ctx, client influxdb2.Client) error {
	serverID := c.Query("server_id")

	// Ensure that  serverID  is provided
	if serverID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "server_id must be provided",
		})
	}
	// Start building the query string
	query := `from(bucket: "` + os.Getenv("INFLUXDB_BUCKET") + `")
              |> range(start: -1d)
              |> filter(fn: (r) => r["_measurement"] == "host_info")`

	// Add filters for server_id or api_key if provided

	query += ` |> filter(fn: (r) => r["server_id"] == "` + serverID + `")`

	query += `|> filter(fn: (r) => r._field == "boot_time" or r._field == "hostname" 
	          or r._field == "kernel_version" or r._field == "os" or r._field == "platform_version" or r._field == "uptime_hours")
              |> last()
  	          |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`

	return queryInfluxDB(c, client, query, "HostInfo")
}
func GetDiskUsage(c *fiber.Ctx, client influxdb2.Client) error {
	// Parse the time range from the query parameters
	start := c.Query("start", "-1h") // Default to -6h if not provided
	stop := c.Query("stop", "now()") // Default to now() if not provided
	serverID := c.Query("server_id")
	fmt.Println("Hello")
	// Ensure that  serverID  is provided
	if serverID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "server_id must be provided",
		})
	}
	// Start building the query string
	query := `from(bucket: "` + os.Getenv("INFLUXDB_BUCKET") + `")
             |> range(start: ` + start + `, stop: ` + stop + `)
              |> filter(fn: (r) => r["_measurement"] == "disk")`

	// Add filters for server_id or api_key if provided

	query += ` |> filter(fn: (r) => r["server_id"] == "` + serverID + `")`

	query += `|> filter(fn: (r) => r._field == "free_gb" or r._field == "total_gb" or r._field == "used_percent" or r._field == "used_gb")
  	          |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`

	return queryInfluxDB(c, client, query, "disk")
}
func GetNetworkStats(c *fiber.Ctx, client influxdb2.Client) error {
	// Parse the time range from the query parameters
	start := c.Query("start", "-1h") // Default to -6h if not provided
	stop := c.Query("stop", "now()")
	serverID := c.Query("server_id")
	fmt.Println(start) // Default to now() if not provided
	// Ensure that  serverID  is provided
	if serverID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Either server_id or api_key must be provided",
		})
	}
	// Start building the query string
	query := `from(bucket: "` + os.Getenv("INFLUXDB_BUCKET") + `")
              |> range(start: ` + start + `, stop: ` + stop + `)
              |> filter(fn: (r) => r["_measurement"] == "network")`

	// Add filters for server_id or api_key if provided

	query += ` |> filter(fn: (r) => r["server_id"] == "` + serverID + `")`

	// Filter on fields (e.g., bytes sent/received, drops, errors, etc.)
	query += ` |> filter(fn: (r) => r._field == "bytes_recv_mb" or r._field == "bytes_sent_mb" 
	                           or r._field == "drops_in" or r._field == "drops_out" 
	                           or r._field == "errors_in" or r._field == "errors_out" 
	                           or r._field == "interface_name" or r._field == "packets_recv" 
	                           or r._field == "packets_sent")
  	          |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`

	// Execute the query

	return queryInfluxDB(c, client, query, "Network")
}

func queryInfluxDB(c *fiber.Ctx, client influxdb2.Client, query string, metricstype string) error {
	queryAPI := client.QueryAPI(os.Getenv("INFLUXDB_ORG"))
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("Query error: %s", err))
	}
	defer result.Close()

	var data []map[string]interface{}

	switch metricstype {
	case "Disk":
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
			//fmt.Println("data", record)
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
	case "CpuLoad":
		for result.Next() {
			record := result.Record()
			//fmt.Println("data", record)
			// The fields are now columns, so we'll extract them individually
			row := map[string]interface{}{
				"time":    record.Time(),
				"load_1":  record.ValueByKey("load1"),
				"load_5":  record.ValueByKey("load5"),
				"load_15": record.ValueByKey("load15"),
			}
			data = append(data, row)
		}

	}

	if result.Err() != nil {
		return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("Query error: %s", result.Err()))
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": data, "msg": "data fetched succesfully"})
}
