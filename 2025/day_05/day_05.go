package main

import (
	// Math "aoc_2025/tools/Math"
	// Array "aoc_2025/tools/Array"
	// Printer "aoc_2025/tools/Printer"
	rw "aoc_2025/tools/rw"
	"time"

	"fmt"
	"strconv"

	"slices"
	// "sort"
	"strings"
	// "bufio"
	// "os"
)

// var must be used for global variables
var testData = []string {
	"3-5",
	"10-14",
	"16-20",
	"12-18",
	"",
	"1",
	"5",
	"8",
	"11",
	"17",
	"32",
}

var testSolution1, testSolution2 = 3, 14

//------------------------------------------------------

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed   :)"
	} else {
		return "failed   :(   (should be '"+strconv.Itoa(solValue)+"')"
	}
}

//------------------------------------------------------

func initData(fileLines []string) ([][]int, []int) {

	ranges := [][]int {}
	values := []int {}

	processingRanges := true
	for i := range fileLines {
		line := fileLines[i]

		if line == "" {
			processingRanges = false
			continue
		}

		if processingRanges {
			rangeSplit := strings.Split(line, "-")

			rangeMin, err := strconv.Atoi(rangeSplit[0])
			if err != nil {
				panic(err)
			}
			rangeMax, err := strconv.Atoi(rangeSplit[1])
			if err != nil {
				panic(err)
			}

			ranges = append(ranges, []int {rangeMin, rangeMax})
		} else {
			value, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			values = append(values, value)
		}
	}

	return ranges, values
}

func solve1(ranges [][]int, values []int, printout bool) int {

	if printout {
		fmt.Println("ranges", ranges)
		fmt.Println("values", values)
	}

	sum := 0
	for i := range values {
		if printout {
			fmt.Println(values[i])
		}

		for j:= range ranges {
			if values[i] >= ranges[j][0] && values[i] <= ranges[j][1] {

				if printout {
					fmt.Println("   ->", true)
				}

				sum = sum + 1
				break
			}
		}

	}

	return sum
}

//----------------------------------------

func sortRanges(ranges [][]int) [][]int {

	sorted := [][]int {}

	for i := range ranges {

		// get index to insert current range
		j := 0
		for j < len(sorted) {

			if sorted[j][0] > ranges[i][0] {
				break
			}

			j = j + 1
		}

		// insert new range in its place
		_sorted := slices.Clone(sorted[:j])
		_sorted = append(_sorted, ranges[i])

		if j < len(sorted) {
			_sorted = append(_sorted, sorted[j:]...)
		}
		
		sorted = _sorted
	}

	return sorted
}

func solve2(ranges [][]int, printout bool) int {

	if printout {
		fmt.Println(ranges)
	}

	// First sort ranges by lower boundary
	sortedRanges := sortRanges(ranges)

	simplifiedRanges := [][]int {}
	for i := range sortedRanges {

		if i == 0 {
			simplifiedRanges = append(simplifiedRanges, sortedRanges[i])
		} else {
			_range := sortedRanges[i]
			
			foundOverlap := false
			for j := range simplifiedRanges {
				rangeSimp := simplifiedRanges[j]

				if _range[1] > rangeSimp[1] && _range[0] <= rangeSimp[1]  {
					// check _range overlaps upper bounds
					simplifiedRanges[j][1] = _range[1]
					foundOverlap = true
				} else if _range[1] >= rangeSimp[0] && _range[0] < rangeSimp[0]  {
					// check _range overlaps lower bounds
					simplifiedRanges[j][0] = _range[0]
					foundOverlap = true
				} else if _range[0] == rangeSimp[0] && _range[1] > rangeSimp[1] {
					// expanding upper limit (lower is the same point)
					simplifiedRanges[j][1] = _range[1]
					foundOverlap = true
				} else if _range[1] == rangeSimp[1] && _range[0] < rangeSimp[0] {
					// expanding lower limit (upper is the same point)
					simplifiedRanges[j][0] = _range[0]
					foundOverlap = true
				}  else if _range[1] >= rangeSimp[1] && _range[0] <= rangeSimp[0] {
					// _range is bigger than simplified range
					simplifiedRanges[j] = _range
					foundOverlap = true
				} else if _range[1] <= rangeSimp[1] && _range[0] >= rangeSimp[0] {
					// _range is inside of simplified range
					foundOverlap = true
				} 
			}

			if !foundOverlap {
				simplifiedRanges = append(simplifiedRanges, _range)
			}
		}
	}

	sum := 0
	for i := range simplifiedRanges {
		_range := simplifiedRanges[i]
		sum = sum + (_range[1]-_range[0] + 1)
	}
	return sum
}

func main() {

	// data gathering and parsing
	rangesTest, valuesPosTest := initData(testData)

	fileName := "day_05_data.txt"
	fileData := rw.ReadFile(fileName)
	ranges1, values1 := initData(fileData)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_1_test := solve1(rangesTest, valuesPosTest, true)
	fmt.Println("Test solution 1 =", sol1_1_test, "->", checkSolution(sol1_1_test, testSolution1))

	t1 := time.Now()
	sol1 := solve1(ranges1, values1, false)
	dur := time.Since(t1)
	fmt.Println("Solution part 1 =", sol1, "(ET =", dur, ")")

	// ---------------------------------------------
	fmt.Println()
	fmt.Println("=== Part 2 ===")
	sol2_1_test := solve2(rangesTest, true)
	fmt.Println("Test solution 2 =", sol2_1_test, "->", checkSolution(sol2_1_test, testSolution2))

	t1 = time.Now()
	sol2 := solve2(ranges1, false)
	dur = time.Since(t1)
	fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
