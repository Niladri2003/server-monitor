package InfluxSetup

import (
	"context"
	"fmt"
	"github.com/Niladri2003/server-monitor/server/metrics"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"log"
	"time"
)

func StoreMetricsFromKafka(client influxdb2.Client, msg metrics.MetricMessage) error {
	// Create a WriteAPIBlocking

	writeAPI := client.WriteAPIBlocking("sysmos", "test-bucket")

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

	cpu := make([]*write.Point, 0) // Create a slice to hold all the points

	for i, times := range msg.Metrics.CPU.Times {
		cpuMetricsPoint := influxdb2.NewPointWithMeasurement("cpu_metrics").
			AddTag("server_id", msg.ServerID).
			AddTag("api_key", msg.APIKey).
			AddTag("core", fmt.Sprintf("%d", i)).
			AddField("model", msg.Metrics.CPU.Model).                            // CPU model (same for all cores)
			AddField("cores", msg.Metrics.CPU.Cores).                            // Total number of cores (same for all cores)
			AddField("usage_per_core_percent", msg.Metrics.CPU.UsagePerCore[i]). // Usage for each specific core
			AddField("user_time_sec", times.User).                               // CPU time in user mode
			AddField("system_time_sec", times.System).                           // CPU time in system mode
			AddField("idle_time_sec", times.Idle).                               // CPU time idle
			AddField("iowait_time_sec", times.Iowait).                           // CPU time waiting for I/O
			SetTime(timestamp)

		err := writeAPI.WritePoint(context.Background(), cpuMetricsPoint)
		if err != nil {
			return err
		}
		// Add the point to the batch
		cpu = append(cpu, cpuMetricsPoint)
	}
	err = writeAPI.Flush(context.Background())
	if err != nil {
		return err
	}

	// Create the point for load averages
	loadAvgPoint := influxdb2.NewPointWithMeasurement("load_averages").
		AddTag("server_id", msg.ServerID).
		AddTag("api_key", msg.APIKey).
		AddField("load1", msg.Metrics.LoadAverages.Load1).
		AddField("load5", msg.Metrics.LoadAverages.Load5).
		AddField("load15", msg.Metrics.LoadAverages.Load15).
		SetTime(timestamp)

	// Create a slice to hold all disk points
	diskPoints := make([]*write.Point, 0)

	// Add a point for disk usage
	diskUsagePoint := influxdb2.NewPointWithMeasurement("disk_metrics").
		AddTag("server_id", msg.ServerID).
		AddTag("api_key", msg.APIKey).
		AddField("total_gb", msg.Metrics.Disk.Total).
		AddField("free_gb", msg.Metrics.Disk.Free).
		AddField("used_gb", msg.Metrics.Disk.Used).
		AddField("used_percent", msg.Metrics.Disk.UsedPercent).
		SetTime(timestamp)

	// Add the disk usage point to the slice
	diskPoints = append(diskPoints, diskUsagePoint)

	// Add points for disk I/O stats
	for _, diskStat := range msg.Metrics.Disk.IOStats {
		diskMetricsPoint := influxdb2.NewPointWithMeasurement("disk_metrics").
			AddTag("server_id", msg.ServerID).
			AddTag("api_key", msg.APIKey).
			AddTag("disk_name", diskStat.Name).
			AddField("read_bytes_mb", diskStat.ReadBytes).
			AddField("write_bytes_mb", diskStat.WriteBytes).
			SetTime(timestamp)
		err := writeAPI.WritePoint(context.Background(), diskMetricsPoint)
		if err != nil {
			return err
		}
		// Add the disk I/O point to the slice
		diskPoints = append(diskPoints, diskMetricsPoint)
	}
	err = writeAPI.Flush(context.Background())
	if err != nil {
		return err
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
		//if err := writeAPI.WritePoint(context.Background(), networkPoint); err != nil {
		//	return fmt.Errorf("failed to write network point: %w", err)
		//}
		err := writeAPI.WritePoint(context.Background(), networkPoint)
		if err != nil {
			return err
		}
	}
	err = writeAPI.Flush(context.Background())
	if err != nil {
		return err
	}

	// Add points for top processes by cpu
	top5cpu := make([]*write.Point, 0)
	for _, proc := range msg.Top5CPU {
		processPoint := influxdb2.NewPointWithMeasurement("top_processes_by_cpu").
			AddTag("server_id", msg.ServerID).
			AddTag("api_key", msg.APIKey).
			AddTag("pid", fmt.Sprintf("%d", proc.Pid)).
			AddField("name", proc.Name).
			AddField("cpu_percent", proc.CPU).
			AddField("memory_percent", proc.Memory).
			SetTime(timestamp)
		// Write each point individually
		top5cpu = append(top5cpu, processPoint)
		err := writeAPI.WritePoint(context.Background(), processPoint)
		if err != nil {
			return err
		}
	}
	err = writeAPI.Flush(context.Background())
	if err != nil {
		return err
	}
	// Add points for top processes by memory
	top5memory := make([]*write.Point, 0)
	for _, proc := range msg.Top5Memory {
		processPoint := influxdb2.NewPointWithMeasurement("top_processes_by_memory").
			AddTag("server_id", msg.ServerID).
			AddTag("api_key", msg.APIKey).
			AddTag("pid", fmt.Sprintf("%d", proc.Pid)).
			AddField("name", proc.Name).
			AddField("cpu_percent", proc.CPU).
			AddField("memory_percent", proc.Memory).
			SetTime(timestamp)
		err := writeAPI.WritePoint(context.Background(), processPoint)
		if err != nil {
			return err
		}
		// Write each point individually
		top5cpu = append(top5memory, processPoint)
	}
	err = writeAPI.Flush(context.Background())
	if err != nil {
		return err
	}

	// Create the point for Host_info
	systemHostInfo := influxdb2.NewPointWithMeasurement("host_info").
		AddTag("server_id", msg.ServerID).
		AddTag("api_key", msg.APIKey).
		AddField("hostname", msg.Metrics.SystemInfo.Hostname).
		AddField("os", msg.Metrics.SystemInfo.OS).
		AddField("platform_version", msg.Metrics.SystemInfo.PlatformVersion).
		AddField("kernel_version", msg.Metrics.SystemInfo.KernelVersion).
		AddField("uptime_hours", msg.Metrics.SystemInfo.UptimeHours).
		AddField("boot_time", msg.Metrics.SystemInfo.BootTime).
		SetTime(timestamp)
	// Create the point for swap_memory
	SwapMemoryInfo := influxdb2.NewPointWithMeasurement("swap_memory").
		AddTag("server_id", msg.ServerID).
		AddTag("api_key", msg.APIKey).
		AddField("total_gb", msg.Metrics.Swap.Total).
		AddField("used_gb", msg.Metrics.Swap.Used).
		AddField("free_gb", msg.Metrics.Swap.Free).
		AddField("used_percent", msg.Metrics.Swap.UsedPercent).
		SetTime(timestamp)

	//slice to hold all points
	allPoints := make([]*write.Point, 0)

	allPoints = append(allPoints, memoryPoint, loadAvgPoint, systemHostInfo, SwapMemoryInfo)
	//allPoints = append(allPoints, cpu...)
	//allPoints = append(allPoints, diskPoints...)
	//allPoints = append(allPoints, top5cpu...)
	//allPoints = append(allPoints, top5memory...)
	//fmt.Println(len(allPoints))
	// Write the points to InfluxSetup
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = writeAPI.WritePoint(ctx, allPoints...)
	if err != nil {
		return fmt.Errorf("failed to write points: %w", err)
	}

	log.Printf("Metrics stored for server: %s", msg.ServerID)
	return nil
}
