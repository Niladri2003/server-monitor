package DataModel

type CPUStats struct {
	User            uint64              `json:"user"`            // Time spent in user mode
	Nice            uint64              `json:"nice"`            // Time spent in user mode with low priority (nice)
	System          uint64              `json:"system"`          // Time spent in system mode
	Idle            uint64              `json:"idle"`            // Time spent idle
	Iowait          uint64              `json:"iowait"`          // Time spent waiting for I/O
	Irq             uint64              `json:"irq"`             // Time spent servicing interrupts
	Softirq         uint64              `json:"softirq"`         // Time spent servicing softirqs
	Steal           uint64              `json:"steal"`           // Time spent in other operating systems (for virtual CPUs)
	Guest           uint64              `json:"guest"`           // Time spent running a guest operating system
	GuestNice       uint64              `json:"guestNice"`       // Time spent running a niced guest
	ContextSwitches uint64              `json:"contextSwitches"` // Number of context switches
	Processes       uint64              `json:"processes"`       // Number of processes created
	SoftIrqTotal    uint64              `json:"softIrqTotal"`    // Total softirq count
	SoftIrqDetails  map[string]uint64   `json:"softIrqDetails"`  // Details for each type of softirq
	LogicalCPUs     map[string]CPUStats `json:"logicalCPUs"`     // Map to store CPU stats for each logical CPU (e.g., cpu0, cpu1, etc.)
}
