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
	"sort"
	"strconv"
)

// var must be used for global variables
var testData = []string {
	"5,4",
	"4,2",
	"4,5",
	"3,0",
	"2,1",
	"6,3",
	"2,4",
	"1,5",
	"0,6",
	"3,3",
	"2,6",
	"5,1",
	"1,2",
	"5,5",
	"2,5",
	"6,5",
	"1,4",
	"0,4",
	"6,4",
	"1,1",
	"6,1",
	"1,0",
	"0,5",
	"1,6",
	"2,0",
}

var testSolution1, testSolution2 = 22, "6,1"

//------------------------------------------------------

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed   :)"
	} else {
		return "failed   :(   (should be '"+strconv.Itoa(solValue)+"')"
	}
}

func checkSolutionStr(testValue string, solValue string) string {
	if testValue == solValue {
		return "passed   :)"
	} else {
		return "failed   :(   (should be '"+solValue+"')"
	}
}

//------------------------------------------------------

func initData(fileLines []string) [][2]int {

	data := [][2]int {}

	for i := range fileLines {
		line := fileLines[i]

		_line := strings.Split(line, ",")

		
		x, errX := strconv.Atoi(_line[0])
		y, errY := strconv.Atoi(_line[1])
		if errX == nil && errY == nil {
			data = append(data, [2]int {x,y})
		} else {
			panic("so sad")
		}
	}

	return data
}

func createMaze(data [][2]int, Nx int, Ny int, iMax int) [][]rune {

	maze := Array.InitArrayValuesRune(Nx, Ny, '.')

	for i := 0; i < len(data) && i < iMax; i++ {
		maze[data[i][1]][data[i][0]] = '#'
	}

	return maze
}

func sortQueue(queue [][]int, priceIndex int) {
	
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

func findShortestPath(maze [][]rune, startInd [2]int, endInd [2]int) int {

	mazeCopy := Array.CopyRune2D(maze)
	queue := [][]int {{startInd[0], startInd[1], 0}}
	popIndex := 0

	Ny, Nx := len(mazeCopy), len(mazeCopy[0])
	visited := Array.InitArrayValuesInt(len(mazeCopy), len(mazeCopy[0]), 0)

	k := 0
	for k = range 999999 {
		if len(queue) == 0 {
			// fmt.Println("Empty queue!", k)
			return -2
		}

		// fmt.Println("yay", queue, popIndex)
		currState := queue[popIndex]
		y,x, currPrice := currState[0], currState[1], currState[2]
		queue = append(queue[:popIndex], queue[popIndex+1:]...)

		// Check for exit condition
		if y == endInd[0] && x == endInd[1] {
			// fmt.Println("Reached finish point.", y,x, k, "|", currPrice)
			return currPrice
		}

		// add right
		if x+1 < Nx && mazeCopy[y][x+1] != '#' && visited[y][x+1] == 0 {
			newState := []int {y,x+1, currPrice+1}
			queue = append(queue, newState)
			visited[y][x+1]++
		}
		// add left
		if x-1 >= 0 && mazeCopy[y][x-1] != '#' && visited[y][x-1] == 0 {
			newState := []int {y,x-1, currPrice+1}
			queue = append(queue, newState)
			visited[y][x-1]++
		}
		// add down
		if y+1 < Ny && mazeCopy[y+1][x] != '#' && visited[y+1][x] == 0 {
			newState := []int {y+1,x, currPrice+1}
			queue = append(queue, newState)
			visited[y+1][x]++
		}
		// add up
		if y-1 >= 0 && mazeCopy[y-1][x] != '#' && visited[y-1][x] == 0 {
			newState := []int {y-1,x, currPrice+1}
			queue = append(queue, newState)
			visited[y-1][x]++
		}

		// Sort the queue in regards to the price - the last element of each state
		sortQueue(queue, 2)
	}

	fmt.Println("number of steps:", k)
	return -1
}

func solve1(data [][2]int, Nx int, Ny int, nSteps int, printout bool) int {

	if printout {
		fmt.Println(data)
	}

	maze := createMaze(data, Nx, Ny, nSteps)
	if printout {
		Printer.PrintGridRune(maze, 1)
	}

	startInd := [2]int {0,0}
	endInd := [2]int {Ny-1,Nx-1}

	price := findShortestPath(maze, startInd, endInd)

	return price
}

//----------------------------------------

func solve2(data [][2]int, Nx int, Ny int, printout bool) string {

	if printout {
		fmt.Println(data)
	}

	maxSteps := len(data)
	minSteps := 0
	res := "_"
	for range 9999 {

		nSteps := (maxSteps + minSteps) / 2

		maze := createMaze(data, Nx, Ny, nSteps)
		if printout {
			Printer.PrintGridRune(maze, 1)
		}

		startInd := [2]int {0,0}
		endInd := [2]int {Ny-1,Nx-1}

		price := findShortestPath(maze, startInd, endInd)
		if price < 0 {
			// this maze is not solvable
			maxSteps = nSteps
		} else {
			minSteps = nSteps
		}

		// exit condition
		if maxSteps - minSteps == 1 {
			ind := minSteps
			return strconv.Itoa(data[ind][0])+","+strconv.Itoa(data[ind][1])
		}
	}

	return res
}

func main() {

	// data gathering and parsing
	dataTest := initData(testData)

	fileName := "day_18_data.txt"
	fileData := rw.ReadFile(fileName)
	data := initData(fileData)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_1_test := solve1(dataTest, 7, 7, 12, true)
	fmt.Println("Test solution 1 =", sol1_1_test, "->", checkSolution(sol1_1_test, testSolution1))

	t1 := time.Now()
	sol1 := solve1(data, 71,71, 1024, false)
	dur := time.Since(t1)
	fmt.Println("Solution part 1 =", sol1, "(ET =", dur, ")")

	// ---------------------------------------------
	fmt.Println()
	fmt.Println("=== Part 2 ===")
	sol2_1_test := solve2(dataTest, 7,7, false)
	fmt.Println("Test solution 2 =", sol2_1_test, "->", checkSolutionStr(sol2_1_test, testSolution2))

	t1 = time.Now()
	sol2 := solve2(data, 71,71, false)
	dur = time.Since(t1)
	fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
