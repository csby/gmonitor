package gmonitor

import (
	"bytes"
	"io"
	"os/exec"
	"strconv"
	"strings"
)

func getTcpListenPorts(ports *ListenCollection) {
	var stdout bytes.Buffer
	cmd := exec.Command("ss", "-ltn")
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		return
	}

	for {
		line, err := stdout.ReadString('\n')
		if err == io.EOF {
			break
		}
		if len(line) < 6 {
			continue
		}
		if line[0] != 'L' || line[5] != 'N' {
			continue
		}
		listen := parseListenPort(line)
		if listen == nil {
			continue
		}
		listen.Protocol = "tcp"
		*ports = append(*ports, listen)
	}
}

func getUdpListenPorts(ports *ListenCollection) {
	var stdout bytes.Buffer
	cmd := exec.Command("ss", "-lun")
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		return
	}

	for {
		line, err := stdout.ReadString('\n')
		if err == io.EOF {
			break
		}
		if len(line) < 6 {
			continue
		}
		if line[0] != 'U' || line[5] != 'N' {
			continue
		}
		listen := parseListenPort(line)
		if listen == nil {
			continue
		}
		listen.Protocol = "udp"
		*ports = append(*ports, listen)
	}
}

func parseListenPort(line string) *Listen {
	fields := make([]string, 0)
	for _, field := range strings.Split(line, " ") {
		val := strings.TrimSpace(field)
		if len(val) < 1 {
			continue
		}
		fields = append(fields, val)
	}
	if len(fields) < 4 {
		return nil
	}
	ipPort := fields[3]
	pos := strings.LastIndex(ipPort, ":")
	if pos < 1 {
		return nil
	}
	ip := ipPort[0:pos]
	port := ipPort[pos+1:]
	portVal, err := strconv.Atoi(port)
	if err != nil {
		return nil
	}

	listen := &Listen{
		Address: ip,
		Port:    portVal,
	}

	return listen
}
