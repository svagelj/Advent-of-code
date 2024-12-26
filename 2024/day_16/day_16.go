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

var testSolution1_1, testSolution1_2, testSolution2_1, testSolution2_2 = 7036, 11048, 45, 64

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

func getPathWithSamePoint(pathsMap map[int]([][2]int), y int,x int, dy int,dx int) int {

	for key := range pathsMap {
		// fmt.Println("key", key)
		N := len(pathsMap[key])
		p_1 := pathsMap[key][N-1]

		if p_1[0] == y && p_1[1] == x {
			p_2 := pathsMap[key][N-2]
			_dy := p_1[0] - p_2[0]
			_dx := p_1[1] - p_2[1]
			// fmt.Println("   ", _dy, _dx, "|", dy,dx)
			if _dy == dy && _dx == dx {
				return key
			}
		}
	}

	return -999
}

func mergePaths(path1 [][2]int, path2 [][2]int) [][2]int {

	path := append([][2]int {}, path1...)
	// fmt.Println("yolo", path, "#", path1, "#", path2)

	if len(path) == 0 {
		return append([][2]int {}, path2...)
	}

	for i := range path2 {
		found := false
		for j := range path1 {
			if path1[j][0] == path2[i][0] && path1[j][1] == path2[i][1] {
				found = true
				break
			}
		}

		if !found {
			path = append(path, path2[i])
		}
	}

	return path
}

func getPriceFromVisited(visitedLocation [][]int, dy int, dx int) int {

	// fmt.Println("visited", visitedLocation, dy, dx)

	if len(visitedLocation) == 1 && len(visitedLocation[0]) == 0 {
		return -1
	}

	for k := range visitedLocation {
		if visitedLocation[k][0] == dy && visitedLocation[k][1] == dx {
			return visitedLocation[k][2]
		}
	}

	return -1
}

func addToQueue2(queue [][]int, data [][]rune, stepPrice int, turnPrice int, currState []int, visited [][][][]int, pathsMap map[int]([][2]int)) ([][]int, map[int]([][2]int)) {

	y,x, dy,dx, currPrice, pathInd := currState[0], currState[1], currState[2], currState[3], currState[4], currState[5]
	currPath := append([][2]int {}, pathsMap[pathInd]...)
	added := false

	// add right
	if data[y][x+1] != '#' {
		newPrice := stepPrice
		if dy == 0 && dx == -1 {
			newPrice = newPrice + 2*turnPrice
		} else if dx == 0 && (dy == 1 || dy == -1) {
			newPrice = newPrice + turnPrice
		}

		newState := []int {y,x+1, 0,1, currPrice+newPrice, pathInd}

		if !newStateInVisited(newState, visited[y][x+1]) {
			queue = append(queue, newState)
			// fmt.Println("add right", newPrice, newState)
			if len(visited[y][x+1]) == 1 && len(visited[y][x+1][0]) == 0 {
				visited[y][x+1] = [][]int {{0,1, currPrice+newPrice}}
			} else {
				visited[y][x+1] = append(visited[y][x+1], []int {0,1, currPrice+newPrice})
			}

			pathsMap[pathInd] = append(pathsMap[pathInd], [2]int {y,x+1})

			added = true
		} else if getPriceFromVisited(visited[y][x+1], 0,1) == currPrice+newPrice {
			// fmt.Println("merge paths", y,x,"(adding right)")
			key := getPathWithSamePoint(pathsMap, y,x+1,0,1)

			if key != -1 {
				pathsMap[key] = mergePaths(pathsMap[pathInd], pathsMap[key])
			} else {
				panic("key -1")
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
		
		var newState []int
		if !added {
			newState = []int {y,x-1, 0,-1, currPrice+newPrice, pathInd}
		} else {
			newState = []int {y,x-1, 0,-1, currPrice+newPrice, len(pathsMap)+1}
		}

		if !newStateInVisited(newState, visited[y][x-1]) {
			queue = append(queue, newState)
			// fmt.Println("add left", newPrice, newState)
			if len(visited[y][x-1]) == 1 && len(visited[y][x-1][0]) == 0 {
				visited[y][x-1] = [][]int {{0,-1, currPrice+newPrice}}
			} else {
				visited[y][x-1] = append(visited[y][x-1], []int {0,-1, currPrice+newPrice})
			}

			_, foundKey := pathsMap[newState[5]]
			if !foundKey {
				// new branch of paths
				pathsMap[newState[5]] = append([][2]int{}, currPath...)
				pathsMap[newState[5]] = append(pathsMap[newState[5]], [2]int {y,x-1})
			} else {
				pathsMap[pathInd] = append(pathsMap[pathInd], [2]int {y,x-1})
			}

			added = true
		} else if getPriceFromVisited(visited[y][x-1], 0,-1) == currPrice+newPrice {
			// fmt.Println("merge paths", y,x,"(adding left)")
			key := getPathWithSamePoint(pathsMap, y,x-1,0,-1)

			if key != -1 {
				pathsMap[key] = mergePaths(pathsMap[pathInd], pathsMap[key])
			} else {
				panic("key -1")
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

		var newState []int
		if !added {
			newState = []int {y+1,x, 1,0, currPrice+newPrice, pathInd}
		} else {
			newState = []int {y+1,x, 1,0, currPrice+newPrice, len(pathsMap)+1}
		}

		if !newStateInVisited(newState, visited[y+1][x]) {
			queue = append(queue, newState)
			// fmt.Println("add down", newPrice, newState)
			if len(visited[y+1][x]) == 1 && len(visited[y+1][x][0]) == 0 {
				visited[y+1][x] = [][]int {{1,0, currPrice+newPrice}}
			} else {
				visited[y+1][x] = append(visited[y+1][x], []int {1,0, currPrice+newPrice})
			}

			_, foundKey := pathsMap[newState[5]]
			if !foundKey {
				// new branch of paths
				pathsMap[newState[5]] = append([][2]int {}, currPath...)
				pathsMap[newState[5]] = append(pathsMap[newState[5]], [2]int {y+1,x})
			} else {
				pathsMap[pathInd] = append(pathsMap[pathInd], [2]int {y+1,x})
			}

			added = true
		} else if getPriceFromVisited(visited[y+1][x], 1,0) == currPrice+newPrice {
			// fmt.Println("merge paths at", y,x,"(adding down)")
			key := getPathWithSamePoint(pathsMap, y+1,x,1,0)

			if key != 1 {
				pathsMap[key] = mergePaths(pathsMap[pathInd], pathsMap[key])
			} else {
				panic("key -1")
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

		var newState []int
		if !added {
			newState = []int {y-1,x, -1,0, currPrice+newPrice, pathInd}
		} else {
			newState = []int {y-1,x, -1,0, currPrice+newPrice, len(pathsMap)+1}
		}

		if !newStateInVisited(newState, visited[y-1][x]) {
			queue = append(queue, newState)
			// fmt.Println("add up", newPrice, newState)
			if len(visited[y-1][x]) == 1 && len(visited[y-1][x][0]) == 0 {
				visited[y-1][x] = [][]int {{-1,0, currPrice+newPrice}}
			} else {
				visited[y-1][x] = append(visited[y-1][x], []int {-1,0, currPrice+newPrice})
			}

			_, foundKey := pathsMap[newState[5]]
			if !foundKey {
				// new branch of paths
				pathsMap[newState[5]] = append([][2]int{}, currPath...)
				pathsMap[newState[5]] = append(pathsMap[newState[5]], [2]int {y-1,x})
			} else {
				pathsMap[pathInd] = append(pathsMap[pathInd], [2]int {y-1,x})
			}
		} else if getPriceFromVisited(visited[y-1][x], -1,0) == currPrice+newPrice {
			// fmt.Println("merge paths", y,x,"(adding up)")
			key := getPathWithSamePoint(pathsMap, y-1,x,-1,0)

			if key != 0 {
				pathsMap[key] = mergePaths(pathsMap[pathInd], pathsMap[key])
			} else {
				panic("key -1")
			}
		}
	}

	// Sort the queue in regards to the price - the last element of each state
	sortQueue(queue)

	return queue, pathsMap
}

func stepping2(data [][]rune, startInd [2]int, startDir [2]int, endInd [2]int) (map[int]([][2]int), []int) {

	queue := [][]int {{startInd[0], startInd[1], startDir[0], startDir[1], 0, 0}}
	popIndex := 0

	turnPrice := 1000
	stepPrice := 1

	visited := Array.InitArrayValuesInt4D(len(data), len(data[0]))
	visited[startInd[0]][startInd[1]] = [][]int {{startDir[0], startDir[1], 0}}
	pathsMap := make(map[int]([][2]int))
	pathsMap[0] = [][2]int {startInd}
	solutionPathInds := []int {}

	nSolutions := 0
	bestScore := -1
	k := 0
	for k = range 99999999 {
		if len(queue) == 0 {
			fmt.Println("empty queue")
			break
		}

		currState := queue[popIndex]
		y,x, _,_, currPrice, pathInd := currState[0], currState[1], currState[2], currState[3], currState[4], currState[5]
		queue = append(queue[:popIndex], queue[popIndex+1:]...)

		// Check for exit condition
		if y == endInd[0] && x == endInd[1] {
			// fmt.Println("\tReached finish point.", y,x, k, "|", currPrice)
			
			if bestScore == -1 {
				bestScore = currPrice
			}

			if currPrice == bestScore {
				solutionPathInds = append(solutionPathInds, pathInd)

				nSolutions++			
				continue
			} else {
				return pathsMap, solutionPathInds
			}
		} else if bestScore != -1 && currPrice > bestScore {
			// we already found the best price and current other best price is larger than the best one
			// so any new paths will be more expensive - stop search
			return pathsMap, solutionPathInds
		}

		// fmt.Println("\nstep", k)
		// fmt.Println("\nstep", k, len(queue), "|", y,x, dy,dx, currPrice)
		// fmt.Println("visited:", visited[y][x])
		// fmt.Println("current path:", currPrice, pathsMap[pathInd])
		// fmt.Println("q:", queue)

		queue, pathsMap = addToQueue2(queue, data, stepPrice, turnPrice, currState, visited, pathsMap)
	}

	fmt.Println("Max number of steps reached:", k)
	return pathsMap, solutionPathInds
}

func solve2(data [][]rune, startInd [2]int, endInd [2]int, printout bool) int {

	if printout {
		Printer.PrintGridRune(data, 2)
		fmt.Println("start", startInd)
		fmt.Println("end", endInd)
	}

	dataCopy := Array.CopyRune2D(data)
	dir := [2]int {0,1}
	currPos := [2]int {startInd[0], startInd[1]}
	endPos := [2]int {endInd[0], endInd[1]}

	allPaths, goodPathInds := stepping2(dataCopy, currPos, dir, endPos)

	// Go through found path and change data to something that is not '.'
	for k := range goodPathInds {
		for p := range allPaths[goodPathInds[k]] {
			y,x := allPaths[goodPathInds[k]][p][0], allPaths[goodPathInds[k]][p][1]
			key := '1' + k
			dataCopy[y][x] = rune(key)
		}
	}
	
	if printout {
		Printer.PrintGridRune(dataCopy, 2)
	}

	// count locations of tiles with path
	sum := 0
	for i := range dataCopy {
		for j := range dataCopy[i] {
			if dataCopy[i][j] != '.' && dataCopy[i][j] != '#' {
				sum++
			}
		}
	}

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
	sol1_1_test := solve1(dataTest1, startTest1, endTest1, true)
	fmt.Println("Test solution 1_1 =", sol1_1_test, "->", checkSolution(sol1_1_test, testSolution1_1))
	sol1_2_test := solve1(dataTest2, startTest2, endTest2, false)
	fmt.Println("Test solution 1_2 =", sol1_2_test, "->", checkSolution(sol1_2_test, testSolution1_2))

	t1 := time.Now()
	sol1 := solve1(data, start, end, false)
	dur := time.Since(t1)
	fmt.Println("Solution part 1 =", sol1, "(ET =", dur, ")")

	// ---------------------------------------------
	fmt.Println()
	fmt.Println("=== Part 2 ===")
	sol2_1_test := solve2(dataTest1, startTest1, endTest1, true)
	fmt.Println("Test solution 2_1 =", sol2_1_test, "->", checkSolution(sol2_1_test, testSolution2_1))
	sol2_2_test := solve2(dataTest2, startTest2, endTest2, false)
	fmt.Println("Test solution 2_2 =", sol2_2_test, "->", checkSolution(sol2_2_test, testSolution2_2))

	t1 = time.Now()
	sol2 := solve2(data, start, end, false)
	dur = time.Since(t1)
	fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
