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
	client := influxdb2.NewClient(os.Getenv("INFLUXDB_URL"), os.Getenv("INFLUXDB_TOKEN"))
	return &InfluxConfig{
		InfluxClient: client,
	}
}
