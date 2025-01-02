package main

import (
	// Math "aoc_2024/tools/Math"
	// Array "aoc_2024/tools/Array"
	// Printer "aoc_2024/tools/Printer"
	rw "aoc_2024/tools/rw"
	"time"

	"fmt"
	"strconv"
	// "slices"
	// "sort"
	// "strings"
	// "bufio"
	// "os"
)

// var must be used for global variables
var testData = []string {
	"#####",
	".####",
	".####",
	".####",
	".#.#.",
	".#...",
	".....",
	"",
	"#####",
	"##.##",
	".#.##",
	"...##",
	"...#.",
	"...#.",
	".....",
	"",
	".....",
	"#....",
	"#....",
	"#...#",
	"#.#.#",
	"#.###",
	"#####",
	"",
	".....",
	".....",
	"#.#..",
	"###..",
	"###.#",
	"###.#",
	"#####",
	"",
	".....",
	".....",
	".....",
	"#....",
	"#.#..",
	"#.#.#",
	"#####",
}


var testSolution1, testSolution2 = 3, -1

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

func initData(fileLines []string) ([][]string, [][]string) {

	keys := [][]string {}
	locks := [][]string {}

	_data := []string {}
	for i := range fileLines {
		line := fileLines[i]
		
		if len(line) == 0 {
			// append to keys or locks
			if _data[0] == "....." {
				keys = append(keys, _data)
			} else if _data[0] == "#####" {
				locks = append(locks, _data)
			}
			_data = []string {}
		} else {
			_data = append(_data, line)
		}
	}

	// append the last one
	if _data[0] == "....." {
		keys = append(keys, _data)
	} else if _data[0] == "#####" {
		locks = append(locks, _data)
	}

	return keys, locks
}

func convertToHeightData(keys [][]string, locks [][]string) ([][]int, [][]int) {

	keysH := [][]int {}
	locksH := [][]int {}

	for k := range keys {
		keysH = append(keysH, []int {})
		key := keys[k]
		for i := 0; i < len(key[0]); i++ {
			n := 0
			for j := len(key)-2; j >= 0; j-- {
				if key[j][i] == '.' {
					break
				}
				n++
			}
			keysH[k] = append(keysH[k], n)
		}
	}

	for k := range locks {
		locksH = append(locksH, []int {})
		lock := locks[k]
		for i := 0; i < len(lock[0]); i++ {
			n := 0
			for j := 1; j < len(lock); j++ {
				if lock[j][i] == '.' {
					break
				}
				n++
			}
			locksH[k] = append(locksH[k], n)
		}
	}

	return keysH, locksH
}

func solve1(keys [][]string, locks [][]string, printout bool) int {

	keysH, locksH :=  convertToHeightData(keys, locks)

	if printout {
		fmt.Println("keys:")
		for k := range keys {
			fmt.Println("    ", keysH[k])
		}

		fmt.Println("locks:")
		for k := range locksH {
			fmt.Println("    ", locksH[k])
		}
	}

	sum := 0

	maxHeight := 5
	for i := range keysH {
		for j := range locksH {

			goodCombo := true
			for k := range keysH[i] {
				if keysH[i][k] + locksH[j][k] > maxHeight {
					goodCombo = false 
					break
				}
			}
			if goodCombo {
				sum++
			}
		}
	}

	return sum
}

//----------------------------------------

func solve2(data [][]string, printout bool) int {

	if printout {
		fmt.Println("data:", data)
	}

	sum := 0
	return sum
}

func main() {

	// data gathering and parsing
	keysTest, locksTest := initData(testData)

	fileName := "day_25_data.txt"
	fileData := rw.ReadFile(fileName)
	keys, locks := initData(fileData)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_1_test := solve1(keysTest, locksTest, true)
	fmt.Println("Test solution 1 =", sol1_1_test, "->", checkSolution(sol1_1_test, testSolution1))

	t1 := time.Now()
	sol1 := solve1(keys, locks, false)
	dur := time.Since(t1)
	fmt.Println("Solution part 1 =", sol1, "(ET =", dur, ")")

	// ---------------------------------------------
	// fmt.Println()
	// fmt.Println("=== Part 2 ===")
	// sol2_1_test := solve2(dataTest, true)
	// fmt.Println("Test solution 2 =", sol2_1_test, "->", checkSolutionStr(sol2_1_test, testSolution2))

	// t1 = time.Now()
	// sol2 := solve2(data, false)
	// dur = time.Since(t1)
	// fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
