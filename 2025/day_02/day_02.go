package main

import (
	// Math "aoc_2025/tools/Math"
	// Array "aoc_2025/tools/Array"
	// Printer "aoc_2025/tools/Printer"
	rw "aoc_2025/tools/rw"
	"time"

	"fmt"
	"strconv"
	// "slices"
	// "sort"
	"strings"
	// "bufio"
	// "os"
)

// var must be used for global variables
var testData = []string {
	"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124",
}

var testSolution1, testSolution2 = 1227775554, 4174379265

//------------------------------------------------------

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed   :)"
	} else {
		return "failed   :(   (should be '"+strconv.Itoa(solValue)+"')"
	}
}

//------------------------------------------------------

func initData(fileLines []string) ([][]int) {

	intervalsInt := [][]int {}

	for i := range fileLines {
		line := fileLines[i]

		intervals := strings.Split(line, ",")

		for j := range intervals {
			intervalStr := strings.Split(intervals[j], "-")

			intervalMin, err := strconv.Atoi(intervalStr[0])
			if err != nil {
				panic(err)
			}
			intervalMax, err := strconv.Atoi(intervalStr[1])
			if err != nil {
				panic(err)
			}

			intervalsInt = append(intervalsInt, []int {intervalMin, intervalMax})
		}
	}

	return intervalsInt
}

func solve1(intervals [][]int, printout bool) int {

	if printout {
		fmt.Println("intervals:", intervals)
	}

	sum := 0
	maxIntervalLength := 0
	invalidIds := []int {}

	for i := range intervals {

		minValue := intervals[i][0]
		maxValue := intervals[i][1]
		if maxValue - minValue > maxIntervalLength {
			maxIntervalLength = maxValue - minValue
		}
		
		// loop over each value in current interval
		value := minValue
		for value <= maxValue {

			valueStr := strconv.Itoa(value)

			n := len(valueStr)
			nHalf := n/2

			if n%2 == 0 && valueStr[:nHalf] == valueStr[nHalf:] {
				sum = sum + value
				invalidIds = append(invalidIds, value)
			}

			value = value + 1
		}

	}

	if printout {
		fmt.Println("max interval length:", maxIntervalLength)
		fmt.Println("Invalid ids:", invalidIds)
	}

	return sum
}

//----------------------------------------

func isBodyMadeOfKeys(key string, body string) bool {

	if len(key) > len(body) {
		return false
	}
	
	if len(body) % len(key) != 0 {
		return false
	}

	n := len(key)
	i := 0
	for i + n <= len(body) {
		if key != body[i:i+n] {
			return false
		}
		i = i + n
	}

	return true
}

func solve2(intervals [][]int, printout bool) int {

	if printout {
		fmt.Println("intervals:", intervals)
	}

	sum := 0

	maxIntervalLength := 0
	invalidIds := []int {}

	for i := range intervals {

		minValue := intervals[i][0]
		maxValue := intervals[i][1]
		if maxValue - minValue > maxIntervalLength {
			maxIntervalLength = maxValue - minValue
		}
		
		// loop over each value in current interval
		value := minValue
		for value <= maxValue {

			valueStr := strconv.Itoa(value)
			n := len(valueStr)
				
			// loop over each key lengths
			keyLen := 1
			for keyLen <= n/2 {
				key := valueStr[:keyLen]
				body := valueStr[keyLen:]
				
				if isBodyMadeOfKeys(key, body) {
					sum = sum + value
					invalidIds = append(invalidIds, value)
					break
				}

				keyLen = keyLen + 1
			}
			
			value = value + 1
		}
	}

	if printout {
		fmt.Println("Invalid ids:", invalidIds)
	}

	return sum
}

func main() {

	// data gathering and parsing
	intervalsTest := initData(testData)

	fileName := "day_02_data.txt"
	fileData := rw.ReadFile(fileName)
	intervals1 := initData(fileData)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_1_test := solve1(intervalsTest, true)
	fmt.Println("Test solution 1 =", sol1_1_test, "->", checkSolution(sol1_1_test, testSolution1))

	t1 := time.Now()
	sol1 := solve1(intervals1, false)
	dur := time.Since(t1)
	fmt.Println("Solution part 1 =", sol1, "(ET =", dur, ")")

	// ---------------------------------------------
	fmt.Println()
	fmt.Println("=== Part 2 ===")
	sol2_1_test := solve2(intervalsTest, true)
	fmt.Println("Test solution 2 =", sol2_1_test, "->", checkSolution(sol2_1_test, testSolution2))

	t1 = time.Now()
	sol2 := solve2(intervals1, false)
	dur = time.Since(t1)
	fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
