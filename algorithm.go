package main

import "fmt"

func algorithm(config Config, unMap unavailablesMap, servers []*Server) int {
	rows := make([]Row, config.rows)
	// Assign Server
	for sPos := 0; sPos < len(servers); sPos++ {
		server := servers[sPos]
		placeServer(config, unMap, server, sPos, rows)
	}

	// Assign pool
	for rPos := 0; rPos < len(rows); rPos++ {
		for sPos := 0; sPos < len(rows[rPos].servers); sPos++ {
			rows[rPos].servers[sPos].assignedPool = sPos % config.nPools
		}
	}

	return 42
}

type Row struct {
	servers []*Server
}

func placeServer(config Config, unMap unavailablesMap, server *Server, sPos int, rows []Row) {
	for i := 0; i < config.rows; i++ {
		for j := 0; j < config.slots; j++ {
			if j+server.size > config.slots {
				break
			}
			canFit := true
			for k := j; k < j+server.size && k < config.slots; k++ {
				if ok := unMap[fmt.Sprintf("%d %d", i, k)]; ok {
					canFit = false
					break
				}
			}
			if !canFit {
				continue
			}

			server.assignedRow = i
			server.assignedSlot = j
			server.assigned = true

			// Slots unavailable
			if rows[i].servers == nil {
				rows[i].servers = make([]*Server, 0)
			}
			rows[i].servers = append(rows[i].servers, server)
			for k := j; k < j+server.size && k < config.slots; k++ {
				unMap[fmt.Sprintf("%d %d", i, k)] = true
			}
			return
		}
	}
}
