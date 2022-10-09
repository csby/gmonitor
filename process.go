package gmonitor

import "github.com/shirou/gopsutil/process"

type Process struct {
	Pid     int    `json:"pid" note:"进程ID"`
	Name    string `json:"name" note:"进程名称"`
	Exe     string `json:"exe" note:"程序路径"`
	Cmdline string `json:"cmdline" note:"命令行参数"`
}

func GetProcessInfo(pid int) (*Process, error) {
	p, e := process.NewProcess(int32(pid))
	if e != nil {
		return nil, e
	}

	info := &Process{Pid: pid}
	info.Name, e = p.Name()
	info.Exe, e = p.Exe()
	info.Cmdline, e = p.Cmdline()

	return info, nil
}
