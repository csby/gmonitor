package gmonitor

import "github.com/shirou/gopsutil/mem"

type Memory struct {
	Total         uint64  `json:"total"`
	TotalText     string  `json:"totalText"`
	Available     uint64  `json:"available"`
	AvailableText string  `json:"availableText"`
	Used          uint64  `json:"used"`
	UsedText      string  `json:"usedText"`
	UsedPercent   float64 `json:"usedPercent"`
}

func (s *Memory) Stat() error {
	v, err := mem.VirtualMemory()
	if err != nil {
		return err
	}

	s.Total = v.Total
	s.Available = v.Available
	s.Used = v.Used
	s.UsedPercent = v.UsedPercent
	s.TotalText = toText(float64(v.Total))
	s.AvailableText = toText(float64(v.Available))
	s.UsedText = toText(float64(v.Used))

	return nil
}
