package tools

import (
	// "bufio"
	"fmt"
	// "os"
)

// Functions with names starting with upper case are exported

func PrintGridRune(grid [][]rune){

	for i := range grid {
		for j:= range grid[i] {
			fmt.Printf("%c ", grid[i][j])
		}
		fmt.Printf("\n")
	}
}

func PrintGridInt(grid [][]int){

	for i := range grid {
		for j:= range grid[i] {
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Printf("\n")
	}
}