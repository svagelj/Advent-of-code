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
	"aaa: you hhh",
	"you: bbb ccc",
	"bbb: ddd eee",
	"ccc: ddd eee fff",
	"ddd: ggg",
	"eee: out",
	"fff: out",
	"ggg: out",
	"hhh: ccc fff iii",
	"iii: out",
}

var testData2 = []string {
	"svr: aaa bbb",
	"aaa: fft",
	"fft: ccc",
	"bbb: tty",
	"tty: ccc",
	"ccc: ddd eee",
	"ddd: hub",
	"hub: fff",
	"eee: dac",
	"dac: fff",
	"fff: ggg hhh",
	"ggg: out",
	"hhh: out",
}

var testSolution1, testSolution2 = 5, 2

//------------------------------------------------------

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed   :)"
	} else {
		return "failed   :(   (should be '"+strconv.Itoa(solValue)+"')"
	}
}

//------------------------------------------------------

func initData(fileLines []string) (map[string][](string)) {

	nodes := make(map[string]([]string))

	for i := range fileLines {
		line := fileLines[i]

		_line := strings.Split(line, " ")

		key := strings.Split(_line[0], ":")[0]
		outputs := []string {}
		for _,v := range _line[1:] {
			if slices.Contains(outputs, v){
				panic("Duplicate nodes!!")
			} else {
				outputs = append(outputs, v)
			}
		}

		nodes[key] = outputs
	}

	return nodes
}

func solve1(nodes map[string][](string), printout bool) int {

	if printout {
		fmt.Println("nodes", nodes)
	}

	endRoutes := [][]string {}
	routes := [][]string {{"you"}}

	sum := 0
	maxIter := 10
	for k := range maxIter {

		_newRoutes := [][]string {}

		for _, route := range routes {
			lastNode := route[len(route)-1]
			for _, nextNode := range nodes[lastNode] {

				if nextNode == "out" {
					sum = sum + 1
					endRoutes = append(endRoutes, append(route, nextNode))
					continue
				}

				if slices.Contains(route, nextNode) {
					// skip this route as it is a loop - there should be no loops otherwise the problem has infinite solutions
					fmt.Println("  we in a loop", route, nextNode)
					continue
				} else {
					// newRoute must be copied to be independent
					newRoute := make([]string, len(route))
					copy(newRoute, route)
					newRoute = append(newRoute, nextNode)

					_newRoutes = append(_newRoutes, newRoute)
				}
			}
		}

		routes = _newRoutes

		if k >= maxIter - 1 {
			fmt.Println("WARNING max iter", k+1, "/", maxIter, "reached!")
		}

		if len(routes) == 0 {
			break
		}
		
	}

	if printout {
		for _,v := range endRoutes {
			fmt.Println(v)
		}
	}

	return sum
}

//----------------------------------------

func findNumberOfPathsRecursive(nodes map[string][](string), startNode string, endNode string, memo map[string]([]int), depth int) int {

	// assumptions 
	// 1) every path that we are currently on must have reach the desired end node
	// 2) if 1) is not true we are in a loop of infinite length, never reaches the end
	// 3) memoization works on depth first search -> easier with recursion than queue

	// memo key is for start node because end node is fixed
	value, success := memo[startNode]
	if success {
		return value[0]
	}

	if startNode == endNode {
		return 1
	}

	sum := 0
	for _, nextNode := range nodes[startNode] {

		if nextNode == endNode {
			return 1
		}

		nPaths := findNumberOfPathsRecursive(nodes, nextNode, endNode, memo, depth+1)
		sum = sum + nPaths
	}

	// update memo with current data
	memo[startNode] = []int {sum}

	return sum
}

func solve2(nodes map[string][](string), printout bool) int {

	if printout {
		fmt.Println("nodes:", nodes)
	}

	// We have to get from 'svr' to 'out' but through two other nodes ('dac' and 'fft')
	// this means that we can find number of paths for each of these sections: 
	//	1) svr - fft (alternative: svr - dac)
	//	2) fft - dac (alternative: dac - fft)
	//	3) dac - out (alternative: fft - out)
	// The result is multiplication of all three sectors plus multiplication of alternative path
	// Number of paths for one of these two paths might always be 0

	n11 := findNumberOfPathsRecursive(nodes, "svr", "fft", make(map[string]([]int)), 1)
	n12 := findNumberOfPathsRecursive(nodes, "svr", "dac", make(map[string]([]int)), 1)
	n21 := findNumberOfPathsRecursive(nodes, "fft", "dac", make(map[string]([]int)), 1)
	n22 := findNumberOfPathsRecursive(nodes, "dac", "fft", make(map[string]([]int)), 1)
	n31 := findNumberOfPathsRecursive(nodes, "dac", "out", make(map[string]([]int)), 1)
	n32 := findNumberOfPathsRecursive(nodes, "fft", "out", make(map[string]([]int)), 1)

	if printout {
		fmt.Println("svr -> fft", n11)
		fmt.Println("svr -> fft", n12)
		fmt.Println("fft -> dac", n21)
		fmt.Println("dac -> fft", n22)
		fmt.Println("dac -> out", n31)
		fmt.Println("fft -> out", n32)

		fmt.Println("a) svr - fft - dac - out", n11*n21*n31)
		fmt.Println("b) svr - dac - fft - out", n12*n22*n32)
	}

	return n11*n21*n31 + n12*n22*n32
}

func main() {

	// data gathering and parsing
	nodesTest1 := initData(testData)
	nodesTest2 := initData(testData2)

	fileName := "day_11_data.txt"
	fileData := rw.ReadFile(fileName)
	nodes1 := initData(fileData)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_1_test := solve1(nodesTest1, true)
	fmt.Println("Test solution 1 =", sol1_1_test, "->", checkSolution(sol1_1_test, testSolution1))

	t1 := time.Now()
	sol1 := solve1(nodes1, false)
	dur := time.Since(t1)
	fmt.Println("Solution part 1 =", sol1, "(ET =", dur, ")")

	// ---------------------------------------------

	fmt.Println()
	fmt.Println("=== Part 2 ===")
	sol2_1_test := solve2(nodesTest2, true)
	fmt.Println("Test solution 2 =", sol2_1_test, "->", checkSolution(sol2_1_test, testSolution2))

	t1 = time.Now()
	sol2 := solve2(nodes1, false)
	dur = time.Since(t1)
	fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
