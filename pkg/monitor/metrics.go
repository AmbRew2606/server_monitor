package monitor

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

type Metrics struct {
	CPUUsage float64 `json:"cpu_usage"`
	RAMUsage float64 `json:"ram_usage"`
}

func GetMetrics() (*Metrics, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	c, err := cpu.Percent(0, false)
	if err != nil {
		return nil, err
	}

	return &Metrics{
		CPUUsage: c[0],
		RAMUsage: v.UsedPercent,
	}, nil
}
