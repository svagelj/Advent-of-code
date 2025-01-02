package main

import (
	// Array "aoc_2024/tools/Array"
	// "sort"
	// Math "aoc_2024/tools/Math"
	rw "aoc_2024/tools/rw"
	"time"
	// Printer "aoc_2024/tools/Printer"
	"fmt"
	"strconv"
	"strings"
)

// var must be used for global variables
var testData1 = []string {
	// "0 1 10 99 999",
	"125 17",
}

var testSolution1, testSolution2 = 55312, 55312

//------------------------------------------------------

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed   :)"
	} else {
		return "failed   :(   (should be '"+strconv.Itoa(solValue)+"')"
	}
}

//------------------------------------------------------

func initData(fileLines []string) []int {

	data := []int{}

	for i := range fileLines {
		line := fileLines[i]
		_line := strings.Split(line, " ")

		lineData := []int{}
		for j := range _line {
			_int, err := strconv.Atoi(_line[j])
			if err == nil {
				lineData = append(lineData, _int)
			} else {
				panic(err)
			}
		}
		data = lineData
		break
	}

	return data
}

func changeStone1(stoneNumber int) []int {

	// to get number of digits
	numberStr := strconv.Itoa(stoneNumber)

	if stoneNumber == 0 {
		return []int {1}
	} else if len(numberStr) % 2 == 0 {
		left := numberStr[:len(numberStr)/2]
		right := numberStr[len(numberStr)/2:]
		// fmt.Println(numberStr, len(numberStr), "=>", left, right)

		leftInt, err := strconv.Atoi(left)
		if err != nil {
			panic(err)
		}
		rightInt, err := strconv.Atoi(right)
		if err != nil {
			panic(err)
		}

		return []int {leftInt, rightInt}
	} else {
		return []int {stoneNumber*2024}
	}

}

func solve1(data []int, maxBlinks int, printout bool) int {

	state := make([]int, len(data))
	copy(state, data)

	for i := range maxBlinks {

		newState := []int{}
		for k := range state {
			change := changeStone1(state[k])
			newState = append(newState, change...)
			// fmt.Println("  change", state[k], "=>", change)
		}
		// fmt.Println(i, state, "=>", newState)
		if printout {
			fmt.Println(i+1, "=>", len(newState))
		}

		state = make([]int, len(newState))
		copy(state, newState)
	}
	
	sum := len(state)
	return sum
}

//----------------------------------------

func solve2(data []int, maxBlinks int, printout bool) int {

	// State for this part is defined differently 
	// state is a map[stoneNumber]: numberOfStones and iterate of this state (stoneNumbers)
	// This way we process just stone numbers that are different from each other
	// and prevent processing same stone numbers multiple times at one state change
	state := make(map[int](int))
	for i := range data {
		_, found := state[data[i]]
		if !found {
			state[data[i]] = 1
		} else {
			state[data[i]]++
		}
	}

	if printout {
		fmt.Println("starting state:", state)
	}

	for i := range maxBlinks {

		newState := make(map[int](int))
		for key, N := range state {
			newStones := changeStone1(key)

			// append transformed stones with frequencies to newState
			for j := range newStones {
				newStone := newStones[j]
				_, found := newState[newStone]
				if !found {
					newState[newStone] = N
				} else {
					newState[newStone] = newState[newStone] + N
				}
			}
		}
		if printout {
			_sum := 0
			for _, value := range newState {
				_sum = _sum + value
			}
			fmt.Println(i+1, "=>", _sum)
		}

		state = newState
	}
	
	sum := 0
	for _, value := range state {
		sum = sum + value
	}
	return sum
}

func main() {

	// data gathering and parsing
	dataTest := initData(testData1)

	fileName := "day_11_data.txt"
	fileData := rw.ReadFile(fileName)
	data := initData(fileData)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_test := solve1(dataTest, 25, true)
	fmt.Println("Test solution 1 =", sol1_test, "->", checkSolution(sol1_test, testSolution1))

	sol1 := solve1(data, 25, false)
	fmt.Println("Solution part 1 =", sol1)

	// ---------------------------------------------
	fmt.Println()
	fmt.Println("=== Part 2 ===")
	sol2_test := solve2(dataTest, 25, false)
	fmt.Println("Test solution 2 =", sol2_test, "->", checkSolution(sol2_test, testSolution2))

	t1 := time.Now()
	sol2 := solve2(data, 75, false)
	dur := time.Since(t1)
	fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
