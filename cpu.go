package gmonitor

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
)

func CpuName() (string, error) {
	infos, err := cpu.Info()
	if err != nil {
		return "", err
	}

	c := len(infos)
	if c < 1 {
		return "", fmt.Errorf("not found")
	}

	info := infos[0]
	return fmt.Sprintf("%s x%d", info.ModelName, c), nil
}

func AllCpuTotalBusyTime() (float64, float64, error) {
	ts, err := cpu.Times(false)
	if err != nil {
		return 0, 0, err
	}
	if len(ts) < 1 {
		return 0, 0, fmt.Errorf("not found")
	}
	t := ts[0]
	busy := t.User +
		t.System +
		t.Nice +
		t.Iowait +
		t.Irq +
		t.Softirq +
		t.Steal

	return busy + t.Idle, busy, nil
}

func ToCpuPercent(all1, busy1, all2, busy2 float64) float64 {
	if busy2 <= busy1 {
		return 0
	}

	if all2 <= all1 {
		return 100
	}

	percent := 100 * (busy2 - busy1) / (all2 - all1)
	if percent < 0 {
		percent = 0
	} else if percent > 100 {
		percent = 100
	}

	return percent
}
