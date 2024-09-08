package controllers

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"net/http"
	"os"
)

func GetMetrics(c *fiber.Ctx, client influxdb2.Client) error {
	query := `from(bucket: "` + os.Getenv("INFLUXDB_BUCKET") + `")
              |> range(start: -1h)
              |> filter(fn: (r) => r["_measurement"] == "system_metrics")`

	return queryInfluxDB(c, client, query)
}

func GetDiskUsage(c *fiber.Ctx, client influxdb2.Client) error {
	query := `from(bucket: "` + os.Getenv("INFLUXDB_BUCKET") + `")
              |> range(start: -6h)
              |> filter(fn: (r) => r._measurement == "disk")
              |> filter(fn: (r) => r._field == "free_gb" or r._field == "total_gb" or r._field == "used_percent" or r._field == "used_gb")
  	          |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`

	return queryInfluxDB(c, client, query)
}

func GetNetworkStats(c *fiber.Ctx, client influxdb2.Client) error {
	query := `from(bucket: "` + os.Getenv("INFLUXDB_BUCKET") + `")
              |> range(start: -1h)
              |> filter(fn: (r) => r["_measurement"] == "network_stats")`

	return queryInfluxDB(c, client, query)
}

func queryInfluxDB(c *fiber.Ctx, client influxdb2.Client, query string) error {
	queryAPI := client.QueryAPI(os.Getenv("INFLUXDB_ORG"))
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("Query error: %s", err))
	}
	defer result.Close()

	var data []map[string]interface{}
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
	if result.Err() != nil {
		return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("Query error: %s", result.Err()))
	}

	return c.JSON(data)
}
