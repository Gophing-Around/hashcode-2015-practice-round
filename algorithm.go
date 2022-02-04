package main

import "fmt"

func algorithm(config Config, unMap unavailablesMap, servers []*Server) int {
	// occupiedMap := make(unavailablesMap, 0)
	for sPos := 0; sPos < len(servers); sPos++ {
		server := servers[sPos]
		fmt.Printf("Working for server %d - %+v\n", sPos, server)
		fmt.Printf("UN MAP %+v\n", unMap)
		placeServer(config, unMap, server, sPos)
	}

	return 42
}

func placeServer(config Config, unMap unavailablesMap, server *Server, sPos int) {

	for i := 0; i < config.rows; i++ {
		for j := 0; j < config.slots; j++ {
			// funziona?

			// è occupato

			// ho spazio negli slot successivi rispetto a size
			canFit := true
			for k := j; k < j+server.size && k < config.slots; k++ {
				if ok := unMap[fmt.Sprintf("%d %d", i, k)]; ok {
					canFit = false
					break
				}
				// if ok := occupiedMap[fmt.Sprintf("%d %d", i, k)]; ok {
				// 	break
				// }
			}
			if !canFit {
				continue
			}

			server.assignedRow = i
			server.assignedSlot = j
			server.assignedPool = 0

			// Slots unavailable
			fmt.Printf("BEFORE ITERATING FOR SERVER %d - %+v\n", sPos, unMap)
			for k := j; k < j+server.size && k < config.slots; k++ {

				fmt.Printf("SETTING TRUE %d %d\n", i, k)
				unMap[fmt.Sprintf("%d %d", i, k)] = true
			}
			return
			// j+=server.size
		}
	}
}
