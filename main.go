package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	files := []string{
		// Uncomment the line with the desired files (add other lines if needed)
		// "a",
		"b",
		// "a", "b", "c", "d", "e", "f",
		// "a", "b",
		// "a", "b", "e", "f",
		// "c",
		// "d",
	}

	for _, fileName := range files {
		fmt.Printf("****************** INPUT: %s\n", fileName)
		inputSet := readFile(fmt.Sprintf("./inputFiles/%s.in", fileName))

		config, unMap, servers := buildInput(inputSet)
		// fmt.Printf("CONFIG %+v\nunMap: %+v\nservers %+v\n", config, unMap, servers)
		// for _, s := range servers {
		// 	fmt.Printf("Server: %+v\n", s)
		// }
		// printInputMetrics(input)

		algorithm(config, unMap, servers)

		// for _, s := range servers {
		// 	fmt.Printf("Server End: %+v\n", s)
		// }
		output := buildOutput(servers)
		// printResultMetrics(servers)

		ioutil.WriteFile(fmt.Sprintf("./result/%s.out", fileName), []byte(output), 0644)
	}
}
