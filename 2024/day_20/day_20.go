package main

import (
	// Math "aoc_2024/tools/Math"
	Array "aoc_2024/tools/Array"
	Printer "aoc_2024/tools/Printer"
	rw "aoc_2024/tools/rw"
	"time"

	"fmt"
	// "sort"
	"strconv"
	"strings"

	// "bufio"
	// "os"
)

// var must be used for global variables
var testData = []string {
	"###############",
	"#...#...#.....#",
	"#.#.#.#.#.###.#",
	"#S#...#.#.#...#",
	"#######.#.#.###",
	"#######.#.#...#",
	"#######.#.###.#",
	"###..E#...#...#",
	"###.#######.###",
	"#...###...#...#",
	"#.#####.#.###.#",
	"#.#...#.#.#...#",
	"#.#.#.#.#.#.###",
	"#...#...#...###",
	"###############",
}

var testSolution1, testSolution2 = 44, -1

//------------------------------------------------------

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed   :)"
	} else {
		return "failed   :(   (should be '"+strconv.Itoa(solValue)+"')"
	}
}

func checkSolutionStr(testValue string, solValue string) string {
	if testValue == solValue {
		return "passed   :)"
	} else {
		return "failed   :(   (should be '"+solValue+"')"
	}
}

//------------------------------------------------------

func initData(fileLines []string) ([][]rune, [2]int, [2]int) {

	data := [][]rune {}
	startInd := [2]int {-1,-1}
	endInd := [2]int {-1,-1}

	startChar := "S"
	endChar := "E"

	for i := range fileLines {
		line := fileLines[i]

		data = append(data, []rune(line))

		// get starting position
		if strings.Contains(line, startChar) {
			startInd[0] = i
			startInd[1] = Array.GetIndexString(line, startChar)
		}

		// get end position
		if strings.Contains(line, endChar) {
			endInd[0] = i
			endInd[1] = Array.GetIndexString(line, endChar)
		}
	}

	return data, startInd, endInd
}

func getOriginalPath(data [][]rune, startInd [2]int, endInd [2]int) [][]int {

	path := [][]int {{startInd[0], startInd[1]}}
	Ny := len(data)
	Nx := len(data[0])

	for k := range Nx*Ny {
		y,x := path[k][0], path[k][1]
		if y == endInd[0] && x == endInd[1] {
			break
		}

		if data[y][x+1] != '#' && Array.GetIndexInt2D(path, []int{y,x+1}) == -1 {
			path = append(path, []int {y,x+1})
		} else if data[y][x-1] != '#' && Array.GetIndexInt2D(path, []int{y,x-1}) == -1 {
			path = append(path, []int {y,x-1})
		} else if data[y+1][x] != '#' && Array.GetIndexInt2D(path, []int{y+1,x}) == -1 {
			path = append(path, []int {y+1,x})
		} else if data[y-1][x] != '#'  && Array.GetIndexInt2D(path, []int{y-1,x}) == -1 {
			path = append(path, []int {y-1,x})
		}
	}

	return path
}

func tryShortcut(data [][]rune, path [][]int, pointInd int, Ny int, Nx int) []int {

	queue := [][]int {}
	y, x := path[pointInd][0], path[pointInd][1]

	// look all four direction if there is path on the other side of the path
	if x+2 < Nx && data[y][x+1] == '#' && data[y][x+2] != '#' {
		queue = append(queue, []int {y,x+2})
	}
	if x-2 >= 0 && data[y][x-1] == '#' && data[y][x-2] != '#' {
		queue = append(queue, []int {y,x-2})
	}
	if y+2 < Ny && data[y+1][x] == '#' && data[y+2][x] != '#' {
		queue = append(queue, []int {y+2,x})
	}
	if y-2 >= 0 && data[y-1][x] == '#' && data[y-2][x] != '#' {
		queue = append(queue, []int {y-2,x})
	}

	// if there is at least one direction with shortcut
	// find where each point is on path - this is landing point from jumping point (pointInd)
	// subtract landing_point index from jumping_point to get shortcut length (or time cut)
	shortcutLengths := []int {}
	for k :=  range queue {
		landingInd := Array.GetIndexInt2D(path, queue[k])
		if landingInd != -1 {
			shortcutLengths = append(shortcutLengths, (landingInd - pointInd)-2)
		}
	}

	return shortcutLengths
}

func solve1(data [][]rune, startInd [2]int, endInd [2]int, minShortcutLength int, printout bool) int {

	if printout {
		Printer.PrintGridRune(data, 2)
		fmt.Println("start:", startInd, "end:", endInd)
	}

	dataCopy := Array.CopyRune2D(data)
	path := getOriginalPath(dataCopy, startInd, endInd)

	// Try shortcut at every point of the path
	Ny, Nx := len(data), len(data[0])
	sum := 0
	for k := range path {
		pathCopy := append([][]int {}, path...)
		shortcutLengths := tryShortcut(dataCopy, pathCopy, k, Ny, Nx)

		for i := range shortcutLengths {
			if shortcutLengths[i] >= minShortcutLength {
				sum = sum + 1
			}
		}
	}

	return sum
}

//----------------------------------------

func solve2(data [][]rune, startInd []int, endInd []int, printout bool) int {

	if printout {
		fmt.Println("start:", startInd, "end:", endInd)
	}

	sum := 0
	return sum
}

func main() {

	// data gathering and parsing
	dataTest, startTets, endTest := initData(testData)

	fileName := "day_20_data.txt"
	fileData := rw.ReadFile(fileName)
	data, start, end := initData(fileData)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_1_test := solve1(dataTest, startTets, endTest, 0, true)
	fmt.Println("Test solution 1 =", sol1_1_test, "->", checkSolution(sol1_1_test, testSolution1))

	t1 := time.Now()
	sol1 := solve1(data, start, end, 100, false)
	dur := time.Since(t1)
	fmt.Println("Solution part 1 =", sol1, "(ET =", dur, ")")

	// ---------------------------------------------
	// fmt.Println()
	// fmt.Println("=== Part 2 ===")
	// sol2_1_test := solve2(towelsTest, patternsTest, true)		// 2, 1, 4, 6, 0, 1, 0
	// fmt.Println("Test solution 2 =", sol2_1_test, "->", checkSolution(sol2_1_test, testSolution2))

	// t1 = time.Now()
	// sol2 := solve2(towels, patterns, false)
	// dur = time.Since(t1)
	// fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
