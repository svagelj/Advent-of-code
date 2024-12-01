package main

import (
	rw "aoc_2024/tools"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// var must be used for global variables
var testData1 = []string{
	"3   4",
	"4   3",
	"2   5",
	"1   3",
	"3   9",
	"3   3",
}

var testSolution1, testSolution2 = 11, 31

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed"
	} else {
		return "failed"
	}
}

func initData(fileLiens []string, printout bool) [][]int {

	column1 := []int{}
	column2 := []int{}

	i := 0
	for i < len(fileLiens) {
		line := fileLiens[i]

		// Parse the two columns
		_data := strings.Fields(line)

		if printout == true {
			fmt.Printf("%2d: %s => %s, %s\n", i+1, line, _data[0], _data[1])
		}

		_int, err := strconv.Atoi(_data[0])
		if err == nil {
			column1 = append(column1, _int)
		} else {
			panic(err)
		}

		_int, err = strconv.Atoi(_data[1])
		if err == nil {
			column2 = append(column2, _int)
		} else {
			panic(err)
		}

		i++
	}

	if len(column1) != len(column2) {
		println("[ERROR] Columns are not the same length:", len(column1), ",", len(column2))
	}

	return [][]int{column1, column2}

}

func solve1(data [][]int, printout bool) int {

	_column1, _column2 := data[0], data[1]

	// copy to new array and sort only new arrays
	column1 := append([]int{}, _column1...)
	column2 := append([]int{}, _column2...)
	sort.Ints(column1[:])
	sort.Ints(column2[:])

	sum := 0
	var d int
	for i := 0; i < len(column1); i++ {

		d = column1[i] - column2[i]
		if d < 0 {
			d = -d
		}

		if printout == true {
			fmt.Printf("%2d: %d, %d => %d\n", i+1, column1[0], column2[1], d)
		}

		sum = sum + d
	}

	return sum
}

func solve2(data [][]int, printout bool) int {

	column1, column2 := data[0], data[1]

	N := len(column1)
	sum := 0
	var n int
	for i := 0; i < N; i++ {

		n = 0
		for j := 0; j < N; j++ {
			if column1[i] == column2[j] {
				n++
			}
		}

		if printout == true {
			fmt.Printf("%2d: %d, %d => %d\n", i+1, column1[i], column2[i], n*column1[i])
		}

		sum = sum + n*column1[i]
	}

	return sum
}

func main() {

	// data gathering and parsing
	testData := initData(testData1, false)

	fileName := "day_01_data.txt"
	fmt.Println("Reading file '" + fileName + "'")
	lines := rw.ReadFile(fileName)
	fileData := initData(lines, false)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_test := solve1(testData, false)
	fmt.Println("Test solution =", sol1_test, "->", checkSolution(sol1_test, testSolution1))

	sol1 := solve1(fileData, false)
	fmt.Println("Solution part =", sol1)

	// ---------------------------------------------
	fmt.Println("=== Part 2 ===")
	sol2_test := solve2(testData, false)
	fmt.Println("Test solution =", sol2_test, "->", checkSolution(sol2_test, testSolution2))

	sol2 := solve2(fileData, false)
	fmt.Println("Solution part =", sol2)
}
