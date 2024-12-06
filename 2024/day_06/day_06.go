package main

import (
	Math "aoc_2024/tools/Math"
	rw "aoc_2024/tools/rw"
	"fmt"
	"strings"
	// "slices"
	// "strconv"
)

// var must be used for global variables
var testData1 = []string {
	"....#.....",
	".........#",
	"..........",
	"..#.......",
	".......#..",
	"..........",
	".#..^.....",
	"........#.",
	"#.........",
	"......#...",
}

var testSolution1, testSolution2 = 41, -1
var rotate90CwMatrix = [][]int {{0,1}, {-1,0}}

//------------------------------------------------------

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed"
	} else {
		return "failed"
	}
}

func getIndexInt(array []int, element int) int {
	for i := range array {
		if array[i] == element {
			return i
		}
	}
	return -1
}

func getIndexString(array string, element string) int {
	for i := range array {
		if array[i] == element[0] {
			return i
		}
	}
	return -1
}

func initArrayValuesInt(M int, N int, value int) [][] int {
	array := [][]int {}

	for range M {
		_line := []int {}
		for range N {
			_line = append(_line, value)
		}
		array = append(array, _line)
	}

	return array
}

func initData(fileLines []string, startingChar string) ([][] rune, []int) {

	data := [][]rune {}
	startInd := []int {-1, -1}

	i := 0
	for i < len(fileLines) {
		line := []rune(fileLines[i])
		data = append(data, line)

		if strings.Contains(fileLines[i], startingChar){
			startInd[0] = i
			ind := getIndexString(fileLines[i], startingChar)
			if ind != -1 {
				startInd[1] = ind
			} else {
				panic("Could not find x index in array!")
			}
		}
		i++
	}

	return data, startInd

}

func printGridRune(grid [][]rune){

	for i := range grid {
		for j:= range grid[i] {
			fmt.Printf("%c", grid[i][j])
		}
		fmt.Printf("\n")
	}
}

func printGridInt(grid [][]int){

	for i := range grid {
		for j:= range grid[i] {
			fmt.Printf("%d", grid[i][j])
		}
		fmt.Printf("\n")
	}
}

func stepping(data [][]rune, startingInd []int, maxSteps int, printout bool) [][]int {

	visited := initArrayValuesInt(len(data), len(data[0]), 0)
	visited[startingInd[0]][startingInd[1]] = 1

	// printGridInt(visited)
	// fmt.Println("-----")
	// printGridRune(data)
	// fmt.Println()

	M := len(data)
	N := len(data[0])
	currentPos := []int {startingInd[0], startingInd[1]}
	currDirection := []int {-1, 0}
	obstacle := '#'
	for k := range maxSteps {

		nextI := currentPos[0] + currDirection[0]
		nextJ := currentPos[1] + currDirection[1]
		
		// Exit condition
		if nextI < 0 || nextI >= M || nextJ < 0 || nextJ >= N {
			fmt.Println("Reach freedom at", k, "steps")
			return visited
		}

		potNextValue := data[nextI][nextJ]
		if potNextValue != obstacle {
			// save this as current position and move on
			currentPos[0] = currentPos[0] + currDirection[0]
			currentPos[1] = currentPos[1] + currDirection[1]
			visited[currentPos[0]][currentPos[1]] = 1
		} else {
			x := Math.MatrixDotVector(rotate90CwMatrix, currDirection)
			currDirection[0] = int(x[0])
			currDirection[1] = int(x[1])

			if printout {
				fmt.Println("new direction:", currDirection, currentPos)
			}
		}

		// printGridInt(visited)
		// fmt.Println("----------")
	}

	fmt.Println("Max steps reached, am at:", currentPos, currDirection)
	return visited
}

func solve1(data [][]rune, startingInd []int, maxSteps int, printout bool) int {

	sum := 0

	visited := stepping(data, startingInd, maxSteps, printout)

	for i := range visited {
		for j := range visited[i] {
			sum = sum + visited[i][j]
		}
	}

	return sum
}

//----------------------------------------

func solve2(data []int, updates [][]int, printout bool) int {

	sum := 0
	for i := range updates {
		fmt.Println("yay:", data[i])
	}

	return sum
}

func main() {

	// data gathering and parsing
	startString := "^"
	dataTest, startingIndTest := initData(testData1, startString)

	fileName := "day_06_data.txt"
	fileData := rw.ReadFile(fileName)
	data, startingInd := initData(fileData, startString)

	// ---------------------------------------------
	// fmt.Println("=== Part 1 ===")
	sol1_test := solve1(dataTest, startingIndTest, 100, false)
	fmt.Println("Test solution 1 =", sol1_test, "->", checkSolution(sol1_test, testSolution1))

	sol1 := solve1(data, startingInd, 9999999, false)
	fmt.Println("Solution part 1 =", sol1)

	// ---------------------------------------------
	// fmt.Println()
	// fmt.Println("=== Part 2 ===")
	// sol2_test := solve2(testRules, testUpdates, true)
	// fmt.Println("Test solution 2 =", sol2_test, "->", checkSolution(sol2_test, testSolution2))

	// sol2 := solve2(rules, updates, false)
	// fmt.Println("Solution part 2 =", sol2)

	// fmt.Println("yolo", 4/2, 3/2)
}
