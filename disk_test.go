package gmonitor

import (
	"testing"
)

func TestDiskPartitions(t *testing.T) {
	ps, err := DiskPartitions()
	if err != nil {
		t.Error(err)
		return
	}
	c := len(ps)
	t.Log("count: ", c)
	for i := 0; i < c; i++ {
		p := ps[i]
		if p == nil {
			continue
		}

		t.Logf("%d: %#v", i+1, p)
	}

	for i := 0; i < c; i++ {
		p := ps[i]
		if p == nil {
			continue
		}

		u, e := StatDiskUsage(p.MountPoint)
		if e != nil {
			t.Errorf("%d: %v", i+1, e)
		} else {
			t.Logf("%d: %#v", i+1, u)
		}
	}
}
