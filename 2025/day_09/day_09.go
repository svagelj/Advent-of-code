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
	"strings"
	"math"
	// "bufio"
	// "os"
)

// var must be used for global variables
var testData = []string {
	"7,1",
	"11,1",
	"11,7",
	"9,7",
	"9,5",
	"2,5",
	"2,3",
	"7,3",
}

var testSolution1, testSolution2 = 50, -1

//------------------------------------------------------

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed   :)"
	} else {
		return "failed   :(   (should be '"+strconv.Itoa(solValue)+"')"
	}
}

//------------------------------------------------------

func initData(fileLines []string) ([][]int) {

	nodes := [][]int {}

	for i := range fileLines {
		line := fileLines[i]

		node := []int {}
		_line := strings.Split(line, ",")
		for j := range _line {
			value, err := strconv.Atoi(_line[j])
			if err != nil {
				panic(err)
			}
			node = append(node, value)
		}
		nodes = append(nodes, node)
	}

	return nodes
}

func solve1(nodes [][]int, printout bool) int {

	if printout {
		fmt.Println("nodes", nodes)
	}

	maxS := -1

	for i := range nodes {
		for j := i+1; j < len(nodes); j++ {
			a := Math.AbsInt(nodes[i][0] - nodes[j][0]) + 1
			b := Math.AbsInt(nodes[i][1] - nodes[j][1]) + 1
			s := a*b
			if s > maxS {
				maxS = s
			}
		}
	}

	if maxS < 0 {
		panic("Did not find max area")
	}
	
	return maxS
}

//----------------------------------------

func solve2(diagram [][]rune, start []int, printout bool) int {

	if printout {
		fmt.Println("start", start)
	}

	sum := 0

	return sum
}

func main() {

	// data gathering and parsing
	nodesTest1 := initData(testData)

	fileName := "day_09_data.txt"
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

	// fmt.Println()
	// fmt.Println("=== Part 2 ===")
	// sol2_1_test := solve2(diagramTest1, startTest1, true)
	// fmt.Println("Test solution 2 =", sol2_1_test, "->", checkSolution(sol2_1_test, testSolution2))

	// t1 = time.Now()
	// sol2 := solve2(diagram1, start1, false)
	// dur = time.Since(t1)
	// fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
