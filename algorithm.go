package main

import "fmt"

func algorithm(config Config, unMap unavailablesMap, servers []*Server) int {
	// occupiedMap := make(unavailablesMap, 0)
	rows := make([]*Row, config.rows)
	for sPos := 0; sPos < len(servers); sPos++ {
		server := servers[sPos]
		fmt.Printf("Working for server %d - %+v\n", sPos, server)
		fmt.Printf("UN MAP %+v\n", unMap)
		placeServer(config, unMap, server, sPos, rows)
	}

	fmt.Printf("%+v\n", rows)

	return 42
}

type Row struct {
	servers []*Server
}

func placeServer(config Config, unMap unavailablesMap, server *Server, sPos int, rows []*Row) {
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
			server.assignedPool = 1

			// Slots unavailable
			fmt.Printf("BEFORE ITERATING FOR SERVER %d - %+v\n", sPos, unMap)
			rows[i].servers = append(rows[i].servers, server)
			for k := j; k < j+server.size && k < config.slots; k++ {
				fmt.Printf("SETTING TRUE %d %d\n", i, k)
				unMap[fmt.Sprintf("%d %d", i, k)] = true
			}
			return
		}
	}
}
