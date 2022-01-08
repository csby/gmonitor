package gmonitor

import (
	"testing"
	"time"
)

func TestTotalCPUPercent(t *testing.T) {
	all1, busy1, err := AllCpuTotalBusyTime()
	if err != nil {
		t.Error(err)
		return
	}

	time.Sleep(time.Second)

	all2, busy2, err := AllCpuTotalBusyTime()
	if err != nil {
		t.Error(err)
		return
	}

	percent := ToCpuPercent(all1, busy1, all2, busy2)
	t.Log(percent)
}

func TestCpuName(t *testing.T) {
	name, err := CpuName()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(name)
}
