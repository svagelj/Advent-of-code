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
	"..@@.@@@@.",
	"@@@.@.@.@@",
	"@@@@@.@.@@",
	"@.@@@@..@.",
	"@@.@@@@.@@",
	".@@@@@@@.@",
	".@.@.@.@@@",
	"@.@@@.@@@@",
	".@@@@@@@@.",
	"@.@.@@@.@.",
}

var testSolution1, testSolution2 = 13, -1

//------------------------------------------------------

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed   :)"
	} else {
		return "failed   :(   (should be '"+strconv.Itoa(solValue)+"')"
	}
}

//------------------------------------------------------

func initData(fileLines []string, rollChar rune) ([][]rune, [][]int) {

	rollsMap := [][]rune {}
	rollsPositions := [][]int {}

	for i := range fileLines {
		line := fileLines[i]

		_line := []rune {}
		for j := range line {
			_line = append(_line, rune(line[j]))
			if rune(line[j]) == rollChar {
				rollsPositions = append(rollsPositions, []int {i,j})
			}
		}

		rollsMap = append(rollsMap, _line)
	}

	return rollsMap, rollsPositions
}

func solve1(rollsMap [][]rune, rollsPositions [][]int, rollChar rune, printout bool) int {

	if printout {
		Printer.PrintGridRune(rollsMap, 1)
		fmt.Println(rollsPositions)
	}

	yMax := len(rollsMap)
	xMax := len(rollsMap[0])
	found := Array.CopyRune2D(rollsMap)

	sum := 0
	for k := range rollsPositions {

		pos := rollsPositions[k]
		y := pos[0]
		x := pos[1]
		nFound := 0

		// loops to get all combinations of x=+-1, y=+-1
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				}
				if y + i < 0 || y + i >= yMax {
					continue
				}
				if x + j < 0 || x + j >= xMax {
					continue
				}
				
				if rollsMap[y+i][x+j] == rollChar {
					nFound = nFound + 1
				}
			}
		}

		if nFound <= 3{
			found[y][x] = rune('x')
			sum = sum + 1
		}
	}

	if printout {
		Printer.PrintGridRune(found, 1)
	}

	return sum
}

//----------------------------------------

func solve2(batteryBanks []string, printout bool) int {

	if printout {
		fmt.Println("battery bank:", batteryBanks)
	}

	sum := 0
	return sum
}

func main() {

	// data gathering and parsing
	rollsChar := '@'
	rollsMapTest, rollsPosTest := initData(testData, rollsChar)

	fileName := "day_04_data.txt"
	fileData := rw.ReadFile(fileName)
	rollsMap1, rollsPos1 := initData(fileData, rollsChar)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_1_test := solve1(rollsMapTest, rollsPosTest, rollsChar, true)
	fmt.Println("Test solution 1 =", sol1_1_test, "->", checkSolution(sol1_1_test, testSolution1))

	t1 := time.Now()
	sol1 := solve1(rollsMap1, rollsPos1, rollsChar, false)
	dur := time.Since(t1)
	fmt.Println("Solution part 1 =", sol1, "(ET =", dur, ")")

	// ---------------------------------------------
	// fmt.Println()
	// fmt.Println("=== Part 2 ===")
	// sol2_1_test := solve2(batteryBanksTest, true)
	// fmt.Println("Test solution 2 =", sol2_1_test, "->", checkSolution(sol2_1_test, testSolution2))

	// t1 = time.Now()
	// sol2 := solve2(batteryBanks1, false)
	// dur = time.Since(t1)
	// fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
