package memory

import "encoding/json"

type Memorymet struct {
	// Total amount of RAM on this system
	Total uint64 `json:"total"`

	// RAM available for programs to allocate
	Available uint64 `json:"available"`

	// RAM used by programs
	Used uint64 `json:"used"`

	// Percentage of RAM used by programs
	UsedPercent float64 `json:"usedPercent"`

	// This is the kernel's notion of free memory
	Free uint64 `json:"free"`

	// Memory used by the page cache and slabs
	Cached uint64 `json:"cached"`

	// Memory used by kernel buffers
	Buffers uint64 `json:"buffers"`

	// Total swap memory
	SwapTotal uint64 `json:"swapTotal"`

	// Free swap memory
	SwapFree uint64 `json:"swapFree"`

	// Timestamp to capture the time of the metrics
	Timestamp int64 `json:"timestamp"`
}

func (m Memorymet) string() string {
	s, _ := json.Marshal(m)
	return string(s)

}
