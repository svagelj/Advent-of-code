package main

import (
	rw "aoc_2024/tools"
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

func solve1(data [][]int, printout bool) int {

	minD := 1
	maxD := 3

	var isSafe int
	var sign int
	sum := 0
	var d int
	for i := range data {

		sign = 0
		isSafe = 1
		for j := 0; j < len(data[i]) - 1; j++ {

			d = data[i][j] - data[i][j+1]

			if absInt(d) < minD || absInt(d) > maxD || ((d > 0 && sign == -1) || (d < 0 && sign == 1)) {
				isSafe = 0
				break
			}

			// Get sign of the difference -> increasing or decreasing levels
			if data[i][j] > data[i][j+1] {
				sign = 1
			} else {
				sign = -1
			}
		}

		if printout == true {
			fmt.Printf("%2d: %v => %d\n", i+1, data[i], isSafe)
		}

		sum = sum + isSafe
	}

	return sum
}

func checkOneReport(_report []int, minD int, maxD int, maxMistakes int) int {

	nMistakes := 0
	sign := 0
	isSafe := 1
	var isSafeA, isSafeB int
	var sliceA, sliceB []int

	report := append([]int{}, _report...)
	fmt.Println("\tchecking", _report)

	var d int
	j := 0
	for j < len(report) - 1 {

		d = report[j] - report[j+1]

		if absInt(d) < minD || absInt(d) > maxD || ((d > 0 && sign == -1) || (d < 0 && sign == 1)) {
			nMistakes = nMistakes + 1
			fmt.Printf("\t\tmistake at j = %d, %d\n", j, report[j])

			if nMistakes > maxMistakes {
				// fmt.Println("\t\tmistakes are bad", nMistakes, maxMistakes, report)
				return 0
			} else {
				isSafeA, isSafeB = 0,0

				// Skip first element of this comparison
				fmt.Println("\t\tskip first element")
				if (j != 0){
					sliceA = append([] int {}, report[:j-1]...)
					sliceA = append(sliceA, report[j:]...)
				} else {
					sliceA = report[1:]
				}
				fmt.Println("\t\tyay1", sliceA)
				isSafeA = checkOneReport(sliceA, minD, maxD, 0)
				fmt.Println("\t   ", sliceA, "->", isSafeA)

				// Skip second element of this comparison
				if j != len(report) - 1{
					fmt.Println("\t\tskip second element a")
					sliceB = append([] int {}, report[:j]...)
					sliceB = append(sliceB, report[j+1:]...)
					// fmt.Println("\t\tyaya", sliceB)
					isSafeB = checkOneReport(sliceB, minD, maxD, 0)
					fmt.Println("\t   ", sliceB, "-->", isSafeB)
				} else {
					fmt.Println("\t\tskip second element b")
					isSafeB = checkOneReport(report[:j], minD, maxD, 0)
					fmt.Println("\t   ", report[:j], "==>", isSafeB)
				}
			}

			if isSafeA == 0 && isSafeB == 0 {
				// if isSafeA == 0 {
				// 	fmt.Println("\t\t A is bad", report)
				// }
				// if isSafeB == 0 {
				// 	fmt.Println("\t\t B is bad", report)
				// }
				return 0
			} else {
				return 1
			}
		}

		// Get sign of the difference -> increasing or decreasing levels
		if report[j] > report[j+1] {
			sign = 1
		} else {
			sign = -1
		}

		j++
	}

	return isSafe
}

func solve2(data [][]int, printout bool) int {

	minD := 1
	maxD := 3
	maxMistakes := 1

	var isSafe int
	sum := 0
	for i := range data {

		if printout == true {
			fmt.Printf("%2d: %v\n", i+1, data[i])
		}

		isSafe = checkOneReport(data[i], minD, maxD, maxMistakes)

		if printout == true {
			fmt.Printf("     => %d\n", isSafe)
		}

		sum = sum + isSafe
	}

	return sum
}

func main() {

	// data gathering and parsing
	testData := initData(testData1, true)

	fileName := "day_02_data.txt"
	fmt.Println("Reading file '" + fileName + "'")
	lines := rw.ReadFile(fileName)
	fileData := initData(lines, false)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_test := solve1(testData, true)
	fmt.Println("Test solution =", sol1_test, "->", checkSolution(sol1_test, testSolution1))

	sol1 := solve1(fileData, false)
	fmt.Println("Solution part =", sol1)

	// ---------------------------------------------
	fmt.Println("=== Part 2 ===")
	sol2_test := solve2(testData, true)
	fmt.Println("Test solution =", sol2_test, "->", checkSolution(sol2_test, testSolution2))

	// sol2 := solve2(fileData, true)
	// fmt.Println("Solution part =", sol2)

	// a := []int {1,2,3,4,5,6}
	// fmt.Println("YOLO", a[:5])
}
