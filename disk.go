package gmonitor

import (
	"github.com/shirou/gopsutil/disk"
)

type DiskPartition struct {
	Device     string `json:"device" note:"设备"`
	MountPoint string `json:"mountPoint" note:"挂载点"`
	FsType     string `json:"fsType" note:"文件系统"`
}

func DiskPartitions() ([]*DiskPartition, error) {
	ps, err := disk.Partitions(true)
	if err != nil {
		return nil, err
	}

	results := make([]*DiskPartition, 0)
	c := len(ps)
	for i := 0; i < c; i++ {
		p := ps[i]
		results = append(results, &DiskPartition{
			Device:     p.Device,
			MountPoint: p.Mountpoint,
			FsType:     p.Fstype,
		})
	}

	return results, nil
}

type DiskUsage struct {
	Path        string  `json:"path" note:"路径"`
	FsType      string  `json:"fsType" note:"文件系统"`
	Total       uint64  `json:"total" note:"总容量, 单位字节"`
	Free        uint64  `json:"free" note:"已用容量, 单位字节"`
	Used        uint64  `json:"used" note:"剩余容量, 单位字节"`
	UsedPercent float64 `json:"usedPercent" note:"使用率"`

	TotalText string `json:"totalText"`
	UsedText  string `json:"usedText"`
	FreeText  string `json:"freeText"`
}

func StatDiskUsage(path string) (*DiskUsage, error) {
	stat, err := disk.Usage(path)
	if err != nil {
		return nil, err
	}

	result := &DiskUsage{
		Path:        stat.Path,
		FsType:      stat.Fstype,
		Total:       stat.Total,
		Free:        stat.Free,
		Used:        stat.Used,
		UsedPercent: stat.UsedPercent,
	}
	result.TotalText = toText(float64(stat.Total))
	result.UsedText = toText(float64(stat.Used))
	result.FreeText = toText(float64(stat.Free))

	return result, nil
}
