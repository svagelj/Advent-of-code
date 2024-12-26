package tools

import (
	"fmt"
	Array "aoc_2024/tools/Array"
)

// Functions with names starting with upper case are exported

func PrintGridRune(grid [][]rune, pad int){

	for i := range grid {
		for j:= range grid[i] {
			fmt.Printf("%*c", pad, grid[i][j])
		}
		fmt.Printf("\n")
	}
}

func PrintGridInt(grid [][]int, pad int){

	for i := range grid {
		for j:= range grid[i] {
			fmt.Printf("%*d", pad, grid[i][j])
		}
		fmt.Printf("\n")
	}
}

func PrintRegion(nodes [][2]int, M int, N int, pad int) {

	area := Array.InitArrayValuesRune(M,N, '.')
	for k := range nodes {
		i,j := nodes[k][0], nodes[k][1]
		area[i][j] = '#'
	}
	PrintGridRune(area, pad)
}