package main

import (
	// Math "aoc_2025/tools/Math"
	// Array "aoc_2025/tools/Array"
	Printer "aoc_2025/tools/Printer"
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
	"123 328  51 64 " ,
	" 45 64  387 23 ",
	"  6 98  215 314",
	"*   +   *   +  " ,
}

var testSolution1, testSolution2 = 4277556, 3263827

//------------------------------------------------------

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed   :)"
	} else {
		return "failed   :(   (should be '"+strconv.Itoa(solValue)+"')"
	}
}

//------------------------------------------------------

func initData(fileLines []string) ([][]int, []string) {

	values := [][]int {}
	operations := []string {}

	for i := range fileLines {
		line := fileLines[i]

		if line[0] == '*' || line[0] == '+' {
			_opers := strings.Split(line, " ")
			for j := range _opers {
				if _opers[j] != "" {
					operations = append(operations, _opers[j])
				}
			}
		} else {
			_line := []int {}

			_num := strings.Split(line, " ")
			for j := range _num {

				// There is white space between values -> trim it
				if _num[j] != "" {
					value, err := strconv.Atoi(_num[j])
					if err != nil {
						panic(err)
					}
					_line = append(_line, value)
				}
			}
			values = append(values, _line)
		}
	}

	return values, operations
}

func solve1(values [][]int, operations []string, printout bool) int {

	if printout {
		fmt.Println("values", values)
		fmt.Println("operations", operations)
	}

	sum := 0
	for j := range values[0] {

		oper := operations[j]
		var result int
		if oper == "+" {
			result = 0
			for i:= range values {
				result = result + values[i][j]
			}
		} else if oper == "*" {
			result = 1
			for i:= range values {
				result = result * values[i][j]
			}
		}

		if printout {
			fmt.Println(j, oper, "->", result)
		}

		sum = sum + result
	}

	return sum
}

//----------------------------------------
func initData2(fileLines []string) ([][]string) {

	data := [][]string {}

	for i := range fileLines {
		line := fileLines[i]
		
		_line := []string {}
		for _, v := range line {
			_line = append(_line, string(v))
		}
		data = append(data, _line)
	}

	return data
}

func solve2(data [][]string, printout bool) int {

	if printout {
		Printer.PrintGridStr(data, 1)
	}

	maxY := len(data)

	sum := 0

	numList := []int {}
	// loop from right to left
	for j := len(data[0]) - 1; j >= 0; j-- {
		
		currNumStr := ""
		// loop from top to bottom
		for i := 0; i < len(data)-1; i++ {
			if data[i][j] != " " {
				currNumStr = currNumStr + data[i][j]
			}
		}

		value, err := strconv.Atoi(currNumStr)
		if err != nil {
			panic(err)
		}
		numList = append(numList, value)
		
		// sum if last element is operation char
		oper := data[maxY-1][j]
		if oper == "*" || oper == "+" {
			var result int
			if oper == "*" {
				result = 1
				for _,num := range numList {
					result = result * num
				}
			} else if oper == "+" {
				result = 0
				for _,num := range numList {
					result = result + num
				}
			}

			if printout {
				fmt.Println(oper, numList, "->", result)
			}

			sum = sum + result
			numList = []int {}
			j = j - 1	// skip empty column
		}
	}

	return sum
}

func main() {

	// data gathering and parsing
	valuesTest1, operationsTest1 := initData(testData)

	fileName := "day_06_data.txt"
	fileData := rw.ReadFile(fileName)
	values1, operations1 := initData(fileData)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_1_test := solve1(valuesTest1, operationsTest1, true)
	fmt.Println("Test solution 1 =", sol1_1_test, "->", checkSolution(sol1_1_test, testSolution1))

	t1 := time.Now()
	sol1 := solve1(values1, operations1, false)
	dur := time.Since(t1)
	fmt.Println("Solution part 1 =", sol1, "(ET =", dur, ")")

	// ---------------------------------------------
	// fmt.Println("-----------------------")
	dataTest2 := initData2(testData)
	data2 := initData2(fileData)

	fmt.Println()
	fmt.Println("=== Part 2 ===")
	sol2_1_test := solve2(dataTest2, true)
	fmt.Println("Test solution 2 =", sol2_1_test, "->", checkSolution(sol2_1_test, testSolution2))

	t1 = time.Now()
	sol2 := solve2(data2, false)
	dur = time.Since(t1)
	fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
