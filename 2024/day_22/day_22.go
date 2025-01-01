package main

import (
	// Math "aoc_2024/tools/Math"
	// Array "aoc_2024/tools/Array"
	// Printer "aoc_2024/tools/Printer"
	rw "aoc_2024/tools/rw"
	"time"

	"fmt"
	// "sort"
	"strconv"
	// "strings"

	// "bufio"
	// "os"
)

// var must be used for global variables
var testData1 = []string {
	"1",
	"10",
	"100",
	"2024",
}

var testData2 = []string {
	"1",
	"2",
	"3",
	"2024",
}

var testSolution1, testSolution2 = 37327623, 23

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

func initData(fileLines []string) ([]int) {

	data := []int {}

	for i := range fileLines {
		line := fileLines[i]

		_int, err := strconv.Atoi(line)
		if err == nil {
			data = append(data, _int)
		}
	}

	return data
}

func decodeOneNumber(number int, nSteps int) int {

	decoded := number

	for range nSteps {
		// Step 1
		decoded = (decoded * 64) ^ decoded
		decoded = decoded % 16777216

		// Step 2
		decoded = int(float64(decoded) / float64(32)) ^ decoded
		decoded = decoded % 16777216

		// Step 3
		decoded = (decoded * 2048) ^ decoded
		decoded = decoded % 16777216
	}

	return decoded
}

func solve1(data []int, nSteps int, printout bool) int {

	if printout {
		fmt.Println("data:", data)
		fmt.Println("nSteps:", nSteps)
	}

	sum := 0
	for k := range data {
		decoded := decodeOneNumber(data[k], nSteps)

		sum = sum + decoded

		if printout {
			fmt.Println(data[k], "->", decoded)
		}
	}

	return sum
}

//----------------------------------------

func getSequenceFromOneNumber(number int, nSteps int) []int {

	decoded := number
	sequence := []int {decoded%10}

	for range nSteps {
		// Step 1
		decoded = (decoded * 64) ^ decoded
		decoded = decoded % 16777216

		// Step 2
		decoded = int(float64(decoded) / float64(32)) ^ decoded
		decoded = decoded % 16777216

		// Step 3
		decoded = (decoded * 2048) ^ decoded
		decoded = decoded % 16777216

		sequence = append(sequence, decoded%10)
	}

	return sequence
}

func getDifferencesFromSequence(sequence []int) []int {

	diff := []int {}
	for k := 1; k < len(sequence); k++ {
		diff = append(diff, sequence[k]-sequence[k-1])
	}

	return diff
}

func testInstruction(sequences [][]int, differences [][]int, instruction [4]int) int {
	
	sum := 0

	for k := range sequences {
		seq := sequences[k]
		diff := differences[k]

		for i := 3; i < len(diff); i++ {
			if instruction[0] == diff[i-3] && instruction[1] == diff[i-2] && instruction[2] == diff[i-1] && instruction[3] == diff[i] {
				sum = sum + seq[i]
				break	// we found match in this sequence - move to the next one
			}
		}
	}

	return sum
}

func solve2(data []int, nSteps int, printout bool) int {

	if printout {
		fmt.Println("data:", data)
		fmt.Println("nSteps:", nSteps)
	}

	dataCopy := append([]int {}, data...)

	// Create sequences
	sequences := [][]int {}
	differences := [][]int {}
	for k := range dataCopy {
		seq := getSequenceFromOneNumber(dataCopy[k], nSteps)
		diff := getDifferencesFromSequence(seq)
		sequences = append(sequences, seq[1:])	// the first number in sequence is ignored because we can
		differences = append(differences, diff)
	}

	// The original solution:
	// This is brute force approach and it worked - kinda
	// Currently best instruction is printed while looking for others
	// The solution was stabilized quickly and so tested against AOC website
	// it's just luck that solution was at the starting few tries
	// With optimization of don't test instructions that were tested already - solution took 3 to 4 minutes
	//
	// Another optimization is to get number of occurrences for all possible instructions
	// and the check just a few top ones
	// the best solution probably has a lot of occurrences

	// Try every possible 4 number difference
	bestInstruction := [4]int {}
	bestOutcome := -1
	visited := make(map[[4]int]bool)
	for k := range dataCopy {
		for i := 3; i < len(differences[k]); i++ {
			instruction := [4]int {differences[k][i-3], differences[k][i-2], differences[k][i-1], differences[k][i]}

			// test this instruction for all data
			_, found := visited[instruction]
			if !found {
				nBananas := testInstruction(sequences, differences, instruction)
				visited[instruction] = true

				if nBananas > bestOutcome {
					bestInstruction = instruction
					bestOutcome = nBananas
					fmt.Println(instruction, "->", nBananas)
				}
			} 
		}
	}

	if printout {
		for k := range dataCopy {
			_sum := 0
			for i := 3; i < len(differences[k]); i++ {
				if bestInstruction[0] == differences[k][i-3] && bestInstruction[1] == differences[k][i-2] && 
						bestInstruction[2] == differences[k][i-1] && bestInstruction[3] == differences[k][i] {
					_sum = _sum + sequences[k][i]
					break	// we found match in this sequence - move to the next one
				}
			}
			fmt.Printf("%4d -> %d\n", dataCopy[k], _sum)
		}
	}

	if printout {
		fmt.Println("best solution:", bestOutcome, bestInstruction)
	}

	sum := bestOutcome
	return sum
}

func main() {

	// data gathering and parsing
	dataTest1 := initData(testData1)
	dataTest2 := initData(testData2)

	fileName := "day_22_data.txt"
	fileData := rw.ReadFile(fileName)
	data := initData(fileData)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_1_test := solve1(dataTest1, 2000, true)
	fmt.Println("Test solution 1 =", sol1_1_test, "->", checkSolution(sol1_1_test, testSolution1))

	t1 := time.Now()
	sol1 := solve1(data, 2000, false)
	dur := time.Since(t1)
	fmt.Println("Solution part 1 =", sol1, "(ET =", dur, ")")

	// ---------------------------------------------
	fmt.Println()
	fmt.Println("=== Part 2 ===")
	sol2_1_test := solve2(dataTest2, 2000, true)
	fmt.Println("Test solution 2 =", sol2_1_test, "->", checkSolution(sol2_1_test, testSolution2))

	t1 = time.Now()
	sol2 := solve2(data, 2000, false)
	dur = time.Since(t1)
	fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
