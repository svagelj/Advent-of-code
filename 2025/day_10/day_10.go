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
	// "math"
	// "bufio"
	// "os"
)

// var must be used for global variables
var testData = []string {
	"[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
	"[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}",
	"[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}",
}

var testSolution1, testSolution2 = 7, -1

//------------------------------------------------------

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed   :)"
	} else {
		return "failed   :(   (should be '"+strconv.Itoa(solValue)+"')"
	}
}

//------------------------------------------------------

func initData(fileLines []string) ([]string, [][][]int, [][]int) {

	states := []string {}
	buttons := [][][]int {}
	voltages := [][]int {}

	for i := range fileLines {
		line := fileLines[i]
		_line := strings.Split(line, " ")

		// first element is state
		states = append(states, _line[0][1:len(_line[0])-1])

		// last element is voltages
		lastInd := len(_line)-1
		_volts := strings.Split(_line[lastInd][1:len(_line[lastInd])-1], ",")
		volts := []int {}
		for _,v := range _volts {
			value, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			volts = append(volts, value)
		}
		voltages = append(voltages, volts)

		// middle elements are buttons
		btns := _line[1:lastInd]
		button := [][]int {}
		for _,buttonStr := range btns {
			
			_button := strings.Split(buttonStr[1:len(buttonStr)-1], ",")
			leds := []int {}
			for _, led := range _button {
				value, err := strconv.Atoi(led)
				if err != nil {
					panic(err)
				}
				leds = append(leds, value)
			}
			button = append(button, leds)
		}
		buttons = append(buttons, button)
	}
	
	// fmt.Println("voltages", voltages)

	return states, buttons, voltages
}

func pressButton(state string, button []int) string {

	_state := []rune {}
	for _,v := range state {
		_state = append(_state, v)
	}

	for _,ind := range button {
		if _state[ind] == '.' {
			_state[ind] = '#'
		} else {
			_state[ind] = '.'
		}
	}

	return string(_state)
}

func searchLeastPresses(stateGoal string, buttons [][]int) int {

	visited := []string {}
	queue := []string {""}
	depthQ := []int {0}
	for range stateGoal {
		queue[0] = queue[0] + "."
	}

	maxIter := 999
	k := 0
	for k < maxIter && len(queue) > 0 {

		state := queue[0]
		queue = queue[1:]
		depth := depthQ[0]
		depthQ = depthQ[1:]

		// press all the buttons and add results to queue
		for _, button := range buttons {
			newState := pressButton(state, button)

			// Check exit condition
			if stateGoal == newState {
				return depth+1
			}

			// if new state was already visited, there is a better path to this new state -> skip it here
			if !slices.Contains(visited, newState) {
				queue = append(queue, newState)
				depthQ = append(depthQ, depth+1)
				visited = append(visited, newState)
			}
		}

		k = k + 1

	}

	fmt.Println("ERROR Max iteration steps was reached without finding goal state!")
	return -999
}

func solve1(states []string, buttonsArray [][][]int, printout bool) int {

	if printout {
		fmt.Println("states", states)
		fmt.Println("buttons", buttonsArray)
	}

	n := 0

	if len(states) != len(buttonsArray) {
		panic("Len of states != len of buttons array")
	}

	// loop over all cases (machines)
	for i := range states {

		stateGoal := states[i]
		buttons := buttonsArray[i]

		if printout {
			fmt.Println()
			fmt.Println(i, stateGoal, buttons)
		}

		// Breadth first search
		k := searchLeastPresses(stateGoal, buttons)
		n = n + k

		if printout {
			fmt.Println("   ->", k)
		}

	}
	
	return n
}

//----------------------------------------

func solve2(nodes [][]int, printout bool) int {

	if printout {
		fmt.Println("nodes", nodes)
	}

	maxS := -1
	return maxS
}

func main() {

	// data gathering and parsing
	statesTest1, buttonsTest1, _ := initData(testData)

	fileName := "day_10_data.txt"
	fileData := rw.ReadFile(fileName)
	states1, buttons1, _ := initData(fileData)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_1_test := solve1(statesTest1, buttonsTest1, true)
	fmt.Println("Test solution 1 =", sol1_1_test, "->", checkSolution(sol1_1_test, testSolution1))

	t1 := time.Now()
	sol1 := solve1(states1, buttons1, false)
	dur := time.Since(t1)
	fmt.Println("Solution part 1 =", sol1, "(ET =", dur, ")")

	// ---------------------------------------------

	// fmt.Println()
	// fmt.Println("=== Part 2 ===")
	// sol2_1_test := solve2(nodesTest1, true)
	// fmt.Println("Test solution 2 =", sol2_1_test, "->", checkSolution(sol2_1_test, testSolution2))

	// t1 = time.Now()
	// sol2 := solve2(diagram1, start1, false)
	// dur = time.Since(t1)
	// fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
