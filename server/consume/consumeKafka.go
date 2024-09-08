package consume

import (
	"context"
	"encoding/json"
	"github.com/Niladri2003/server-monitor/server/InfluxSetup"
	"github.com/Niladri2003/server-monitor/server/metrics"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/segmentio/kafka-go"
	"log"
)

//type MetricMessage struct {
//	APIKey     string                `json:"api_key"`
//	ServerID   string                `json:"server_id"`
//	Timestamp  string                `json:"timestamp"`
//	Metrics    metrics.SystemMetrics `json:"metrics"`
//	Top5CPU    []metrics.ProcessInfo `json:"top5_cpu_processes"`
//	Top5Memory []metrics.ProcessInfo `json:"top5_memory_processes"`
//}

func ConsumeKafka(influxClient influxdb2.Client) error {

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "agent-data-topic",
		GroupID: "central-server-group",
	})

	for {
		// Read message from Kafka
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("failed to read message:", err)
		}
		// Parse the message into MetricMessage struct
		var msg metrics.MetricMessage
		err = json.Unmarshal(m.Value, &msg)
		if err != nil {
			log.Printf("failed to unmarshal message: %v", err)
			continue
		}

		//// Now you can access the metrics and metadata
		//log.Printf("Received data from server: %s", msg.ServerID)
		//log.Printf("Metrics: %v", msg.Metrics.Disk.Free)
		//log.Printf("Top 5 CPU processes: %v", msg.Top5CPU)
		//log.Printf("Top 5 Memory processes: %v", msg.Top5Memory)
		//log.Printf("Api Key: %v", msg.APIKey)
		//Store data in InfluxSetup
		// Store metrics in InfluxSetup
		if err := InfluxSetup.StoreMetricsFromKafka(influxClient, msg); err != nil {
			log.Printf("Failed to store metrics: %v", err)
		}
	}
}
