package main

import (
	// Array "aoc_2024/tools/Array"
	// Math "aoc_2024/tools/Math"
	rw "aoc_2024/tools/rw"
	// "time"
	// Printer "aoc_2024/tools/Printer"
	"fmt"
	"strconv"
	"strings"
)

// var must be used for global variables
var testData1 = []string {
	"190: 10 19",
	"3267: 81 40 27",
	"83: 17 5",
	"156: 15 6",
	"7290: 6 8 6 15",
	"161011: 16 10 13",
	"192: 17 8 14",
	"21037: 9 7 18 13",
	"292: 11 6 16 20",
}

var testSolution1, testSolution2 = 3749, -1
var rotate90CwMatrix = [][]int {{0,1}, {-1,0}}

//------------------------------------------------------

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed"
	} else {
		return "failed"
	}
}

func initData(fileLines []string) ([][] int) {

	data := [][]int {}

	for i := range fileLines {
		line := fileLines[i]

		lineData := []int {}
		_line := strings.Split(line, ":")

		// Get first number - target
		_int, err := strconv.Atoi(_line[0])
		if err == nil {
			lineData = append(lineData, _int)
		} else {
			panic(err)
		}

		// Get other numbers - elements to do operations on
		_data := strings.Fields(_line[1])
		for j := range _data {
			_int, err := strconv.Atoi(_data[j])
			if err == nil {
				lineData = append(lineData, _int)
			} else {
				panic(err)
			}
		}

		data = append(data, lineData)

		// fmt.Printf("%2d: %s => %v (%d)\n", i+1, line, lineData, len(lineData))
	}

	return data

}

func doOperation(operator rune, value1 int, value2 int) int {
	if operator == '+' {
		return value1 + value2
	} else if operator == '*' {
		return value1 * value2
	} else {
		return -1
	}
}

func checkIfSolvable(target int, numbers []int, result int) bool {

	operators := []rune {'+', '*'}

	for i := range operators {
		if result == -1 {
			_result := doOperation(operators[i], numbers[0], numbers[1])
			// fmt.Printf("\tstart operator %c -> %d\n", operators[i], _result)
			
			// either call next recursion or end it
			if len(numbers) > 2 {
				if checkIfSolvable(target, numbers[2:], _result) {
					return true
				} else {
					continue
				}
			} else {
				// fmt.Println("test end:", _result == target)
				if _result == target {
					return true
				} else {
					// this operator resulted in false solvability - try next one
					continue
				}
			}

		} else {
			_result := doOperation(operators[i], result, numbers[0])
			// fmt.Printf("\toperator %c -> %d\n", operators[i], _result)

			// either call next recursion or end it
			if len(numbers) > 1 {
				if checkIfSolvable(target, numbers[1:], _result) {
					return true
				} else {
					// this operator resulted in false solvability - try next one
					continue
				} 
			} else {
				if _result == target {
					return true
				} else {
					continue
				}
			}
		}
	}

	// all operators were tested and none resulted in true solvability
	return false
}

func solve1(data [][]int, printout bool) int {

	sum := 0
	for i := range data {

		target := data[i][0]
		numbers := data[i][1:]
		isSolvable := checkIfSolvable(target, numbers, -1)

		if printout {
			fmt.Println(i, data[i], "=>", isSolvable)
		}

		if isSolvable {
			sum = sum + target
		}
	}

	return sum
}

//----------------------------------------

func solve2(data [][]rune, startingInd []int, maxSteps int, printout bool) int {

	sum := 0
	for i := range data {
		for j := range data[i] {
			sum = sum + j
		}
	}

	return sum
}

func main() {

	// data gathering and parsing
	dataTest := initData(testData1)

	fileName := "day_07_data.txt"
	fileData := rw.ReadFile(fileName)
	data := initData(fileData)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_test := solve1(dataTest, true)
	fmt.Println("Test solution 1 =", sol1_test, "->", checkSolution(sol1_test, testSolution1))

	sol1 := solve1(data, false)
	fmt.Println("Solution part 1 =", sol1)

	// ---------------------------------------------
	// fmt.Println()
	// fmt.Println("=== Part 2 ===")
	// sol2_test := solve2(dataTest, startingIndTest, 9999, true)
	// fmt.Println("Test solution 2 =", sol2_test, "->", checkSolution(sol2_test, testSolution2))

	// t1 := time.Now()
	// sol2 := solve2(data, startingInd, 99999999, false)
	// dur := time.Since(t1)
	// fmt.Println("Solution part 2 =", sol2, "(Et =", dur, ")")
}
