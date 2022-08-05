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
	cmd := exec.Command("ss", "-ltn", "-p")
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
	cmd := exec.Command("ss", "-lun", "-p")
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

	pid := 0
	if len(fields) > 5 {
		//  users:(("sshd",pid=4128,fd=3))
		pidInfo := fields[5]
		pos = strings.LastIndex(pidInfo, "pid=")
		if pos >= 0 {

			for _, pidField := range strings.Split(pidInfo, ",") {
				pidFieldVal := strings.TrimSpace(pidField)
				if len(pidFieldVal) < 5 {
					continue
				}
				if strings.LastIndex(pidFieldVal, "pid=") == 0 {
					pidNumber := strings.ReplaceAll(pidFieldVal, "pid=", "")
					pidVal, pe := strconv.Atoi(pidNumber)
					if pe == nil {
						pid = pidVal
					}
				}
			}
		}
	}

	listen := &Listen{
		Address: ip,
		Port:    portVal,
		PId:     pid,
	}

	return listen
}
