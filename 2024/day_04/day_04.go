package main

import (
	rw "aoc_2024/tools/rw"
	Math "aoc_2024/tools/Math"
	"fmt"
	"strings"
)

// var must be used for global variables
var testData1 = []string {
	"MMMSXXMASM",
	"MSAMXMSMSA",
	"AMXSXMAAMM",
	"MSAMASMSMX",
	"XMASAMXAMM",
	"XXAMMXXAMA",
	"SMSMSASXSS",
	"SAXAMASAAA",
	"MAMMMXMMMM",
	"MXMXAXMASX",
}

var testSolution1, testSolution2 = 18, 9

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed"
	} else {
		return "failed"
	}
}

func findTarget1(data []string, i int, j int, target string) int {

	n := len(data)
	m := len(data[i])
	targetR := rw.ReverseString(target)
	N := 0

	// fmt.Println(i,j, target)

	// check horizontal
	maxJ := Math.MinInt(j+len(target), m)
	// fmt.Println("\t", data[i][j:maxJ], strings.Contains(data[i][j:maxJ], target))
	// fmt.Println("\t", data[i][j:maxJ], strings.Contains(data[i][j:maxJ], targetR))
	if strings.Contains(data[i][j:maxJ], target) == true {
		N++
	}
	if strings.Contains(data[i][j:maxJ], targetR) == true {
		N++
	}

	// check vertical
	maxI := Math.MinInt(i+len(target), n)
	vertStr := ""
	for k := range len(target) {
		if k+i >= maxI {
			break
		}
		vertStr = vertStr + string(data[i+k][j])
	}
	// fmt.Println("\t", vertStr, strings.Contains(vertStr, target))
	// fmt.Println("\t", vertStr, strings.Contains(vertStr, targetR))
	if strings.Contains(vertStr, target) == true {
		N++
	}
	if strings.Contains(vertStr, targetR) == true {
		N++
	}

	// diagonal +,+
	diag1 := ""
	for k := range len(target) {
		if k+i >= maxI || k+j >= maxJ {
			break
		}
		diag1 = diag1 + string(data[i+k][j+k])
	}
	// fmt.Println("\t", diag1, strings.Contains(diag1, target))
	// fmt.Println("\t", diag1, strings.Contains(diag1, targetR))
	if strings.Contains(diag1, target) == true {
		N++
	}
	if strings.Contains(diag1, targetR) == true {
		N++
	}

	// diagonal +,-
	diag2 := ""
	for k := range len(target) {
		if i+k >= maxI || j-k < 0 {
			break
		}
		diag2 = diag2 + string(data[i+k][j-k])
	}
	// fmt.Println("\t", diag2, strings.Contains(diag2, target))
	// fmt.Println("\t", diag2, strings.Contains(diag2, targetR))
	if strings.Contains(diag2, target) == true {
		N++
	}
	if strings.Contains(diag2, targetR) == true {
		N++
	}

	return N
}

func solve1(data []string, target string, printout bool) int {

	sum := 0
	for i := range data {
		// fmt.Println("yay:", data[i])

		n := 0
		for j := range data[i] {
			n = findTarget1(data, i, j, target)
			sum = sum + n
			if printout && n > 0 {
				fmt.Println("found at:", i,j, "n =", n)
			}
		}

		// break
	}

	return sum
}

//----------------------------------------

func findTarget2(data []string, i int, j int, target string) int {

	n := len(data)
	m := len(data[i])
	targetR := rw.ReverseString(target)
	N := 0

	// fmt.Println(i,j, target)

	// check horizontal
	maxJ := Math.MinInt(j+len(target), m)
	// fmt.Println("\t", data[i][j:maxJ], strings.Contains(data[i][j:maxJ], target))
	// fmt.Println("\t", data[i][j:maxJ], strings.Contains(data[i][j:maxJ], targetR))
	if strings.Contains(data[i][j:maxJ], target) == true {
		N++
	}
	if strings.Contains(data[i][j:maxJ], targetR) == true {
		N++
	}

	// check vertical
	maxI := Math.MinInt(i+len(target), n)
	vertStr := ""
	for k := range len(target) {
		if k+i >= maxI {
			break
		}
		vertStr = vertStr + string(data[i+k][j])
	}
	// fmt.Println("\t", vertStr, strings.Contains(vertStr, target))
	// fmt.Println("\t", vertStr, strings.Contains(vertStr, targetR))
	if strings.Contains(vertStr, target) == true {
		N++
	}
	if strings.Contains(vertStr, targetR) == true {
		N++
	}

	// diagonal +,+
	diag1 := ""
	for k := range len(target) {
		if k+i >= maxI || k+j >= maxJ {
			break
		}
		diag1 = diag1 + string(data[i+k][j+k])
	}
	// fmt.Println("\t", diag1, strings.Contains(diag1, target))
	// fmt.Println("\t", diag1, strings.Contains(diag1, targetR))
	if strings.Contains(diag1, target) == true {
		N++
	}
	if strings.Contains(diag1, targetR) == true {
		N++
	}

	// diagonal +,-
	diag2 := ""
	for k := range len(target) {
		if i+k >= maxI || j-k < 0 {
			break
		}
		diag2 = diag2 + string(data[i+k][j-k])
	}
	// fmt.Println("\t", diag2, strings.Contains(diag2, target))
	// fmt.Println("\t", diag2, strings.Contains(diag2, targetR))
	if strings.Contains(diag2, target) == true {
		N++
	}
	if strings.Contains(diag2, targetR) == true {
		N++
	}

	return N
}

func solve2(data []string, target string, printout bool) int {

	sum := 0
	for i := range data {
		// fmt.Println("yay:", data[i])

		n := 0
		for j := range data[i] {
			n = findTarget2(data, i, j, target)
			sum = sum + n
			if printout && n > 0 {
				fmt.Println("found at:", i,j, "n =", n)
			}
		}

		// break
	}

	return sum
}

func main() {

	// data gathering and parsing
	fileName := "day_04_data.txt"
	fmt.Println("Reading file '" + fileName + "'")
	fileData := rw.ReadFile(fileName)

	target := "XMAS"

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_test := solve1(testData1, target, true)
	fmt.Println("Test solution =", sol1_test, "->", checkSolution(sol1_test, testSolution1))

	sol1 := solve1(fileData, target, false)
	fmt.Println("Solution part 1 =", sol1)

	// ---------------------------------------------
	// fmt.Println("=== Part 2 ===")
	// sol2_test := solve2(testData1, "MAS", true)
	// fmt.Println("Test solution =", sol2_test, "->", checkSolution(sol2_test, testSolution2))

	// sol2 := solve2(fileData, false)
	// fmt.Println("Solution part 2 =", sol2)
}
