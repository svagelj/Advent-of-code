package main

import (
	// Array "aoc_2024/tools/Array"
	// Math "aoc_2024/tools/Math"
	// Printer "aoc_2024/tools/Printer"
	rw "aoc_2024/tools/rw"
	"time"

	"fmt"
	// "sort"
	"strconv"
	"strings"

	// "bufio"
	// "os"
)

// var must be used for global variables
var testData = []string {
	"r, wr, b, g, bwu, rb, gb, br",
	"",
	"brwrr",
	"bggr",
	"gbbr",
	"rrbgbr",
	"ubwu",
	"bwurrg",
	"brgr",
	"bbrgwb",
}

var testSolution1, testSolution2 = 6, 16

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

func initData(fileLines []string) ([]string, []string) {

	towels := []string {}
	patterns := []string {}

	for i := range fileLines {
		line := fileLines[i]

		if i == 0 {
			towels = strings.Split(line, ", ")
		} else if i > 1 {
			patterns = append(patterns, line)
		}
	}

	return towels, patterns
}

func isPatternPossible(towels []string, pattern string) bool {

	for i := range towels {

		// exit condition
		if towels[i] == pattern {
			return true
		}

		if len(pattern) > len(towels[i]) && towels[i] == pattern[:len(towels[i])] {
			// current towel matches the beginning of the patterns
			// call next iteration
			isPossible := isPatternPossible(towels, pattern[len(towels[i]):])
			if isPossible {
				return true
			}
		} 
	}

	return false
}

func solve1(towels []string, patterns []string, printout bool) int {

	if printout {
		fmt.Println(towels)
		fmt.Println(patterns)
	}

	sum := 0
	for k := 0; k < len(patterns); k++ {
		isPossible := isPatternPossible(towels, patterns[k]) 

		if printout {
			fmt.Println(patterns[k], "->", isPossible)
		}

		if isPossible {
			sum++
		}
	}

	return sum
}

//----------------------------------------

func countPossibleSolutions(towels []string, pattern string, cache map[string](int)) (int, map[string](int)) {

	// reading from cache first
	value, found := cache[pattern]
	if found {
		return value, cache
	}

	var _nPoss int
	nPossible := 0	// local counter
	for i := range towels {

		// exit condition
		// if towels[i] == pattern {return +1} - this is not ok because there are possibilities for other matches as well (br, b+r)
		if len(pattern) == 0 {
			return 1, cache
		}

		if len(pattern) >= len(towels[i]) && towels[i] == pattern[:len(towels[i])] {
			// current towel matches the beginning of the pattern
			// call next iteration with pattern without it's beginning as matched with the towel
			_nPoss, cache = countPossibleSolutions(towels, pattern[len(towels[i]):], cache)
			nPossible = nPossible + _nPoss
		} 
	}

	// save to cache after all possible starting towels were processed
	_, f := cache[pattern]
	if !f {
		cache[pattern] = nPossible
	}

	return nPossible, cache
}

func solve2(towels []string, patterns []string, printout bool) int {

	if printout {
		fmt.Println(towels)
		fmt.Println(patterns)
	}

	cache := make(map[string](int))
	var nPossible int

	sum := 0
	for k := 0; k < len(patterns); k++ {
		nPossible, cache = countPossibleSolutions(towels, patterns[k], cache) 

		if printout {
			fmt.Println(patterns[k], "->", nPossible)
		}

		sum = sum + nPossible
	}

	return sum
}

func main() {

	// data gathering and parsing
	towelsTest, patternsTest := initData(testData)

	fileName := "day_19_data.txt"
	fileData := rw.ReadFile(fileName)
	towels, patterns := initData(fileData)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_1_test := solve1(towelsTest, patternsTest, true)
	fmt.Println("Test solution 1 =", sol1_1_test, "->", checkSolution(sol1_1_test, testSolution1))

	t1 := time.Now()
	sol1 := solve1(towels, patterns, false)
	dur := time.Since(t1)
	fmt.Println("Solution part 1 =", sol1, "(ET =", dur, ")")

	// ---------------------------------------------
	fmt.Println()
	fmt.Println("=== Part 2 ===")
	sol2_1_test := solve2(towelsTest, patternsTest, true)		// 2, 1, 4, 6, 0, 1, 0
	fmt.Println("Test solution 2 =", sol2_1_test, "->", checkSolution(sol2_1_test, testSolution2))

	t1 = time.Now()
	sol2 := solve2(towels, patterns, false)
	dur = time.Since(t1)
	fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
