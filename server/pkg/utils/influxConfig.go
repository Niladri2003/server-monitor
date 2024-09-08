package utils

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"os"
)

type InfluxConfig struct {
	InfluxClient influxdb2.Client
	// Add other dependencies if needed
}

func NewConfig() *InfluxConfig {
	client := influxdb2.NewClientWithOptions(os.Getenv("INFLUXDB_URL"), os.Getenv("INFLUXDB_TOKEN"), influxdb2.DefaultOptions().SetBatchSize(20))
	return &InfluxConfig{
		InfluxClient: client,
	}
}
