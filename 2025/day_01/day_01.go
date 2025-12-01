package main

import (
	Math "aoc_2025/tools/Math"
	// Array "aoc_2025/tools/Array"
	// Printer "aoc_2025/tools/Printer"
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
	"L68",
	"L30",
	"R48",
	"L5",
	"R60",
	"L55",
	"L1",
	"L99",
	"R14",
	"L82",
}


var testSolution1, testSolution2 = 3, 6

//------------------------------------------------------

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed   :)"
	} else {
		return "failed   :(   (should be '"+strconv.Itoa(solValue)+"')"
	}
}

//------------------------------------------------------

func initData(fileLines []string) ([]int) {

	rotations := []int {}

	for i := range fileLines {
		line := fileLines[i]

		val, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		if line[0] == 'L'{
			rotations = append(rotations, -val)
		} else if line[0] == 'R' {
			rotations = append(rotations, val)
		} else {
			fmt.Println("Bad rotation direction")
		}
	}

	return rotations
}

func solve1(rotations []int, startValue int, printout bool) int {

	if printout {
		fmt.Println("rotations:", rotations)
	}

	sum := 0

	minValue := 0
	maxValue := 99
	currentValue := startValue

	for i := range rotations {

		rot := rotations[i] %(maxValue-minValue + 1)

		currentValue = currentValue + rot

		if currentValue > maxValue {
			currentValue = currentValue%(maxValue-minValue + 1)
		}
		if currentValue < minValue {
			// currentValue = maxValue + currentValue%(maxValue - minValue + 1) + 1
			currentValue = maxValue + currentValue + 1
		}

		if printout {
			fmt.Println("rotation",  rotations[i], "=", rot, "->", currentValue)
		}

		if currentValue == 0 {
			sum = sum + 1
		}
	}

	return sum
}

//----------------------------------------

func solve2(rotations []int, startValue int, printout bool) int {

	if printout {
		fmt.Println("rotations:", rotations)
	}

	sum := 0

	minValue := 0
	maxValue := 99
	currentValue := startValue

	for i := range rotations {

		lastZero := currentValue == 0

		nPasses := rotations[i] / (maxValue-minValue + 1)
		sum = sum + Math.AbsInt(nPasses)

		rot := rotations[i] % (maxValue-minValue + 1)
		currentValue = currentValue + rot

		if currentValue > maxValue {
			currentValue = currentValue%(maxValue-minValue + 1)
			if currentValue != 0 && !lastZero {
				sum = sum + 1
			}
		} else if currentValue < minValue {
			currentValue = maxValue + currentValue + 1
			if currentValue != 0 && !lastZero {
				sum = sum + 1
			}
		}

		if printout {
			fmt.Println("rotation",  rotations[i], "=", rot, "->", currentValue, nPasses)
		}

		if currentValue == 0 {
			sum = sum + 1
		}
	}

	return sum
}

func main() {

	// data gathering and parsing
	rotationsTest := initData(testData)

	fileName := "day_01_data.txt"
	fileData := rw.ReadFile(fileName)
	rotations1 := initData(fileData)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_1_test := solve1(rotationsTest, 50, true)
	fmt.Println("Test solution 1 =", sol1_1_test, "->", checkSolution(sol1_1_test, testSolution1))

	t1 := time.Now()
	sol1 := solve1(rotations1, 50, false)
	dur := time.Since(t1)
	fmt.Println("Solution part 1 =", sol1, "(ET =", dur, ")")

	// ---------------------------------------------
	fmt.Println()
	fmt.Println("=== Part 2 ===")
	sol2_1_test := solve2(rotationsTest, 50, true)
	fmt.Println("Test solution 2 =", sol2_1_test, "->", checkSolution(sol2_1_test, testSolution2))

	t1 = time.Now()
	sol2 := solve2(rotations1, 50, false)
	dur = time.Since(t1)
	fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
