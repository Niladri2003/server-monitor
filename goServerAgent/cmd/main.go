//package main
//
//import (
//	"fmt"
//	"github.com/google/gopacket"
//	"github.com/google/gopacket/layers"
//	"github.com/google/gopacket/pcap"
//	"log"
//	"net/http"
//	"strings"
//	"sync/atomic"
//)

//	func main() {
//		// Periodically collect and display metrics
//		ticker := time.NewTicker(5 * time.Second) // Adjust the interval as needed
//		defer ticker.Stop()
//
//		for {
//			select {
//			case <-ticker.C:
//				goServerAgent.CollectMetrics()
//			}
//		}
//	}
package main

import (
	"fmt"
	"github.com/Niladri2003/server-monitor/goServerAgent/network"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	requestCount uint64
	ipAddresses  = make(map[string]uint64) // Track request counts by IP address
	methodCounts = make(map[string]uint64) // Track request methods (GET, POST, etc.)
	urlCounts    = make(map[string]uint64) // Track requested URLs
	statusCounts = make(map[int]uint64)    // Track HTTP status codes
	headerCounts = make(map[string]uint64) // Track specific headers
	payloadSizes = make(map[string]uint64) // Track payload sizes by URL
)

func main() {
	port := 5000 // Default port; this should be configurable

	// Monitor HTTP requests on the specified port
	go monitorHTTPRequests(port)

	// Start your main application logic here
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server is running. Requests count on port %d: %d\n", port, atomic.LoadUint64(&requestCount))
		fmt.Fprintf(w, "Requests by IP:\n")
		for ip, count := range ipAddresses {
			fmt.Fprintf(w, "IP %s: %d requests\n", ip, count)
		}
		fmt.Fprintf(w, "Request Methods:\n")
		for method, count := range methodCounts {
			fmt.Fprintf(w, "%s: %d requests\n", method, count)
		}
		fmt.Fprintf(w, "Requested URLs:\n")
		for url, count := range urlCounts {
			fmt.Fprintf(w, "%s: %d requests\n", url, count)
		}
		fmt.Fprintf(w, "HTTP Status Codes:\n")
		for status, count := range statusCounts {
			fmt.Fprintf(w, "%d: %d responses\n", status, count)
		}
		fmt.Fprintf(w, "Payload Sizes by URL:\n")
		for url, size := range payloadSizes {
			fmt.Fprintf(w, "%s: %d bytes\n", url, size)
		}
	})
	select {}
}

func monitorHTTPRequests(port int) {
	handle, err := pcap.OpenLive("ens5", 1600, true, pcap.BlockForever) // Replace "eth0" with your network interface
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	filter := fmt.Sprintf("tcp port %d", port)
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		ip, method, url, statusCode, payloadSize, isHTTP := processPacket(packet)
		if isHTTP {
			atomic.AddUint64(&requestCount, 1)
			ipAddresses[ip]++
			methodCounts[method]++
			urlCounts[url]++
			statusCounts[statusCode]++
			payloadSizes[url] += payloadSize
			fmt.Printf("HTTP Request from IP %s. Method: %s, URL: %s, Status: %d, Payload Size: %d bytes\n", ip, method, url, statusCode, payloadSize)
		}
	}
}

func processPacket(packet gopacket.Packet) (string, string, string, int, uint64, bool) {
	// Extract IP layer
	var ip gopacket.Layer
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		ip = ipLayer
	} else {
		ipLayer = packet.Layer(layers.LayerTypeIPv6)
		if ipLayer != nil {
			ip = ipLayer
		} else {
			return "", "", "", 0, 0, false
		}
	}

	// Extract TCP layer
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer == nil {
		return "", "", "", 0, 0, false
	}
	tcp, _ := tcpLayer.(*layers.TCP)

	// Check if the packet has a payload
	if len(tcp.Payload) == 0 {
		return "", "", "", 0, 0, false
	}

	// Extract IP address
	var sourceIP string
	if ipv4, ok := ip.(*layers.IPv4); ok {
		sourceIP = ipv4.NetworkFlow().Src().String()
	} else if ipv6, ok := ip.(*layers.IPv6); ok {
		sourceIP = ipv6.NetworkFlow().Src().String()
	}

	// Extract HTTP request details from payload
	payload := string(tcp.Payload)
	if len(payload) > 0 && (payload[:3] == "GET" || payload[:4] == "POST") {
		method, url, statusCode, payloadSize := network.ExtractHTTPDetails(payload)
		return sourceIP, method, url, statusCode, payloadSize, true
	}

	return "", "", "", 0, 0, false
}
