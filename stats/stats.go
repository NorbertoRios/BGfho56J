package stats

import "github.com/shirou/gopsutil/process"

//NewProcessStat ...
func NewProcessStat() *Process {
	return &Process{}
}

//Process ...
type Process struct {
	process.Process
	CPUPercent float64
}
