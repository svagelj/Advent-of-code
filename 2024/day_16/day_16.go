package main

import (
	Array "aoc_2024/tools/Array"
	// "bufio"
	// "os"
	"strings"

	// Math "aoc_2024/tools/Math"
	rw "aoc_2024/tools/rw"
	"time"

	Printer "aoc_2024/tools/Printer"
	"fmt"
	"strconv"
	"sort"
)

// var must be used for global variables
var testData1 = []string {
	"###############",
	"#.......#....E#",
	"#.#.###.#.###.#",
	"#.....#.#...#.#",
	"#.###.#####.#.#",
	"#.#.#.......#.#",
	"#.#.#####.###.#",
	"#...........#.#",
	"###.#.#####.#.#",
	"#...#.....#.#.#",
	"#.#.#.###.#.#.#",
	"#.....#...#.#.#",
	"#.###.#.#.#.#.#",
	"#S..#.....#...#",
	"###############",
}

var testData2 = []string {
	"#################",
	"#...#...#...#..E#",
	"#.#.#.#.#.#.#.#.#",
	"#.#.#.#...#...#.#",
	"#.#.#.#.###.#.#.#",
	"#...#.#.#.....#.#",
	"#.#.#.#.#.#####.#",
	"#.#...#.#.#.....#",
	"#.#.#####.#.###.#",
	"#.#.#.......#...#",
	"#.#.###.#####.###",
	"#.#.#...#.....#.#",
	"#.#.#.#####.###.#",
	"#.#.#.........#.#",
	"#.#.#.#########.#",
	"#S#.............#",
	"#################",
}

var testSolution1_1, testSolution1_2, testSolution2 = 7036, 11048, -1

//------------------------------------------------------

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed   :)"
	} else {
		return "failed   :(   (should be '"+strconv.Itoa(solValue)+"')"
	}
}

//------------------------------------------------------

func initData(fileLines []string) ([][]rune, [2]int, [2]int) {

	data := [][]rune {}
	startInd := [2]int {-1,-1}
	endInd := [2]int {-1,-1}

	startChar := "S"
	endChar := "E"

	for i := range fileLines {
		line := fileLines[i]

		data = append(data, []rune(line))

		// get starting position
		if strings.Contains(line, startChar) {
			startInd[0] = i
			startInd[1] = Array.GetIndexString(line, startChar)
		}

		// get end position
		if strings.Contains(line, endChar) {
			endInd[0] = i
			endInd[1] = Array.GetIndexString(line, endChar)
		}
	}

	return data, startInd, endInd
}

func sortQueue(queue [][]int) {

	priceIndex := 4
	
	sort.Slice(queue, func(i, j int) bool {
		// edge cases - useless for this but hey why not
		if len(queue[i]) == 0 && len(queue[j]) == 0 {
			return false // two empty slices - so one is not less than other i.e. false
		}
		if len(queue[i]) == 0 || len(queue[j]) == 0 {
			return len(queue[i]) == 0 // empty slice listed "first" (change to != 0 to put them last)
		}

		// both slices len() > 0, so can test this now:
		return queue[i][priceIndex] < queue[j][priceIndex]
	})
}

func newStateInVisited(state []int, visitedLocation [][]int) bool {

	if len(visitedLocation) == 1 && len(visitedLocation[0]) == 0 {
		return false
	} else {
		// checking if direction at current point is already saved in visited
		for i := range visitedLocation {
			if state[2] == visitedLocation[i][0] && state[3] == visitedLocation[i][1] {
				return true
			}
		}
	}

	return false
}

func addToQueue(queue [][]int, data [][]rune, stepPrice int, turnPrice int, currState []int, visited [][][][]int) [][]int {

	y,x, dy,dx, currPrice := currState[0], currState[1], currState[2], currState[3], currState[4]

	// add right
	if data[y][x+1] != '#' {
		newPrice := stepPrice
		if dy == 0 && dx == -1 {
			newPrice = newPrice + 2*turnPrice
		} else if dx == 0 && (dy == 1 || dy == -1) {
			newPrice = newPrice + turnPrice
		}

		newState := []int {y,x+1, 0,1, currPrice+newPrice}
		if !newStateInVisited(newState, visited[y][x+1]) {
			queue = append(queue, newState)
			// fmt.Println("add right", newPrice, newState)
			if len(visited[y][x+1]) == 1 && len(visited[y][x+1][0]) == 0 {
				visited[y][x+1] = [][]int {{0,1, currPrice+newPrice}}
			} else {
				visited[y][x+1] = append(visited[y][x+1], []int {0,1, currPrice+newPrice})
			}
		}
	}

	// add left
	if data[y][x-1] != '#' {
		newPrice := stepPrice
		if dy == 0 && dx == 1 {
			newPrice = newPrice + 2*turnPrice
		} else if dx == 0 && (dy == 1 || dy == -1) {
			newPrice = newPrice + turnPrice
		}
		
		newState := []int {y,x-1, 0,-1, currPrice+newPrice}
		if !newStateInVisited(newState, visited[y][x-1]) {
			queue = append(queue, newState)
			// fmt.Println("add left", newPrice, newState)
			if len(visited[y][x-1]) == 1 && len(visited[y][x-1][0]) == 0 {
				visited[y][x-1] = [][]int {{0,-1, currPrice+newPrice}}
			} else {
				visited[y][x-1] = append(visited[y][x-1], []int {0,-1, currPrice+newPrice})
			}
		}
	}

	// add down
	if data[y+1][x] != '#' {
		newPrice := stepPrice
		if dy == -1 && dx == 0 {
			newPrice = newPrice + 2*turnPrice
		} else if dy == 0 && (dx == 1 || dx == -1) {
			newPrice = newPrice + turnPrice
		}

		newState := []int {y+1,x, 1,0, currPrice+newPrice}
		if !newStateInVisited(newState, visited[y+1][x]) {
			queue = append(queue, newState)
			// fmt.Println("add down", newPrice, newState)
			if len(visited[y+1][x]) == 1 && len(visited[y+1][x][0]) == 0 {
				visited[y+1][x] = [][]int {{1,0, currPrice+newPrice}}
			} else {
				visited[y+1][x] = append(visited[y+1][x], []int {1,0, currPrice+newPrice})
			}
		}
	}

	// add up
	if data[y-1][x] != '#' {
		newPrice := stepPrice
		if dy == 1 && dx == 0 {
			newPrice = newPrice + 2*turnPrice
		} else if dy == 0 && (dx == 1 || dx == -1) {
			newPrice = newPrice + turnPrice
		}

		newState := []int {y-1,x, -1,0, currPrice+newPrice}
		if !newStateInVisited(newState, visited[y-1][x]) {
			queue = append(queue, newState)
			// fmt.Println("add up", newPrice, newState)
			if len(visited[y-1][x]) == 1 && len(visited[y-1][x][0]) == 0 {
				visited[y-1][x] = [][]int {{-1,0, currPrice+newPrice}}
			} else {
				visited[y-1][x] = append(visited[y-1][x], []int {-1,0, currPrice+newPrice})
			}
		}
	}

	// Sort the queue in regards to the price - the last element of each state
	sortQueue(queue)

	return queue
}

func stepping1(data [][]rune, startInd [2]int, startDir [2]int, endInd [2]int) int {

	queue := [][]int {{startInd[0], startInd[1], startDir[0], startDir[1], 0}}
	popIndex := 0

	turnPrice := 1000
	stepPrice := 1

	visitedSimple := Array.InitArrayValuesInt(len(data), len(data[0]), -1)
	visitedSimple[startInd[0]][startInd[1]] = 0

	visited := Array.InitArrayValuesInt4D(len(data), len(data[0]))
	visited[startInd[0]][startInd[1]] = [][]int {{startDir[0], startDir[1], 0}}

	k := 0
	for k = range 999999 {

		// fmt.Println("yay", queue, popIndex)
		currState := queue[popIndex]
		y,x, _,_, currPrice := currState[0], currState[1], currState[2], currState[3], currState[4]
		queue = append(queue[:popIndex], queue[popIndex+1:]...)

		// Check for exit condition
		if y == endInd[0] && x == endInd[1] {
			// fmt.Println("Reached finish point.", y,x, k, "|", currPrice)
			return currPrice
		}

		// fmt.Println("\nstep", k, len(queue), "|", y,x, dy,dx, currPrice)
		// fmt.Println("visited:", visited[y][x])

		queue = addToQueue(queue, data, stepPrice, turnPrice, currState, visited)

		// if k > 500 {
		// 	break
		// }
	}

	fmt.Println("number of steps:", k)
	return -1
}

func solve1(data [][]rune, startInd [2]int, endInd [2]int, printout bool) int {

	if printout {
		Printer.PrintGridRune(data, 2)
		fmt.Println("start", startInd)
		fmt.Println("end", endInd)
	}

	dataCopy := Array.CopyRune2D(data)
	dir := [2]int {0,1}
	currPos := [2]int {startInd[0], startInd[1]}
	endPos := [2]int {endInd[0], endInd[1]}

	score := stepping1(dataCopy, currPos, dir, endPos)

	sum := score
	return sum
}

//----------------------------------------

func solve2(data [][]rune, moves string, printout bool) int {

	sum := 0
	
	return sum
}

func main() {

	// data gathering and parsing
	dataTest1, startTest1, endTest1 := initData(testData1)
	dataTest2, startTest2, endTest2 := initData(testData2)

	fileName := "day_16_data.txt"
	fileData := rw.ReadFile(fileName)
	data, start, end := initData(fileData)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_test := solve1(dataTest1, startTest1, endTest1, true)
	fmt.Println("Test solution 1_1 =", sol1_test, "->", checkSolution(sol1_test, testSolution1_1))
	sol2_test := solve1(dataTest2, startTest2, endTest2, false)
	fmt.Println("Test solution 1_2 =", sol2_test, "->", checkSolution(sol2_test, testSolution1_2))

	t1 := time.Now()
	sol1 := solve1(data, start, end, false)
	dur := time.Since(t1)
	fmt.Println("Solution part 1 =", sol1, "(ET =", dur, ")")

	// ---------------------------------------------
	// fmt.Println()
	// fmt.Println("=== Part 2 ===")
	// sol2_test := solve2(dataTest1_2, movesTest1_2, true)
	// fmt.Println("Test solution 2 =", sol2_test, "->", checkSolution(sol2_test, testSolution2))

	// t1 = time.Now()
	// sol2 := solve2(data, Nx, Ny, 10000, false)
	// dur = time.Since(t1)
	// fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
