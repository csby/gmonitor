package gmonitor

import "testing"

func TestInterfaces(t *testing.T) {
	nics, err := Interfaces()
	if err != nil {
		t.Error(err)
	}

	count := len(nics)
	t.Log("network interface count:", count)
	for i := 0; i < count; i++ {
		nic := nics[i]
		t.Logf("nic-%d: %#v", i, nic)
	}
}

func TestTcpListenPorts(t *testing.T) {
	ports := TcpListenPorts()
	count := len(ports)
	t.Log("count = ", count)

	for i := 0; i < count; i++ {
		item := ports[i]
		t.Logf("%3d %18s:%-6d %s", i+1, item.Address, item.Port, item.Protocol)
	}
}

func TestUdpListenPorts(t *testing.T) {
	ports := UdpListenPorts()
	count := len(ports)
	t.Log("count = ", count)

	for i := 0; i < count; i++ {
		item := ports[i]
		t.Logf("%3d %18s:%-6d %s", i+1, item.Address, item.Port, item.Protocol)
	}
}

func TestListenPorts(t *testing.T) {
	ports := ListenPorts()
	count := len(ports)
	t.Log("count = ", count)

	for i := 0; i < count; i++ {
		item := ports[i]
		t.Logf("%3d %18s:%-6d %s", i+1, item.Address, item.Port, item.Protocol)
	}
}
