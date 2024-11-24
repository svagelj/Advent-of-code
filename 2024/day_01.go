package main

import (
	rw "aoc_2024/tools"
	"fmt"
)

func main() {

	fileName := "day01_data_2023.txt"
	fmt.Println("Reading file '" + fileName + "'")

	lines := rw.ReadFile(fileName)

	i := 0
	for i < len(lines) {
		line := lines[i]

		fmt.Printf("%2d: %s\n", i+1, line)

		j := 0
		for j < len(line) {
			fmt.Printf("\t %c\n", line[j])
			j++
		}

		i++
		if i >= 5 {
			break
		}
	}
}
