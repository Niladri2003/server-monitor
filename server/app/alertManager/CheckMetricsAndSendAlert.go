package alertManager

//func checkMetricsAndTriggerAlerts(serverID string) {
//	// Query the latest metrics from InfluxDB
//	metrics := queryLatestMetricsFromInfluxDB(serverID)
//
//	// Compare metrics with thresholds
//	if metrics.CPUUsage > thresholds["cpu_usage"] {
//		triggerAlert("CPU usage exceeded 80%", serverID)
//	}
//	if metrics.MemoryUsage > thresholds["memory_usage"] {
//		triggerAlert("Memory usage exceeded 90%", serverID)
//	}
//	if metrics.DiskUsage > thresholds["disk_usage"] {
//		triggerAlert("Disk usage exceeded 85%", serverID)
//	}
//}
