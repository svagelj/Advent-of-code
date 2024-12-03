package main

import (
	rw "aoc_2024/tools/rw"
	"fmt"
	"strconv"
	"strings"
)

// var must be used for global variables
var testData1 = []string{
	"7 6 4 2 1",
	"1 2 7 8 9",
	"9 7 6 2 1",
	"1 3 2 4 5",
	"8 6 4 4 1",
	"1 3 6 7 9",
}

var testSolution1, testSolution2 = 2, 4

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed"
	} else {
		return "failed"
	}
}

func initData(fileLiens []string, printout bool) [][]int {

	data := [][] int{}

	var n int
	var lineData []int
	i := 0
	for i < len(fileLiens) {
		line := fileLiens[i]

		// Get array from string, separator is empty space
		_data := strings.Fields(line)

		if printout == true {
			fmt.Printf("%2d: %s => %s\n", i+1, line, _data)
		}
		
		// Go trough _data an convert each element to int
		lineData = []int{}
		n = len(_data)
		for j := range n {
			_int, err := strconv.Atoi(_data[j])
			if err == nil {
				lineData = append(lineData, _int)
			} else {
				panic(err)
			}
		}
		data = append(data, lineData)

		i++
	}

	if printout == true {
		fmt.Println("yay", data)
	}

	return data

}

func absInt(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func checkOneReport1(report []int, minD int, maxD int, ) int {

	sign := 0
	var d int
	for j := 0; j < len(report) - 1; j++ {

		d = report[j] - report[j+1]

		if absInt(d) < minD || absInt(d) > maxD || ((d > 0 && sign == -1) || (d < 0 && sign == 1)) {
			return 0
		}

		// Get sign of the difference -> increasing or decreasing levels
		if report[j] > report[j+1] {
			sign = 1
		} else {
			sign = -1
		}
	}

	return 1
}

func solve1(data [][]int, printout bool) int {

	minD := 1
	maxD := 3

	var isSafe int
	sum := 0
	for i := range data {

		isSafe = checkOneReport1(data[i], minD, maxD)

		if printout == true {
			fmt.Printf("%2d: %v => %d\n", i+1, data[i], isSafe)
		}

		sum = sum + isSafe
	}

	return sum
}

func checkOneReport2(_report []int, minD int, maxD int) int {

	// This approach is not ok because it wrongly skips some edge cases -> "9 12 11 8 7" is deemed unsafe if one removed (remove 9), even though it is save
	// Idea is this. Go trough element by element in given report and compare it with the next one.
	// If this pair is legal, move forward by one
	// If this pair causes error, remove each of the elements in this pair and recheck again the resulting shorter report

	sign := 0
	isSafe := 1
	var isSafeA, isSafeB int
	var sliceA, sliceB []int

	report := append([]int{}, _report...)
	// fmt.Println("\tchecking", _report)

	var d int
	for j := 0; j < len(report) - 1; j++ {

		d = report[j] - report[j+1]

		if absInt(d) < minD || absInt(d) > maxD || ((d > 0 && sign == -1) || (d < 0 && sign == 1)) {
			// error in current compare pair was found
			// try to remove each element of the pair

			fmt.Printf("\t\tmistake at j = %d, (%d, %d)\n", j, report[j], report[j+1])

			// Skip first element of this comparison
			sliceA = append([] int {}, report[:j]...)
			sliceA = append(sliceA, report[j+1:]...)
			isSafeA = checkOneReport1(sliceA, minD, maxD)
			fmt.Println("\t", sliceA, "->", isSafeA)

			if isSafeA == 1 {
				return 1
			}

			// Skip second element of this comparison
			if j == len(report) - 1 {
				sliceB = report[:len(report)-1]
				isSafeB = checkOneReport1(sliceB, minD, maxD)
				fmt.Println("\t", sliceB, "__>", isSafeB)
			} else {
				sliceB = append([] int {}, report[:j+1]...)
				sliceB = append(sliceB, report[j+2:]...)
				isSafeB = checkOneReport1(sliceB, minD, maxD)
				fmt.Println("\t", sliceB, "-->", isSafeB)
			}

			if isSafeA == 1 || isSafeB == 1 {
				return 1
			} else {
				return 0
			}
		}

		// Get sign of the difference -> increasing or decreasing levels
		if report[j] > report[j+1] {
			sign = 1
		} else {
			sign = -1
		}
	}

	return isSafe
}

func solve2(data [][]int, printout bool) int {

	minD := 1
	maxD := 3

	var isSafe int
	sum := 0
	for i := range data {

		if printout == true {
			fmt.Printf("%2d: %v\n", i+1, data[i])
		}

		// Wrong approach
		// isSafe = checkOneReport2(data[i], minD, maxD)

		// Nuclear option: remove every element one by one but only if whole report is unsafe
		isSafe = checkOneReport1(data[i], minD, maxD)
		if isSafe == 0 {
			for j := 0; j < len(data[i]); j++ {
				report := append([] int {}, data[i][:j]...)
				report = append(report, data[i][j+1:]...)
				// fmt.Println("\t", report)

				_isSafe := checkOneReport1(report, minD, maxD)
				if _isSafe == 1 {
					isSafe = 1
					break
				}
			}
		}

		if printout == true {
			fmt.Printf("     => %d\n", isSafe)
		}

		sum = sum + isSafe
	}

	return sum
}

func main() {

	// data gathering and parsing
	testData := initData(testData1, false)

	fileName := "day_02_data.txt"
	fmt.Println("Reading file '" + fileName + "'")
	lines := rw.ReadFile(fileName)
	fileData := initData(lines, false)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_test := solve1(testData, true)
	fmt.Println("Test solution =", sol1_test, "->", checkSolution(sol1_test, testSolution1))

	sol1 := solve1(fileData, false)
	fmt.Println("Solution part 1 =", sol1)

	// ---------------------------------------------
	fmt.Println("=== Part 2 ===")
	sol2_test := solve2(testData, true)
	fmt.Println("Test solution =", sol2_test, "->", checkSolution(sol2_test, testSolution2))

	sol2 := solve2(fileData, false)
	fmt.Println("Solution part 2 =", sol2)
}
