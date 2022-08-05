package gmonitor

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"
)

func getTcpListenPorts(ports *ListenCollection) {
	var stdout bytes.Buffer
	cmd := exec.Command("netstat", "-ano", "-p", "tcp")
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 协议  本地地址          外部地址        状态           PID
	// TCP   0.0.0.0:135     0.0.0.0:0     LISTENING      1108
	for {
		line, err := stdout.ReadString('\n')
		if err == io.EOF {
			break
		}
		if len(line) < 6 {
			continue
		}
		fields := make([]string, 0)
		for _, field := range strings.Split(line, " ") {
			val := strings.TrimSpace(field)
			if len(val) < 1 {
				continue
			}
			fields = append(fields, val)
		}
		if len(fields) < 4 {
			continue
		}
		if fields[3] != "LISTENING" {
			continue
		}
		ipPort := fields[1]
		pos := strings.LastIndex(ipPort, ":")
		if pos < 1 {
			continue
		}
		ip := ipPort[0:pos]
		port := ipPort[pos+1:]
		portVal, err := strconv.Atoi(port)
		if err != nil {
			continue
		}
		pid := 0
		pidVal, err := strconv.Atoi(fields[4])
		if err == nil {
			pid = pidVal
		}

		listen := &Listen{
			Address:  ip,
			Port:     portVal,
			Protocol: "tcp",
			PId:      pid,
		}
		*ports = append(*ports, listen)
	}
}

func getUdpListenPorts(ports *ListenCollection) {
	var stdout bytes.Buffer
	cmd := exec.Command("netstat", "-ano", "-p", "udp")
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 协议  本地地址          外部地址        状态           PID
	// UDP   0.0.0.0:53      *:*                          3820
	addrPort := make(map[string]byte)
	for {
		line, err := stdout.ReadString('\n')
		if err == io.EOF {
			break
		}
		if len(line) < 6 {
			continue
		}
		fields := make([]string, 0)
		for _, field := range strings.Split(line, " ") {
			val := strings.TrimSpace(field)
			if len(val) < 1 {
				continue
			}
			fields = append(fields, val)
		}
		if len(fields) < 4 {
			continue
		}
		ipPort := fields[1]
		_, ok := addrPort[ipPort]
		if ok {
			continue
		}
		addrPort[ipPort] = 0
		pos := strings.LastIndex(ipPort, ":")
		if pos < 1 {
			continue
		}
		ip := ipPort[0:pos]
		port := ipPort[pos+1:]
		portVal, err := strconv.Atoi(port)
		if err != nil {
			continue
		}
		pid := 0
		pidVal, err := strconv.Atoi(fields[3])
		if err == nil {
			pid = pidVal
		}

		listen := &Listen{
			Address:  ip,
			Port:     portVal,
			Protocol: "udp",
			PId:      pid,
		}

		*ports = append(*ports, listen)
	}
}
