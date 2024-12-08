package main

import (
	Array "aoc_2024/tools/Array"
	Math "aoc_2024/tools/Math"
	rw "aoc_2024/tools/rw"
	"time"
	Printer "aoc_2024/tools/Printer"
	"fmt"
	// "strconv"
	// "strings"
)

// var must be used for global variables
var testData1 = []string {
	"............",
	"........0...",
	".....0......",
	".......0....",
	"....0.......",
	"......A.....",
	"............",
	"............",
	"........A...",
	".........A..",
	"............",
	"............",
}

var testSolution1, testSolution2 = 14, 34

//------------------------------------------------------

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed   :)"
	} else {
		return "failed   :("
	}
}

//------------------------------------------------------

func initData(fileLines []string, emptyChar rune) ([][] rune, map[rune][][2]int) {

	data := [][]rune {}
	positions := make(map[rune][][2]int)

	for i := range fileLines {
		line := []rune(fileLines[i])
		data = append(data, line)

		// save positions
		for j := range line {
			if line[j] != emptyChar {
				value, contains := positions[line[j]]
				if contains {
					positions[line[j]] = append(value, [2]int{i,j})
				} else {
					positions[line[j]] = [][2]int{{i,j}}
				}
			}
		}
	}

	return data, positions
}

func addAntinodes1(antinodes map[[2]int]int, positions map[rune][][2]int, antennaType rune, M int,N int) map[[2]int]int {

	// Go trough each pair of antennas and add symmetrically the antinodes
	for i := range positions[antennaType] {

		for j := i+1; j < len(positions[antennaType]); j++ {
			pos1 := positions[antennaType][i]
			pos2 := positions[antennaType][j]
			dx := pos1[1]-pos2[1]
			dy := pos1[0]-pos2[0]

			antinode1 := [2]int {pos1[0]+dy, pos1[1]+dx}
			antinode2 := [2]int {pos2[0]-dy, pos2[1]-dx}

			if antinode1[0] >= 0 && antinode1[0] < M && antinode1[1] >= 0 && antinode1[1] < N {
				_, contains := antinodes[antinode1]
				if contains {
					antinodes[antinode1]++
				} else {
					antinodes[antinode1] = 1
				}
			}
			if antinode2[0] >= 0 && antinode2[0] < M && antinode2[1] >= 0 && antinode2[1] < N {
				_, contains := antinodes[antinode2]
				if contains {
					antinodes[antinode2]++
				} else {
					antinodes[antinode2] = 1
				}
			}
		}
	}

	return antinodes
}

func solve1(data [][]rune, positions map[rune][][2]int, printout bool) int {

	if printout {
		Printer.PrintGridRune(data)
		for key, value := range positions {
			fmt.Printf("%c: %d\n", key, value)
		}
	}

	M := len(data)
	N := len(data[0])
	antinodes := make(map[[2]int]int)
	disAntinodes := Array.InitArrayValuesInt(M,N, 0)

	// go through all antenna types
	sum := 0
	for key, value := range positions {

		if printout{
			fmt.Printf("curr antenna type: %c - %d\n", key, value)
		}

		addAntinodes1(antinodes, positions, key, M,N)
	}

	if printout {
		for key, value := range antinodes {
			i := key[0]
			j := key[1]
			disAntinodes[i][j] = value
		}
		Printer.PrintGridInt(disAntinodes)
	}

	sum = len(antinodes)
	return sum
}

//----------------------------------------

func addAntinodes2(antinodes map[[2]int]int, positions map[rune][][2]int, antennaType rune, M int,N int, repeat int) map[[2]int]int {

	// Go trough each pair of antennas and add symmetrically the antinodes
	for i := range positions[antennaType] {

		for j := i+1; j < len(positions[antennaType]); j++ {
			pos1 := positions[antennaType][i]
			pos2 := positions[antennaType][j]
			dx := pos1[1]-pos2[1]
			dy := pos1[0]-pos2[0]

			for k := range repeat {
				antinode1 := [2]int {pos1[0]+k*dy, pos1[1]+k*dx}
				antinode2 := [2]int {pos2[0]-k*dy, pos2[1]-k*dx}

				if antinode1[0] >= 0 && antinode1[0] < M && antinode1[1] >= 0 && antinode1[1] < N {
					_, contains := antinodes[antinode1]
					if contains {
						antinodes[antinode1]++
					} else {
						antinodes[antinode1] = 1
					}
				}
				if antinode2[0] >= 0 && antinode2[0] < M && antinode2[1] >= 0 && antinode2[1] < N {
					_, contains := antinodes[antinode2]
					if contains {
						antinodes[antinode2]++
					} else {
						antinodes[antinode2] = 1
					}
				}
			}
		}
	}

	return antinodes
}

func solve2(data [][]rune, positions map[rune][][2]int, printout bool) int {

	if printout {
		Printer.PrintGridRune(data)
		for key, value := range positions {
			fmt.Printf("%c: %d\n", key, value)
		}
	}

	M := len(data)
	N := len(data[0])
	antinodes := make(map[[2]int]int)
	disAntinodes := Array.InitArrayValuesInt(M,N, 0)

	// go through all antenna types
	sum := 0
	for key, value := range positions {

		if printout{
			fmt.Printf("curr antenna type: %c - %d\n", key, value)
		}

		addAntinodes2(antinodes, positions, key, M,N, Math.MaxInt(M,N))
	}

	if printout {
		for key, value := range antinodes {
			i := key[0]
			j := key[1]
			disAntinodes[i][j] = value
		}
		Printer.PrintGridInt(disAntinodes)
	}

	sum = len(antinodes)
	return sum
}

func main() {

	// data gathering and parsing
	emptyChar := '.'
	dataTest, positionsTest := initData(testData1, emptyChar)

	fileName := "day_08_data.txt"
	fileData := rw.ReadFile(fileName)
	data, positions := initData(fileData, emptyChar)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_test := solve1(dataTest, positionsTest, true)
	fmt.Println("Test solution 1 =", sol1_test, "->", checkSolution(sol1_test, testSolution1))

	sol1 := solve1(data, positions, false)
	fmt.Println("Solution part 1 =", sol1)

	// ---------------------------------------------
	fmt.Println()
	fmt.Println("=== Part 2 ===")
	sol2_test := solve2(dataTest, positionsTest, true)
	fmt.Println("Test solution 2 =", sol2_test, "->", checkSolution(sol2_test, testSolution2))

	t1 := time.Now()
	sol2 := solve2(data, positions, false)
	dur := time.Since(t1)
	fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
