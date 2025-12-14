package main

import (
	// Math "aoc_2025/tools/Math"
	Array "aoc_2025/tools/Array"
	// Printer "aoc_2025/tools/Printer"
	rw "aoc_2025/tools/rw"
	"time"

	"fmt"
	"strconv"

	"slices"
	"sort"
	"strings"
	"math"
	// "bufio"
	// "os"
)

// var must be used for global variables
var testData = []string {
	"162,817,812",
	"57,618,57",
	"906,360,560",
	"592,479,940",
	"352,342,300",
	"466,668,158",
	"542,29,236",
	"431,825,988",
	"739,650,466",
	"52,470,668",
	"216,146,977",
	"819,987,18",
	"117,168,530",
	"805,96,715",
	"346,949,466",
	"970,615,88",
	"941,993,340",
	"862,61,35",
	"984,92,344",
	"425,690,689",
}

var testSolution1, testSolution2 = 40, 25272

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

func dist(p1 []int, p2 []int) float64 {

	return math.Sqrt( math.Pow(float64(p1[0] - p2[0]), 2) + 
					math.Pow(float64(p1[1] - p2[1]), 2) + 
					math.Pow(float64(p1[2] - p2[2]), 2) )
}

func getAllSortedPairs(nodes [][]int, nodesInd []int) ([][]int) {

	pairs := [][]int {}
	distances := [][]float64 {}
	
	k := float64(0)
	for i := range nodesInd {
		for j := i+1; j < len(nodesInd); j++ {
			// fmt.Println(i,j)
			_p1 := nodes[nodesInd[i]]
			_p2 := nodes[nodesInd[j]]

			distances = append(distances, []float64 {k, dist(_p1, _p2)})
			pairs = append(pairs, []int {nodesInd[i], nodesInd[j]})
			k = k + 1
		}
	}

	// sort by distance
	sortedD := Array.CopyFloat642D(distances)
	sortedP := Array.CopyInt2D(pairs)

	sort.Slice(sortedD, func(i int, j int) bool {
		return sortedD[i][1] < sortedD[j][1]
	})

	for i := range sortedD {
		ind := int(sortedD[i][0])
		sortedP[i] = pairs[ind]
	}

	return sortedP
}

func isSamePoint(p1 int, p2 int, nodes [][]int) bool {

	if nodes[p1][0] == nodes[p2][0] && nodes[p1][1] == nodes[p2][1] && nodes[p1][2] == nodes[p2][2] {
		return true
	} else {
		return false
	}
}

func areAlreadyConnected(p1 int, p2 int, circuits [][]int, nodes [][]int) bool {

	for i := range circuits {
		
		found := false
		for j := range circuits[i] {
			if isSamePoint(circuits[i][j], p1, nodes) || isSamePoint(circuits[i][j], p2, nodes) {
				if !found {
					found = true
				} else {
					// fmt.Println("  found", p1,p2, circuits)
					return true
				}
			}
		}
	}

	return false
}

func solve1(nodes [][]int, nConnections int, printout bool) int {

	if printout {
		fmt.Println("nodes", nodes)
		fmt.Println("n connections", nConnections)
	}

	// work with indexes of nodes not nodes themselves
	nodesInd := []int {}
	nodesLeftInd := []int {}	// track which nodes are left to be connected - for debugging
	for i := range nodes {
		nodesInd = append(nodesInd, i)
		nodesLeftInd = append(nodesLeftInd, i)
	}

	// find distance between all pairs and order them once
	pairs := getAllSortedPairs(nodes, nodesInd)

	circuits := [][]int {}
	k := 0
	for k < nConnections && k < len(pairs) {
		
		p1 := pairs[k][0]
		p2 := pairs[k][1]

		// remove p1 and p2 from nodes left to connect
		_i := Array.GetIndexInt(nodesLeftInd, p1)
		if _i != -1 {
			nodesLeftInd = slices.Delete(nodesLeftInd, _i, _i+1)
		}
		_i = Array.GetIndexInt(nodesLeftInd, p2)
		if _i != -1 {
			nodesLeftInd = slices.Delete(nodesLeftInd, _i, _i+1)
		}

		// check if both p1 and p2 are already in the same circuit -> skip this pair
		if areAlreadyConnected(p1, p2, circuits, nodes) {
			if printout {
				fmt.Println(k, "skip |", p1, p2,  nodes[p1], nodes[p2])
				fmt.Println()
			}
			k = k + 1
			continue
		}

		// find current two nodes in existing circuits
		foundCircuits :=  [][]int {}
		for i := range circuits {
			currCircuit := slices.Clone(circuits[i])

			if slices.Contains(currCircuit, p1) {
				foundCircuits = append(foundCircuits, []int {i, p2})
			} else if slices.Contains(currCircuit, p2) {
				foundCircuits = append(foundCircuits, []int {i, p1})
			}

			if len(foundCircuits) >= 2 {
				break
			}
		}

		// adding current nodes to existing circuits
		if len(foundCircuits) == 0 {
			// add current pair as a new circuit
			circuits = append(circuits, []int {p1,p2})
		}else if len(foundCircuits) == 1 {
			// add one node to existing one circuit
			circuits[foundCircuits[0][0]] = append(circuits[foundCircuits[0][0]], foundCircuits[0][1])
		} else {
			// merge two circuits together as current nodes are in two circuits
			
			// refactor to simple 1d array/slice
			foundCircuitsInd := []int {}
			for i := range foundCircuits {
				foundCircuitsInd = append(foundCircuitsInd, foundCircuits[i][0])
			}

			// separate circuits to merge and the rest
			_circuits := [][]int {}
			_mergeCircuits := [][] int{}
			for i,v := range circuits {
				if slices.Contains(foundCircuitsInd, i) {
					_mergeCircuits = append(_mergeCircuits, v)
				} else {
					_circuits = append(_circuits, v)
				}
			}

			// merge to separate array
			mergedCircuit := slices.Clone(_mergeCircuits[0])
			for _,v := range _mergeCircuits[1:] {
				mergedCircuit = append(mergedCircuit, v...)
			}

			circuits = append(_circuits, mergedCircuit)
		}

		if printout {
			fmt.Println(k, "|", p1,p2, "-", dist(nodes[p1], nodes[p2]), nodes[p1], nodes[p2])
			fmt.Println("  circuits", circuits)
			fmt.Println("      left", nodesLeftInd, len(nodesLeftInd))
		}

		k = k + 1
	}
	
	sorted := Array.SortBySizeInt2D(circuits)

	fmt.Println("\nfinal multiplication")
	mult := 1
	for i := range sorted[:3] {
		fmt.Println(len(sorted[i]))
		mult = mult * len(sorted[i])
	}
	return mult
}

//----------------------------------------

func solve2(nodes [][]int, printout bool) int {

	if printout {
		fmt.Println("nodes", nodes)
	}

	// work with indexes of nodes not nodes themselves
	nodesInd := []int {}
	nodesLeftInd := []int {}
	for i := range nodes {
		nodesInd = append(nodesInd, i)
		nodesLeftInd = append(nodesLeftInd, i)
	}

	// find distance between all pairs and order them once
	pairs := getAllSortedPairs(nodes, nodesInd)

	circuits := [][]int {}
	maxConnections := 9999
	k := 0
	for k < maxConnections && k < len(pairs) {
		
		p1 := pairs[k][0]
		p2 := pairs[k][1]

		// remove p1 and p2 from nodes left to connect
		if len(nodesLeftInd) > 0 {
			_i := Array.GetIndexInt(nodesLeftInd, p1)
			if _i != -1 {
				nodesLeftInd = slices.Delete(nodesLeftInd, _i, _i+1)
			}
			_i = Array.GetIndexInt(nodesLeftInd, p2)
			if _i != -1 {
				nodesLeftInd = slices.Delete(nodesLeftInd, _i, _i+1)
			}
		}

		// check if both p1 and p2 are already in the same circuit -> skip this pair
		if areAlreadyConnected(p1, p2, circuits, nodes) {
			if printout {
				fmt.Println(k, "skip |", p1, p2,  nodes[p1], nodes[p2])
				fmt.Println()
			}
			k = k + 1
			continue
		}

		// find current two nodes in existing circuits
		foundCircuits :=  [][]int {}
		for i := range circuits {
			currCircuit := slices.Clone(circuits[i])

			if slices.Contains(currCircuit, p1) {
				foundCircuits = append(foundCircuits, []int {i, p2})
			} else if slices.Contains(currCircuit, p2) {
				foundCircuits = append(foundCircuits, []int {i, p1})
			}

			if len(foundCircuits) >= 2 {
				break
			}
		}

		// adding current nodes to existing circuits
		if len(foundCircuits) == 0 {
			// add current pair as a new circuit
			circuits = append(circuits, []int {p1,p2})
		}else if len(foundCircuits) == 1 {
			// add one node to existing one circuit
			circuits[foundCircuits[0][0]] = append(circuits[foundCircuits[0][0]], foundCircuits[0][1])
		} else {
			// merge two circuits together as current nodes are in two circuits
			
			// refactor to simple 1d array/slice
			foundCircuitsInd := []int {}
			for i := range foundCircuits {
				foundCircuitsInd = append(foundCircuitsInd, foundCircuits[i][0])
			}

			// separate circuits to merge and the rest
			_circuits := [][]int {}
			_mergeCircuits := [][] int{}
			for i,v := range circuits {
				if slices.Contains(foundCircuitsInd, i) {
					_mergeCircuits = append(_mergeCircuits, v)
				} else {
					_circuits = append(_circuits, v)
				}
			}

			// merge to separate array
			mergedCircuit := slices.Clone(_mergeCircuits[0])
			for _,v := range _mergeCircuits[1:] {
				mergedCircuit = append(mergedCircuit, v...)
			}

			circuits = append(_circuits, mergedCircuit)
		}

		if printout {
			fmt.Println(k, "|", p1,p2, "-", dist(nodes[p1], nodes[p2]), nodes[p1], nodes[p2])
			fmt.Println("  circuits", circuits)
			fmt.Println("      left", nodesLeftInd, len(nodesLeftInd))
		}

		if len(nodesLeftInd) == 0 && len(circuits) == 1 {
			return nodes[p1][0] * nodes[p2][0]
		}

		k = k + 1
	}

	return -1
}

func main() {

	// data gathering and parsing
	nodesTest1 := initData(testData)

	fileName := "day_08_data.txt"
	fileData := rw.ReadFile(fileName)
	nodes1 := initData(fileData)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_1_test := solve1(nodesTest1, 10, true)
	fmt.Println("Test solution 1 =", sol1_1_test, "->", checkSolution(sol1_1_test, testSolution1))

	t1 := time.Now()
	sol1 := solve1(nodes1, 1000, false)
	dur := time.Since(t1)
	fmt.Println("Solution part 1 =", sol1, "(ET =", dur, ")")

	// ---------------------------------------------

	fmt.Println()
	fmt.Println("=== Part 2 ===")
	sol2_1_test := solve2(nodesTest1, true)
	fmt.Println("Test solution 2 =", sol2_1_test, "->", checkSolution(sol2_1_test, testSolution2))

	t1 = time.Now()
	sol2 := solve2(nodes1, false)
	dur = time.Since(t1)
	fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
