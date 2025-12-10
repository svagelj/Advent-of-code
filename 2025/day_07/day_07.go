package main

import (
	// Math "aoc_2025/tools/Math"
	Array "aoc_2025/tools/Array"
	Printer "aoc_2025/tools/Printer"
	rw "aoc_2025/tools/rw"
	"time"

	"fmt"
	"strconv"

	// "slices"
	// "sort"
	// "strings"
	// "bufio"
	// "os"
)

// var must be used for global variables
var testData = []string {
	".......S.......",
	"...............",
	".......^.......",
	"...............",
	"......^.^......",
	"...............",
	".....^.^.^.....",
	"...............",
	"....^.^...^....",
	"...............",
	"...^.^...^.^...",
	"...............",
	"..^...^.....^..",
	"...............",
	".^.^.^.^.^...^.",
	"...............",
}

var testSolution1, testSolution2 = 21, 40

//------------------------------------------------------

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed   :)"
	} else {
		return "failed   :(   (should be '"+strconv.Itoa(solValue)+"')"
	}
}

//------------------------------------------------------

func initData(fileLines []string) ([][]rune, []int) {

	diagram := [][]rune {}
	start := []int {}

	startChar := 'S'

	for i := range fileLines {
		line := fileLines[i]

		_line := []rune {}
		for j := range line {
			_line = append(_line, rune(line[j]))
			if rune(line[j]) == startChar {
				start = []int {i,j}
			}
		}
		diagram = append(diagram, _line)
	}

	return diagram, start
}

func solve1(diagram [][]rune, start []int, printout bool) int {

	if printout {
		Printer.PrintGridRune(diagram, 1)
		fmt.Println("start", start)
	}

	queue := [][]int {start}
	splitterChar := '^'
	laserPath := Array.CopyRune2D(diagram)
	visitedSplitters := Array.InitArrayValuesInt(len(diagram), len(diagram[0]), 0)
	sizeY := len(diagram)

	for k := 0; k < len(queue); k++ {
		// the loop must include len(queue) because it is expanding

		newQueue := [][]int {}
		y := queue[k][0]
		x := queue[k][1]

		// current place was already visited
		if laserPath[y][x] == '|' {
			continue
		}

		laserPath[y][x] = '|'

		for j := y+1; j < sizeY; j++ {
			if diagram[j][x] == splitterChar {
				newQueue = append(newQueue, []int{j, x-1})
				newQueue = append(newQueue, []int{j, x+1})
				
				visitedSplitters[j][x] = visitedSplitters[j][x] + 1
				break
			}

			laserPath[j][x] = '|'
		}

		queue = append(queue, newQueue...)
	}

	if printout {
		fmt.Println("Finished")
		Printer.PrintGridRune(laserPath, 1)
		fmt.Println()
		Printer.PrintGridInt(visitedSplitters, 2)
	}
	
	// Count visited splitters
	sum := 0
	for i := range visitedSplitters {
		for j := range visitedSplitters[i] {
			if visitedSplitters[i][j] != 0 {
				sum = sum + 1
			}
		}
	}

	return sum
}

//----------------------------------------

func solve2(diagram [][]rune, start []int, printout bool) int {

	if printout {
		Printer.PrintGridRune(diagram, 1)
		fmt.Println("start", start)
	}

	splitterChar := '^'
	laserChar := '|'
	sizeY := len(diagram)
	laserPaths := Array.InitArrayValuesInt(len(diagram), len(diagram[0]), 0)

	visitedSplitters := Array.InitArrayValuesInt(len(diagram), len(diagram[0]), 0)
	laserPath := Array.CopyRune2D(diagram)

	// put in start
	laserPaths[start[0]][start[1]] = 1
	laserPath[start[0]][start[1]] = laserChar

	// Change the approach - do it line by line for all lasers
	for i := range (sizeY-1) {

		for j := range diagram[i] {
			if laserPaths[i][j] == 0 {
				continue
			}

			if diagram[i+1][j] == splitterChar {
				// split the beam
				visitedSplitters[i+1][j] = visitedSplitters[i+1][j] + laserPaths[i][j]
				laserPaths[i+1][j+1] = laserPaths[i+1][j+1] + laserPaths[i][j]
				laserPaths[i+1][j-1] = laserPaths[i+1][j-1] + laserPaths[i][j]
			} else {
				laserPaths[i+1][j] = laserPaths[i+1][j] + laserPaths[i][j]
			}
		}
	}

	if printout {
		fmt.Println("Finished")
		Printer.PrintGridRune(laserPath, 1)
		fmt.Println()
		Printer.PrintGridInt(visitedSplitters, 2)
		fmt.Println()
		Printer.PrintGridInt(laserPaths, 3)
	}

	// Sum last row
	sum := 0
	for i := range len(laserPaths[sizeY-1]) {
		sum = sum + laserPaths[sizeY-1][i]
	}

	return sum
}

func main() {

	// data gathering and parsing
	diagramTest1, startTest1 := initData(testData)

	fileName := "day_07_data.txt"
	fileData := rw.ReadFile(fileName)
	diagram1, start1 := initData(fileData)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_1_test := solve1(diagramTest1, startTest1, true)
	fmt.Println("Test solution 1 =", sol1_1_test, "->", checkSolution(sol1_1_test, testSolution1))

	t1 := time.Now()
	sol1 := solve1(diagram1, start1, false)
	dur := time.Since(t1)
	fmt.Println("Solution part 1 =", sol1, "(ET =", dur, ")")

	// ---------------------------------------------

	fmt.Println()
	fmt.Println("=== Part 2 ===")
	sol2_1_test := solve2(diagramTest1, startTest1, true)
	fmt.Println("Test solution 2 =", sol2_1_test, "->", checkSolution(sol2_1_test, testSolution2))

	t1 = time.Now()
	sol2 := solve2(diagram1, start1, false)
	dur = time.Since(t1)
	fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
