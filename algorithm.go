package main

import (
	"fmt"
	"sort"
)

type Row struct {
	servers []*Server
}

type Pools []struct {
	rows   map[int]int
	weight int
}

func (p Pools) GetPool(lastRow int) int {
	min := p[0].rows[lastRow]
	idx := 0
	for k, v := range p {
		lastRowWeight := v.rows[lastRow]
		if lastRowWeight < min || (lastRowWeight == min && v.weight < p[idx].weight) {
			min = lastRowWeight
			idx = k
		}
	}

	return idx
}

func algorithm(config Config, unMap unavailablesMap, initialServers []*Server) {

	servers := sortServers(initialServers)

	rows := make([]Row, config.rows)

	// Assign Server
	startRow := 0
	pools := make(Pools, config.nPools)
	for sPos := 0; sPos < len(servers); sPos++ {
		server := servers[sPos]
		lastRow := placeServer(startRow, config, unMap, server, sPos, rows)
		startRow = (lastRow + 1) % config.rows

		poolIdx := pools.GetPool(lastRow)
		servers[sPos].assignedPool = poolIdx
		if pools[poolIdx].rows == nil {
			pools[poolIdx].rows = make(map[int]int)
		}
		pools[poolIdx].rows[lastRow] += server.capacity
		pools[poolIdx].weight += server.capacity
	}

	// // Assign pool
	// currentPool := 0
	// for rPos := 0; rPos < len(rows); rPos++ {
	// 	for sPos := 0; sPos < len(rows[rPos].servers); sPos++ {
	// 		rows[rPos].servers[sPos].assignedPool = currentPool % config.nPools
	// 		currentPool++
	// 	}
	// }
}

func placeServer(startRow int, config Config, unMap unavailablesMap, server *Server, sPos int, rows []Row) int {
	for i := startRow; i < config.rows; i++ {
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
			return i
		}
	}
	for i := 0; i <= startRow; i++ {
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
			return i
		}
	}
	return -1
}

func sortServers(initialServers []*Server) []*Server {
	sort.Slice(initialServers, func(i, j int) bool {
		a := initialServers[i]
		b := initialServers[j]

		if a.capacity/a.size > b.capacity/b.size {
			// if a.capacity > b.capacity {
			return true
		}
		return false
	})
	return initialServers
	// initialServers = initialServers[:500]

	// part1 := make([]*Server, len(initialServers)/2) //+1)
	// part2 := make([]*Server, len(initialServers)/2)
	// for i := 0; i < len(initialServers); i += 2 {
	// 	part1[i/2] = initialServers[i]
	// 	if (i / 2) < len(part2) {
	// 		part2[(i / 2)] = initialServers[i+1]
	// 	}
	// }

	// sort.Slice(part1, func(i, j int) bool {
	// 	a := part1[i]
	// 	b := part1[j]

	// 	// if a.capacity/a.size > b.capacity/b.size {
	// 	if a.capacity < b.capacity {

	// 		return true
	// 	}
	// 	return false
	// })
	// sort.Slice(part2, func(i, j int) bool {
	// 	a := part2[i]
	// 	b := part2[j]

	// 	// if a.capacity/a.size < b.capacity/b.size {
	// 	if a.capacity > b.capacity {

	// 		return true
	// 	}
	// 	return false
	// })
	// servers := append(part1, part2...)
	// return servers
}
