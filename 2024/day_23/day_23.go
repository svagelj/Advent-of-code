package main

import (
	// Math "aoc_2024/tools/Math"
	// Array "aoc_2024/tools/Array"
	// Printer "aoc_2024/tools/Printer"
	rw "aoc_2024/tools/rw"
	"time"

	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
	// "bufio"
	// "os"
)

// var must be used for global variables
var testData = []string {
	"kh-tc",
	"qp-kh",
	"de-cg",
	"ka-co",
	"yn-aq",
	"qp-ub",
	"cg-tb",
	"vc-aq",
	"tb-ka",
	"wh-tc",
	"yn-cg",
	"kh-ub",
	"ta-co",
	"de-co",
	"tc-td",
	"tb-wq",
	"wh-td",
	"ta-ka",
	"td-qp",
	"aq-cg",
	"wq-ub",
	"ub-vc",
	"de-ta",
	"wq-aq",
	"wq-vc",
	"wh-yn",
	"ka-de",
	"kh-ta",
	"co-tc",
	"wh-qp",
	"tb-vc",
	"td-yn",
}


var testSolution1, testSolution2 = 7, "co,de,ka,ta"

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

func initData(fileLines []string) ([][]string) {

	data := [][]string {}

	for i := range fileLines {
		line := fileLines[i]

		_line := strings.Split(line, "-")
		data = append(data, _line)
		
	}

	return data
}

func createConnectionsMap(data [][]string) map[string]([]string) {

	connections := make(map[string]([]string))

	for k := range data {
		a, b := data[k][0], data[k][1]

		_, found := connections[a]
		if !found {
			connections[a] = []string {b}
		} else {
			connections[a] = append(connections[a], b)
		}

		_, found = connections[b]
		if !found {
			connections[b] = []string {a}
		} else {
			connections[b] = append(connections[b], a)
		}
	}

	return connections
}

func solve1(data [][]string, printout bool) int {

	if printout {
		fmt.Println("data:", data)
	}

	connections := createConnectionsMap(data)
	if printout {
		for key, value := range connections {
			fmt.Println(key, "->", value)
		}
	}

	cycles := make(map[[3]string]bool)
	for key, value := range connections {
		for key2 := range value {
			value2, found := connections[value[key2]]
			if found {
				for key3 := range value2 {
					values3, found2 := connections[value2[key3]]
					if found2 && slices.Contains(values3, key) {
						cycle := [3]string{key, value[key2], value2[key3]}
						sort.Strings(cycle[:])
						_, f := cycles[cycle]
						if !f && (cycle[0][:1] == "t" || cycle[1][:1] == "t" || cycle[2][:1] == "t") {
							cycles[cycle] = true
						}
						
					}
				}
			}
		}
	}

	if printout {
		for key := range cycles {
			fmt.Println(key)
		}
	}

	sum := len(cycles)
	return sum
}

//----------------------------------------

func solve2(data [][]string, printout bool) string {

	if printout {
		fmt.Println("data:", data)
	}

	// making connection for each node
	connections := createConnectionsMap(data)
	keys := []string {}
	for key := range connections {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	if printout {
		for k := range keys {
			conn := connections[keys[k]]
			sort.Strings(conn)
			fmt.Println(keys[k], "->", conn)
		}
		fmt.Println()
	}

	nLargest := 0
	bestList := []string {}

	for i := range keys {
		key := keys[i]
		value := connections[key]

		cycle := []string {key}
		// for each first node check it's connections if they are connected to each other
		for j := range value {
			if slices.Contains(connections[key], value[j]) {
				connectedToAll := true
				for k := range cycle {
					if !slices.Contains(connections[value[j]], cycle[k]) {
						connectedToAll = false
						break
					}
				}

				if connectedToAll {
					cycle = append(cycle, value[j])
				}
			}
		}

		sort.Strings(cycle)
		if len(cycle) > nLargest {
			bestList = cycle
			nLargest = len(cycle)
		}

		if printout {
			fmt.Println(key, "->", cycle)
		}
	}

	out := strings.Join(bestList, ",")
	return out
}

func main() {

	// data gathering and parsing
	dataTest := initData(testData)

	fileName := "day_23_data.txt"
	fileData := rw.ReadFile(fileName)
	data := initData(fileData)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_1_test := solve1(dataTest, true)
	fmt.Println("Test solution 1 =", sol1_1_test, "->", checkSolution(sol1_1_test, testSolution1))

	t1 := time.Now()
	sol1 := solve1(data, false)
	dur := time.Since(t1)
	fmt.Println("Solution part 1 =", sol1, "(ET =", dur, ")")

	// ---------------------------------------------
	fmt.Println()
	fmt.Println("=== Part 2 ===")
	sol2_1_test := solve2(dataTest, true)
	fmt.Println("Test solution 2 =", sol2_1_test, "->", checkSolutionStr(sol2_1_test, testSolution2))

	t1 = time.Now()
	sol2 := solve2(data, false)
	dur = time.Since(t1)
	fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
