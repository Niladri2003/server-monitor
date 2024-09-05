package port

import (
	"fmt"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

// PortScanner holds the IP and a semaphore for limiting concurrent scans.
type PortScanner struct {
	ip   string
	lock *sync.WaitGroup
}

// Ulimit returns the current ulimit for open files.
func Ulimit() int64 {
	out, err := exec.Command("bash", "-c", "ulimit -n").Output()
	if err != nil {
		panic(err)
	}

	s := strings.TrimSpace(string(out))

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return i
}

// ScanPort checks if a port is open on the given IP address.
func ScanPort(ip string, port int, timeout time.Duration) bool {
	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", target, timeout)

	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(timeout)
			return ScanPort(ip, port, timeout)
		}
		return false
	}
	conn.Close()
	return true
}

// Start scans ports from f to l with a specified timeout.
func (ps *PortScanner) Start(f, l int, timeout time.Duration) {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	maxGoroutines := 100 // Limit to 100 concurrent scans
	sema := make(chan struct{}, maxGoroutines)

	for port := f; port <= l; port++ {
		wg.Add(1)
		sema <- struct{}{} // Block if there are already maxGoroutines running
		go func(port int) {
			defer wg.Done()
			defer func() { <-sema }() // Release one slot

			if ScanPort(ps.ip, port, timeout) {
				fmt.Printf("Port %d is open\n", port)
			}
		}(port)
	}
}
