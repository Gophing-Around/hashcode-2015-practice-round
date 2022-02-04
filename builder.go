package main

import (
	"fmt"
	"strings"
)

type Config struct {
	rows        int
	slots       int
	unavailable int
	nPools      int
	nServers    int
}

type Server struct {
	capacity int
	size     int

	assignedRow  int
	assignedSlot int
	assignedPool int
	assigned     bool
}

type unavailablesMap map[string]bool

func buildInput(inputSet string) (Config, unavailablesMap, []*Server) {
	lines := splitNewLines(inputSet)
	configItem := splitSpaces(lines[0])
	config := Config{
		rows:        toint(configItem[0]),
		slots:       toint(configItem[1]),
		unavailable: toint(configItem[2]),
		nPools:      toint(configItem[3]),
		nServers:    toint(configItem[4]),
	}

	unavailbles := make(unavailablesMap, 0)
	for i := 1; i <= config.unavailable; i++ {
		unavailbles[lines[i]] = true
	}

	servers := make([]*Server, 0)
	for i := 1 + config.unavailable; i <= config.unavailable+config.nServers; i++ {
		line := splitSpaces(lines[i])
		servers = append(servers, &Server{
			size:     toint(line[0]),
			capacity: toint(line[1]),
		})
	}

	return config, unavailbles, servers
}

func buildOutput(servers []*Server) string {
	result := ""
	for _, server := range servers {
		if server.assigned {
			result += fmt.Sprintf("%d %d %d\n", server.assignedRow, server.assignedSlot, server.assignedPool)
		} else {
			result += fmt.Sprintf("x\n")
		}
	}

	return strings.Trim(result, "\n")
}
