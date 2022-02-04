package main

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

func buildOutput(result int) string {
	return "42"
}
