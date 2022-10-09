package gmonitor

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"strings"
	"time"
)

type Host struct {
	ID              string   `json:"id" note:"主机标识"`
	Name            string   `json:"name" note:"主机名称"`
	BootTime        DateTime `json:"bootTime" note:"系统启动时间"`
	OS              string   `json:"os" note:"操作系统, 如windows, linux"`
	Platform        string   `json:"platform" note:"系统平台, 如ubuntu, Microsoft Windows 10 企业版"`
	PlatformVersion string   `json:"platformVersion" note:"平台版本, 如10.0.17134 Build 17134"`
	KernelVersion   string   `json:"kernelVersion" note:"内核版本, 如4.15.0-22-generic"`
	CPU             string   `json:"cpu" note:"处理器"`
	Memory          string   `json:"memory" note:"系统内存, 如16GB"`
	TimeZone        string   `json:"timeZone" note:"系统时区, 如CST+8"`
}

func (s *Host) Stat() error {
	v, err := host.Info()
	if err != nil {
		return err
	}

	s.ID = v.HostID
	s.Name = v.Hostname
	s.BootTime = DateTime(time.Unix(int64(v.BootTime), 0))
	s.OS = v.OS
	s.Platform = v.Platform
	s.PlatformVersion = v.PlatformVersion
	s.KernelVersion = v.KernelVersion

	c, err := cpu.Info()
	if err == nil {
		if len(c) > 0 {
			s.CPU = fmt.Sprintf("%s x%d", c[0].ModelName, len(c))
		}
	}
	m, err := mem.VirtualMemory()
	if err == nil {
		if m.Total > 0 {
			s.Memory = fmt.Sprintf("%s / %s (%d%%)", toText(float64(m.Used)), toText(float64(m.Total)), 100*m.Used/m.Total)
		} else {
			s.Memory = fmt.Sprintf("%s / %s", toText(float64(m.Used)), toText(float64(m.Total)))
		}
	}

	zoneName, zoneOffset := time.Now().Local().Zone()
	timeZone := strings.Builder{}
	timeZone.WriteString(zoneName)
	if zoneOffset >= 0 {
		timeZone.WriteString("+")
	}
	timeZone.WriteString(fmt.Sprint(zoneOffset / 60 / 60))
	s.TimeZone = timeZone.String()

	return nil
}
