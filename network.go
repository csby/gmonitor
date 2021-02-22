package gmonitor

import (
	"github.com/shirou/gopsutil/v3/net"
	"sort"
	"strings"
)

type Interface struct {
	Name       string   `json:"name" note:"网卡名称"`
	MTU        int      `json:"mtu" note:"最大传输单元"`
	MacAddress string   `json:"macAddress" note:"MAC地址"`
	IPAddress  []string `json:"ipAddress" note:"IP地址"`
	Flags      []string `json:"flags" note:"标志, 如up, loopback, multicast"`
}

func Interfaces() ([]Interface, error) {
	vs, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	infos := make([]Interface, 0)
	for _, v := range vs {
		info := Interface{
			Name:       v.Name,
			MTU:        v.MTU,
			MacAddress: v.HardwareAddr,
			IPAddress:  make([]string, 0),
			Flags:      make([]string, 0),
		}
		ipCount := len(v.Addrs)
		for i := 0; i < ipCount; i++ {
			info.IPAddress = append(info.IPAddress, v.Addrs[i].Addr)
		}
		flagCount := len(v.Flags)
		for i := 0; i < flagCount; i++ {
			info.Flags = append(info.Flags, v.Flags[i])
		}

		infos = append(infos, info)
	}

	return infos, nil
}

type Listen struct {
	Address  string `json:"address" note:"地址"`
	Port     int    `json:"port" note:"端口"`
	Protocol string `json:"protocol" note:"协议"`
}
type ListenCollection []*Listen

func (s ListenCollection) Len() int {
	return len(s)
}
func (s ListenCollection) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ListenCollection) Less(i, j int) bool {
	if s[i].Port == s[j].Port {
		return strings.Compare(s[i].Address, s[j].Address) < 0
	}
	return s[i].Port < s[j].Port
}

func TcpListenPorts() []*Listen {
	listens := make(ListenCollection, 0)
	getTcpListenPorts(&listens)
	sort.Stable(listens)

	return listens
}

func UdpListenPorts() []*Listen {
	listens := make(ListenCollection, 0)
	getUdpListenPorts(&listens)
	sort.Stable(listens)

	return listens
}

func ListenPorts() []*Listen {
	listens := make(ListenCollection, 0)
	getTcpListenPorts(&listens)
	getUdpListenPorts(&listens)
	sort.Stable(listens)

	return listens
}
