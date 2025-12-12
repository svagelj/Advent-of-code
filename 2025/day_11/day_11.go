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

var testSolution1, testSolution2 = 5, -1

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
				fmt.Println("Duplicate nodes!!")
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
					// newRoute must be copied independently
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

func solve2(nodes [][]int, printout bool) int {

	if printout {
		fmt.Println("nodes", nodes)
	}
	
	sum := 0
	return sum
}

func main() {

	// data gathering and parsing
	nodesTest1 := initData(testData)

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

	// fmt.Println()
	// fmt.Println("=== Part 2 ===")
	// sol2_1_test := solve2(nodesTest1, true)
	// fmt.Println("Test solution 2 =", sol2_1_test, "->", checkSolution(sol2_1_test, testSolution2))

	// t1 = time.Now()
	// sol2 := solve2(diagram1, start1, false)
	// dur = time.Since(t1)
	// fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
