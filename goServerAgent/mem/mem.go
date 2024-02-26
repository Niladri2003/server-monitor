package memory

import "encoding/json"

type Memorymet struct {
	//Total amount of RAM on this system
	Total uint64 `json:"total"`
	//Available Ram for programs
	Available uint64 `json:"available"`
	//Total Ram used by the Programs
	Used uint64 `json:"used"`
	//Total Free Ram avilable
	Free        uint64  `json:"free"`
	Cache       uint64  `json:"cache"`
	Buffers     uint64  `json:"buffers"`
	UsedPercent float64 `json:"usedPercent"`
}

func (m Memorymet) string() string {
	s, _ := json.Marshal(m)
	return string(s)

}
