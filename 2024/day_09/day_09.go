package main

import (
	// Array "aoc_2024/tools/Array"
	// Math "aoc_2024/tools/Math"
	rw "aoc_2024/tools/rw"
	"time"
	// Printer "aoc_2024/tools/Printer"
	"fmt"
	"strconv"
	// "strings"
)

// var must be used for global variables
var testData1 = []string {
	"2333133121414131402",
}

var testSolution1, testSolution2 = 1928, 2858

//------------------------------------------------------

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed   :)"
	} else {
		return "failed   :("
	}
}

//------------------------------------------------------

func initData(fileLines []string) []int {

	data := []int{}

	for i := range fileLines {
		line := fileLines[i]

		for j := range line {
			_int, err := strconv.Atoi(string(line[j]))
			if err == nil {
				data = append(data, _int)
			} else {
				panic(err)
			}
		}

	}

	return data
}

func unpackData(data []int) []int{

	unpacked := []int{}
	k := 0
	for i := range data {
		for range (data[i]) {
			if i % 2 == 0 {
				unpacked = append(unpacked, k)
			} else {
				unpacked = append(unpacked, -1)
			}
		}
		
		if i % 2 == 0 {
			k++
		}
	}

	return unpacked
}

func compactData1(unpackedData []int, printout bool) []int {

	compacted := make([]int, len(unpackedData))
	copy(compacted, unpackedData)

	j := len(compacted)-1
	for i := range compacted {

		if compacted[i] == -1 {
			// find value to put in here
			if compacted[j] != -1 {
				compacted[i] = compacted[j]
				compacted[j] = -1
			} else {
				// Current j space is empty
				for k := 0; j-k > i; k++ {
					if compacted[j-k] != -1 {
						j = j-k
						compacted[i] = compacted[j]
						compacted[j] = -1
						break
					}
				}
			}

			if printout {
				fmt.Println(compacted)
			}
		}

		// the two indexes have reached the same space (yes last step is not necessary)
		if j <= i {
			break
		}
	}

	return compacted
}

func solve1(data []int, printout bool) int {

	unpacked := unpackData(data)
	if printout {
		fmt.Println(data, "=>", unpacked)
	}

	compacted := compactData1(unpacked, printout)

	sum := 0
	for i := range compacted {

		if compacted[i] != -1 {
			sum = sum + i*compacted[i]
		}
	}
	return sum
}

//----------------------------------------

func lookForSpace(data []int, length int) int {

	startIndex := -1

	for i := range data {
		if data[i] == -1 {
			for j := i; j < len(data); j++ {
				if j-i >= length {
					return i
				}
				if data[j] != -1 {
					break
				}
			}
		}
	}

	return startIndex
}

func compactData2(unpackedData []int, printout bool) []int {

	// copy to not change original data
	compacted := make([]int, len(unpackedData))
	copy(compacted, unpackedData)

	// to prevent multiple moves of the same file
	visited := make(map[int]int)

	fLength := 0
	for i := len(compacted)-1; i >= 0; i-- {

		// current space is not empty
		if compacted[i] != -1 {
			fID := compacted[i]

			_, found := visited[fID]
			if found {
				continue
			}

			// get the length of current file
			for j := range compacted {
				// fmt.Println("  j", j, "=>", fID, i-j, compacted[j])
				if i-j >= 0 && compacted[i-j] != fID {
					fLength = j
					break
				}
			}

			// fmt.Println("yay", i, len(compacted)-i, "| file:", fID, fLength)
			emptyInd := lookForSpace(compacted, fLength)

			// move this file if space was found
			if emptyInd != -1 && emptyInd <= i{
				// fmt.Println("  found space at", emptyInd)

				// move the whole file
				for j := range fLength {
					compacted[emptyInd+j] = fID
					compacted[i-j] = -1
				}
				visited[fID] = 1
				// fmt.Println("  visited", fID)

				if printout {
					fmt.Println(compacted)
				}
			} else {
				// jump forward (backward really) the size of current file
				// fmt.Println("  no space found")
				i = i - fLength+1
			}
		}
	}

	fmt.Println("yolo", len(visited))

	return compacted
}

func solve2(data []int, printout bool) int {

	unpacked := unpackData(data)
	if printout {
		fmt.Println(data, "=>", unpacked)
	}

	compacted := compactData2(unpacked, printout)

	sum := 0
	for i := range compacted {

		if compacted[i] != -1 {
			sum = sum + i*compacted[i]
		}
	}
	return sum
}

func main() {

	// data gathering and parsing
	dataTest := initData(testData1)

	fileName := "day_09_data.txt"
	fileData := rw.ReadFile(fileName)
	data := initData(fileData)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_test := solve1(dataTest, true)
	fmt.Println("Test solution 1 =", sol1_test, "->", checkSolution(sol1_test, testSolution1))

	sol1 := solve1(data, false)
	fmt.Println("Solution part 1 =", sol1)

	// ---------------------------------------------
	fmt.Println()
	fmt.Println("=== Part 2 ===")
	sol2_test := solve2(dataTest, true)
	fmt.Println("Test solution 2 =", sol2_test, "->", checkSolution(sol2_test, testSolution2))

	t1 := time.Now()
	sol2 := solve2(data, false)
	dur := time.Since(t1)
	fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
