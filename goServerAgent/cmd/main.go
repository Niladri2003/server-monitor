package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Niladri2003/server-monitor/goServerAgent"
	"github.com/fatih/color"
	"github.com/segmentio/kafka-go"
	"github.com/shirou/gopsutil/v4/process"
	"github.com/spf13/viper"
	"log"
	"time"
)

func printBanner() {
	banner := `
      ______   ______  __  __  ___  ____  
     / ___\ \ / / ___||  \/  |/ _ \/ ___| 
    \___ \\ V /\___ \| |\/| | | | \___ \ 
     ___) || |  ___) | |  | | |_| |___) |
    |____/ |_| |____/|_|  |_|\___/|____/ 
		SysMos Server Monitoring System made with ♡ by Niladri
`
	// Set the color to cyan
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Println(cyan(banner))
}

type Config struct {
	Interval    int    `mapstructure:"interval"`
	APIKey      string `mapstructure:"api_key"`
	KafkaBroker string `mapstructure:"kafka_broker"`
	ServerId    string `mapstructure:"server_id"`
	Topic       string `mapstructure:"topic"`
}

func loadConfig(configPath string) (Config, error) {
	printBanner()
	var config Config

	// Set up Viper to read the config file
	viper.SetConfigFile(configPath)

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		return config, fmt.Errorf("error reading config file: %w", err)
	}

	// Unmarshal the config into the Config struct
	if err := viper.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("error unmarshalling config: %w", err)
	}

	return config, nil
}

type MetricMessage struct {
	APIKey     string                      `json:"api_key"`
	ServerID   string                      `json:"server_id"`
	Timestamp  string                      `json:"timestamp"`
	Metrics    goServerAgent.SystemMetrics `json:"metrics"`
	Top5CPU    []goServerAgent.ProcessInfo `json:"top5_cpu_processes"`
	Top5Memory []goServerAgent.ProcessInfo `json:"top5_memory_processes"`
}

func main() {
	// Load configuration
	// Accept a command-line flag for the config file path
	configPath := flag.String("config", "config.yaml", "Path to the config file")
	flag.Parse()

	// Load configuration
	config, err := loadConfig(*configPath)
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}
	//Print Config file
	fmt.Println("---------config----------\n")
	fmt.Println("Server Id => %s", config.ServerId)
	// Kafka writer configuration
	writer := kafka.Writer{
		Addr:     kafka.TCP(config.KafkaBroker),
		Topic:    config.Topic,
		Balancer: &kafka.LeastBytes{},
	}

	defer writer.Close()

	// Periodically collect and display metrics
	// Check if interval is less than 10, and set it to 10 if necessary
	interval := config.Interval
	if interval < 10 {
		interval = 10 // Minimum interval is 10 seconds
	}

	// Periodically collect and display metrics
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()
	// Adjust the interval as needed

	for {
		select {
		case <-ticker.C:

			metrics := goServerAgent.CollectMetrics()
			processes, err := process.Processes()
			if err != nil {
				fmt.Println(err)
			}
			top5byCpu := goServerAgent.TopProcessesByCPU(processes, 5)
			top5byMemory := goServerAgent.TopProcessesByMemory(processes, 5)
			message := MetricMessage{
				APIKey:     config.APIKey,
				ServerID:   config.ServerId, // Unique server identifier
				Timestamp:  time.Now().Format(time.RFC3339),
				Metrics:    metrics,
				Top5CPU:    top5byCpu,
				Top5Memory: top5byMemory,
			}
			messageBytes, err := json.Marshal(message)
			if err != nil {
				log.Fatal("Failed to marshal message", err)
			}
			log.Println("Sending metrics to Kafka...")

			// Send collected data to Kafka
			err = writer.WriteMessages(context.Background(),
				kafka.Message{
					Key:   []byte(config.APIKey),
					Value: []byte(messageBytes),
				},
			)
			if err != nil {
				log.Fatal("failed to write message:", err)
			}
		}
	}
}

//package main
//
//import (
//	"bufio"
//	"bytes"
//	"fmt"
//	"log"
//	"net/http"
//	"strings"
//	"sync/atomic"
//	"time"
//
//	"github.com/google/gopacket"
//	"github.com/google/gopacket/layers"
//	"github.com/google/gopacket/pcap"
//)
//
//var (
//	requestCount uint64
//	ipAddresses  = make(map[string]uint64) // Track request counts by IP address
//	methodCounts = make(map[string]uint64) // Track request methods (GET, POST, etc.)
//	urlCounts    = make(map[string]uint64) // Track requested URLs
//	statusCounts = make(map[int]uint64)    // Track HTTP status codes
//)
//
//func main() {
//	port := 5000 // Default port; this should be configurable
//
//	// Monitor HTTP requests on the specified port
//	go monitorHTTPRequests(port)
//
//	// Start your main application logic here
//	select {}
//}
//
//func monitorHTTPRequests(port int) {
//	// Get the default network interface
//	iface, err := pcap.FindAllDevs()
//	if err != nil {
//		log.Fatal(err)
//	}
//	if len(iface) == 0 {
//		log.Fatal("No network interfaces found")
//	}
//	fmt.Println(iface[0].Name)
//	// Open a live capture on the first available network interface
//	handle, err := pcap.OpenLive("lo", 1600, true, pcap.BlockForever)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer handle.Close()
//
//	// Set a BPF filter to capture only TCP traffic on the specified port
//	filter := fmt.Sprintf("tcp port %d", port)
//	err = handle.SetBPFFilter(filter)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
//	var tcpBuffer bytes.Buffer
//
//	for packet := range packetSource.Packets() {
//		// Process each packet
//
//		processPacket(packet, &tcpBuffer)
//	}
//}
//
//func processPacket(packet gopacket.Packet, buffer *bytes.Buffer) {
//	// Extract IP layer
//	fmt.Println("Step")
//	ipLayer := packet.Layer(layers.LayerTypeIPv4)
//	if ipLayer == nil {
//		return // Not an IP packet
//	}
//
//	// Extract TCP layer
//	tcpLayer := packet.Layer(layers.LayerTypeTCP)
//	if tcpLayer == nil {
//		return // Not a TCP packet
//	}
//	tcp, _ := tcpLayer.(*layers.TCP)
//
//	// Reassemble TCP payload (HTTP request data)
//	payload := tcp.Payload
//	if len(payload) == 0 {
//		return // No payload, skip
//	}
//	// Append payload to the buffer
//	buffer.Write(payload)
//	fmt.Println("data", buffer.String())
//	if strings.Contains(buffer.String(), "\r\n\r\n") {
//		// Try parsing the payload as an HTTP request once we have full headers
//		req, err := http.ReadRequest(bufio.NewReader(buffer))
//		if err == nil {
//			// Successfully parsed HTTP request
//			atomic.AddUint64(&requestCount, 1)
//			ipSrc := packet.NetworkLayer().NetworkFlow().Src().String()
//			methodCounts[req.Method]++
//			urlCounts[req.URL.String()]++
//			ipAddresses[ipSrc]++
//
//			// Log the request time
//			timestamp := time.Now().Format(time.RFC3339)
//			fmt.Printf("[%s] HTTP Request from IP: %s, Method: %s, URL: %s\n", timestamp, ipSrc, req.Method, req.URL.String())
//
//			// Clear the buffer after successfully reading the request
//			buffer.Reset()
//		} else {
//			// If it's an incomplete HTTP request or parsing error, just continue
//			fmt.Println("Malformed HTTP request")
//			fmt.Println(err)
//			buffer.Reset() // Clear buffer to avoid duplicate errors on partial data
//		}
//	}
//}
//
//// mockReader simulates an io.Reader for parsing raw TCP payloads as HTTP requests
//type mockReader struct {
//	data []byte
//}
//
//func (r *mockReader) Read(p []byte) (n int, err error) {
//	if len(r.data) == 0 {
//		return 0, fmt.Errorf("EOF")
//	}
//	n = copy(p, r.data)
//	r.data = r.data[n:]
//	return n, nil
//}

//package main

//import (
//	"log"
//	"sync"
//	"time"
//)
//
//type PortScanner struct {
//	ip   string
//	lock *sync.WaitGroup
//}
//
//func main() {
//	ps := &PortScanner{
//		ip:   "13.60.54.61",
//		lock: &sync.WaitGroup{},
//	}
//	ps.Start(1, 65535, 500*time.Millisecond)
//	log.Println("Port Scanning completed on specific %d", ps.ip)
//}
