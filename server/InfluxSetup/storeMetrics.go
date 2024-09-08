package InfluxSetup

import (
	"context"
	"fmt"
	"github.com/Niladri2003/server-monitor/server/metrics"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"log"
	"time"
)

func StoreMetricsFromKafka(client influxdb2.Client, msg metrics.MetricMessage) error {
	// Create a WriteAPIBlocking

	writeAPI := client.WriteAPIBlocking("sysmos", "server-metrics")

	// Convert timestamp from string to time.Time
	timestamp, err := time.Parse(time.RFC3339, msg.Timestamp)
	if err != nil {
		return fmt.Errorf("failed to parse timestamp: %w", err)
	}

	// Create the point for memory metrics
	memoryPoint := influxdb2.NewPointWithMeasurement("memory").
		AddTag("server_id", msg.ServerID).
		AddTag("api_key", msg.APIKey).
		AddField("total_gb", msg.Metrics.Memory.Total).
		AddField("used_gb", msg.Metrics.Memory.Used).
		AddField("free_gb", msg.Metrics.Memory.Free).
		AddField("used_percent", msg.Metrics.Memory.UsedPercent).
		SetTime(timestamp)

	// Create the point for CPU metrics
	cpuPoint := influxdb2.NewPointWithMeasurement("cpu").
		AddTag("server_id", msg.ServerID).
		AddTag("api_key", msg.APIKey).
		AddField("model", msg.Metrics.CPU.Model).
		AddField("cores", msg.Metrics.CPU.Cores).
		SetTime(timestamp)

	// Add points for CPU times
	for i, times := range msg.Metrics.CPU.Times {
		cpuTimesPoint := influxdb2.NewPointWithMeasurement("cpu_times").
			AddTag("server_id", msg.ServerID).
			AddTag("api_key", msg.APIKey).
			AddTag("core", fmt.Sprintf("%d", i)).
			AddField("user_time_sec", times.User).
			AddField("system_time_sec", times.System).
			AddField("idle_time_sec", times.Idle).
			AddField("iowait_time_sec", times.Iowait).
			SetTime(timestamp)
		// Write each point individually
		if err := writeAPI.WritePoint(context.Background(), cpuTimesPoint); err != nil {
			return fmt.Errorf("failed to write CPU times point: %w", err)
		}
	}

	// Create the point for load averages
	loadAvgPoint := influxdb2.NewPointWithMeasurement("load_averages").
		AddTag("server_id", msg.ServerID).
		AddTag("api_key", msg.APIKey).
		AddField("load1", msg.Metrics.LoadAverages.Load1).
		AddField("load5", msg.Metrics.LoadAverages.Load5).
		AddField("load15", msg.Metrics.LoadAverages.Load15).
		SetTime(timestamp)

	// Add points for disk metrics
	for _, diskStat := range msg.Metrics.Disk.IOStats {
		diskPoint := influxdb2.NewPointWithMeasurement("disk_io").
			AddTag("server_id", msg.ServerID).
			AddTag("api_key", msg.APIKey).
			AddTag("disk_name", diskStat.Name).
			AddField("read_bytes_mb", diskStat.ReadBytes).
			AddField("write_bytes_mb", diskStat.WriteBytes).
			SetTime(timestamp)
		// Write each point individually
		if err := writeAPI.WritePoint(context.Background(), diskPoint); err != nil {
			return fmt.Errorf("failed to write disk I/O point: %w", err)
		}
	}
	// Add points for network metrics
	for _, netStat := range msg.Metrics.Network {
		networkPoint := influxdb2.NewPointWithMeasurement("network").
			AddTag("server_id", msg.ServerID).
			AddTag("api_key", msg.APIKey).
			AddTag("interface_name", netStat.Name).
			AddField("bytes_sent_mb", netStat.BytesSent).
			AddField("bytes_recv_mb", netStat.BytesRecv).
			AddField("packets_sent", netStat.PacketsSent).
			AddField("packets_recv", netStat.PacketsRecv).
			AddField("errors_in", netStat.ErrorsIn).
			AddField("errors_out", netStat.ErrorsOut).
			AddField("drops_in", netStat.DropsIn).
			AddField("drops_out", netStat.DropsOut).
			SetTime(timestamp)
		// Write each point individually
		if err := writeAPI.WritePoint(context.Background(), networkPoint); err != nil {
			return fmt.Errorf("failed to write network point: %w", err)
		}
	}
	// Add points for top processes
	for _, proc := range msg.Metrics.TopProcesses {
		processPoint := influxdb2.NewPointWithMeasurement("top_processes").
			AddTag("server_id", msg.ServerID).
			AddTag("api_key", msg.APIKey).
			AddTag("pid", fmt.Sprintf("%d", proc.Pid)).
			AddField("name", proc.Name).
			AddField("cpu_percent", proc.CPU).
			AddField("memory_percent", proc.Memory).
			SetTime(timestamp)
		// Write each point individually
		if err := writeAPI.WritePoint(context.Background(), processPoint); err != nil {
			return fmt.Errorf("failed to write process point: %w", err)
		}
	}
	// Create the point for load averages
	systemHostInfo := influxdb2.NewPointWithMeasurement("load_averages").
		AddTag("server_id", msg.ServerID).
		AddTag("api_key", msg.APIKey).
		AddField("hostname", msg.Metrics.SystemInfo.Hostname).
		AddField("os", msg.Metrics.SystemInfo.OS).
		AddField("platform_version", msg.Metrics.SystemInfo.PlatformVersion).
		AddField("kernel_version", msg.Metrics.SystemInfo.KernelVersion).
		AddField("uptime_hours", msg.Metrics.SystemInfo.UptimeHours).
		AddField("boot_time", msg.Metrics.SystemInfo.BootTime).
		SetTime(timestamp)

	// Write the points to InfluxSetup
	err = writeAPI.WritePoint(context.Background(), memoryPoint, cpuPoint, loadAvgPoint, systemHostInfo)
	if err != nil {
		return fmt.Errorf("failed to write points: %w", err)
	}

	log.Printf("Metrics stored for server: %s", msg.ServerID)
	return nil
}
