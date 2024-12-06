package main

import (
	Array "aoc_2024/tools/Array"
	Math "aoc_2024/tools/Math"
	rw "aoc_2024/tools/rw"
	"time"

	// Printer "aoc_2024/tools/Printer"
	"fmt"
	"strings"
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

var testSolution1, testSolution2 = 41, 6
var rotate90CwMatrix = [][]int {{0,1}, {-1,0}}

//------------------------------------------------------

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed"
	} else {
		return "failed"
	}
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
			ind := Array.GetIndexString(fileLines[i], startingChar)
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

func stepping1(data [][]rune, startingInd []int, maxSteps int, printout bool) [][]int {

	visited := Array.InitArrayValuesInt(len(data), len(data[0]), 0)
	visited[startingInd[0]][startingInd[1]] = 1

	// Printer.PrintGridInt(visited)
	// fmt.Println("-----")
	// Printer.PrintGridRune(data)
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

		// Printer.PrintGridInt(visited)
		// fmt.Println("----------")
	}

	fmt.Println("Max steps reached at:", currentPos, currDirection)
	return visited
}

func solve1(data [][]rune, startingInd []int, maxSteps int, printout bool) int {

	sum := 0

	visited := stepping1(data, startingInd, maxSteps, printout)

	for i := range visited {
		for j := range visited[i] {
			sum = sum + visited[i][j]
		}
	}

	return sum
}

//----------------------------------------

func isDirectionInHistory(history [][]int, direction []int) bool {

	for i := range history {
		for j := range history[i] {
			if history[i][j] != direction[j] {
				return false
			}
		}
	}

	return true
}

func stepping2(data [][]rune, startingInd []int, maxSteps int, printout bool) bool {

	visitedDir := Array.InitArrayValuesInt4D(len(data), len(data[0]))
	visitedDir[startingInd[0]][startingInd[1]] = [][]int {{-1,0}}
	visitedPos := Array.InitArrayValuesInt(len(data), len(data[0]), 0)
	visitedPos[startingInd[0]][startingInd[1]] = 1

	// fmt.Println(visitedDir)
	// fmt.Println("-----")
	// Printer.PrintGridRune(data)
	// fmt.Println()

	M := len(data)
	N := len(data[0])
	currentPos := []int {startingInd[0], startingInd[1]}
	currDirection := []int {-1, 0}
	obstacle := '#'
	for k := range maxSteps {

		nextI := currentPos[0] + currDirection[0]
		nextJ := currentPos[1] + currDirection[1]
		
		// Exit condition - we went outside the grid
		if nextI < 0 || nextI >= M || nextJ < 0 || nextJ >= N {
			if printout {
				fmt.Println("Reach freedom at", k, "steps")
			}
			return false
		}

		potNextValue := data[nextI][nextJ]
		if potNextValue != obstacle {
			// save this as current position and move on
			currentPos[0] = currentPos[0] + currDirection[0]
			currentPos[1] = currentPos[1] + currDirection[1]

			// fmt.Println("b) curr history dir at", currentPos, "=", visitedDir[currentPos[0]][currentPos[1]])
			// update visited data
			visitedPos[currentPos[0]][currentPos[1]] = 1
			if len(visitedDir[currentPos[0]][currentPos[1]][0]) == 0  {
				// set first direction in this position

				// Copy element wise to NOT link data and rewrite it later!!
				visitedDir[currentPos[0]][currentPos[1]][0] = []int {currDirection[0], currDirection[1]}
				// fmt.Println("new history dir at", currentPos, visitedDir[currentPos[0]][currentPos[1]])

			} else {
				// Exit condition - we reached the same point: same position and direction
				// fmt.Println("  check if", currDirection, "in", visitedDir[currentPos[0]][currentPos[1]], "at", currentPos)

				// potential edge case: if there is the smallest cycle without strait lines, this would not find it 
				// because new direction at collision position is not checked
				if isDirectionInHistory(visitedDir[currentPos[0]][currentPos[1]], currDirection){
					return true
				} else {
					// append this direction to visitedDir at this position - if it is a new position

					// Copy element wise to NOT link data and rewrite it later!!
					visitedDir[currentPos[0]][currentPos[1]] = append(visitedDir[currentPos[0]][currentPos[1]], []int {currDirection[0], currDirection[1]})
					// fmt.Println("new append dir at", currentPos, visitedDir[currentPos[0]][currentPos[1]])
				}

			}
		} else {
			x := Math.MatrixDotVector(rotate90CwMatrix, currDirection)
			currDirection[0] = int(x[0])
			currDirection[1] = int(x[1])

			// possible solution for one edge case
			// fmt.Println("  check if", currDirection, "in", visitedDir[currentPos[0]][currentPos[1]])
			// fmt.Println("  at", currentPos, k)
			// if isDirectionInHistory(visitedDir[currentPos[0]][currentPos[1]], currDirection){
			// 	return true
			// }
		}
	}

	fmt.Println("Max steps", maxSteps,"reached at:", currentPos, currDirection)
	return false
}

func solve2(data [][]rune, startingInd []int, maxSteps int, printout bool) int {

	// Stepping trough original grid to get visited places. Those are an option to put an obstacle
	visitedOriginal := stepping1(data, startingInd, maxSteps, false)
	// Printer.PrintGridInt(visitedOriginal)
	// fmt.Println()
	// Printer.PrintGridRune(data)
	// fmt.Println()

	sum := 0
	for i := range visitedOriginal {
		for j := range visitedOriginal[i] {

			if visitedOriginal[i][j] == 1 && !(i == startingInd[0] && j == startingInd[1]) {
				// add an obstacle to this position and check if it's a cycle now

				// It's hard to deep copy data structures
				// newData := append([][] rune {}, data...)
				newData := Array.CopyRune2D(data)
				newData[i][j] = '#'

				isCycle := stepping2(newData, startingInd, maxSteps, false)
				if isCycle {
					if printout {
						fmt.Println(i,j, "is cycle")
					}
					sum++
				}
			}
		}
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
	fmt.Println()
	fmt.Println("=== Part 2 ===")
	sol2_test := solve2(dataTest, startingIndTest, 9999, true)
	fmt.Println("Test solution 2 =", sol2_test, "->", checkSolution(sol2_test, testSolution2))

	t1 := time.Now()
	sol2 := solve2(data, startingInd, 99999999, false)
	dur := time.Since(t1)
	fmt.Println("Solution part 2 =", sol2, "(Et =", dur, ")")
}
