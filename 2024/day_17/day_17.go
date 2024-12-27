package main

import (
	// Array "aoc_2024/tools/Array"
	// "bufio"
	// "os"
	"math"
	"strings"

	// Math "aoc_2024/tools/Math"
	rw "aoc_2024/tools/rw"
	"time"

	// Printer "aoc_2024/tools/Printer"
	"fmt"
	"strconv"
	// "sort"
)

// var must be used for global variables
var testData = []string {
	"Register A: 729",
	"Register B: 0",
	"Register C: 0",
	"",
	"Program: 0,1,5,4,3,0",
}

var testSolution1, testSolution2 = "4,6,3,5,6,3,5,2,1,0", -1

//------------------------------------------------------

func checkSolution(testValue string, solValue string) string {
	if testValue == solValue {
		return "passed   :)"
	} else {
		return "failed   :(   (should be '"+solValue+"')"
	}
}

//------------------------------------------------------

func initData(fileLines []string) ([]int, []int) {

	registers := []int {}
	program := []int {}

	firstPart := true
	for i := range fileLines {
		line := fileLines[i]

		if len(line) == 0 {
			firstPart = false
			continue
		}

		if firstPart {
			_line := strings.Split(line, ": ")
			_int, err := strconv.Atoi(_line[1])
			if err == nil {
				registers = append(registers, _int)
			} else {
				panic(err)
			}

		} else {
			_line := strings.Split(line, ": ")
			_line = strings.Split(_line[1], ",")
			for j := range _line {
				_int, err := strconv.Atoi(_line[j])
				if err == nil {
					program = append(program, _int)
				} else {
					panic(err)
				}
			}
		}
	}

	return registers, program
}

func combos(num int, registers []int) int {
	switch num {
	case 0:
		return num
	case 1:
		return num
	case 2:
		return num
	case 3:
		return num
	case 4:
		return registers[0]
	case 5:
		return registers[1]
	case 6:
		return registers[2]
	}
	 return -1
}

func bitwiseXOR(a int, b int) int {
	return a^b
}

func instructions(opcode int, input int, registers []int, printout bool) string {

	// returning float64, changed registers, program pointer

	if printout {
		fmt.Println("opcode:", opcode, "| input:", input, "| registers:", registers)
	}

	switch opcode {
	case 0:
		// division - change reg A
		base :=  math.Pow(2, float64(combos(input, registers)))
		value := float64(registers[0]) / base
		registers[0] = int(value)
	case 1:
		// bitwise XOR - change reg B
		registers[1] = bitwiseXOR(registers[1], input)
	case 2:
		// module 8 - change reg B
		registers[1] = combos(input, registers) % 8
	case 3:
		// jumping pointer - change program pointer
		// have to do it outside this function
	case 4:
		// bitwise XOR - change reg B
		registers[1] = bitwiseXOR(registers[1], registers[2])
	case 5:
		// module 8 - return value
		value := float64(combos(input, registers) % 8)
		return strconv.FormatFloat(value, 'f', -1, 64)
	case 6:
		// division - change reg B
		base :=  math.Pow(2, float64(combos(input, registers)))
		value := float64(registers[0]) / base
		registers[1] = int(value)
	case 7:
		// division - change reg C
		base :=  math.Pow(2, float64(combos(input, registers)))
		value := float64(registers[0]) / base
		registers[2] = int(value)
	}

	return ""
}

func testSolve() {
	
	// example 1
	registers := []int {0,0, 9} 
	program := []int {2,6}
	instructions(program[0], program[1], registers, false)
	if registers[1] == 1 {
		fmt.Println("Example 1 passed\t:)")
	} else {
		fmt.Println("Example 1 failed\t:(")
	}

	// example 2
	registers = []int {0, 29,0} 
	program = []int {1,7}
	instructions(program[0], program[1], registers, false)
	if registers[1] == 26 {
		fmt.Println("Example 2 passed\t:)")
	} else {
		fmt.Println("Example 2 failed\t:(")
	}

	// example 3
	registers = []int {0, 2024, 43690} 
	program = []int {4,0}
	instructions(program[0], program[1], registers, false)
	if registers[1] == 44354 {
		fmt.Println("Example 3 passed\t:)")
	} else {
		fmt.Println("Example 3 failed\t:(")
	}

	// example 4
	registers = []int {10, 0,0} 
	program = []int {5,0,5,1,5,4}
	programPointer := 0
	_res := []string {}

	for programPointer < len(program) {
		out := instructions(program[programPointer], program[programPointer+1], registers, false)
		_res = append(_res, out)
		programPointer = programPointer + 2
	}
	res := strings.Join(_res, ",")
	if res == "0,1,2" {
		fmt.Println("Example 4 passed\t:)")
	} else {
		fmt.Println("Example 4 failed\t:(")
	}

	// example 5
	registers = []int {2024, 0,0} 
	program = []int {0,1,5,4,3,0}
	programPointer = 0
	_res = []string {}

	for programPointer < len(program) {
		if program[programPointer] == 3 {
			if registers[0] != 0 {
				programPointer = program[programPointer+1]
			} else {
				programPointer = programPointer + 2
			}
		} else {
			out := instructions(program[programPointer], program[programPointer+1], registers, false)
			if out != "" {
				_res = append(_res, out)
			}

			programPointer = programPointer + 2
		}
	}
	res = strings.Join(_res, ",")
	if res == "4,2,5,6,7,7,7,7,3,1,0" && registers[0] == 0 {
		fmt.Println("Example 5 passed\t:)")
	} else {
		fmt.Println("Example 5 failed\t:(")
	}
}

func solve1(registers []int, program []int, printout bool) string {

	if printout {
		fmt.Println("registers:", registers)
		fmt.Println("program:  ", program)
	}

	programPointer := 0
	_res := []string {}

	for programPointer < len(program) {
		if program[programPointer] == 3 {
			if registers[0] != 0 {
				programPointer = program[programPointer+1]
			} else {
				programPointer = programPointer + 2
			}
		} else {
			out := instructions(program[programPointer], program[programPointer+1], registers, false)
			if out != "" {
				_res = append(_res, out)
			}

			programPointer = programPointer + 2
		}
	}

	res := strings.Join(_res, ",")
	return res
}

//----------------------------------------

func solve2(data [][]rune, startInd [2]int, endInd [2]int, printout bool) int {

	sum := 0
	return sum
}

func main() {

	// data gathering and parsing
	registersTest, programTest := initData(testData)

	fileName := "day_17_data.txt"
	fileData := rw.ReadFile(fileName)
	registers, program := initData(fileData)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	testSolve()
	sol1_1_test := solve1(registersTest, programTest, true)
	fmt.Println("Test solution 1_1 =", sol1_1_test, "->", checkSolution(sol1_1_test, testSolution1))

	t1 := time.Now()
	sol1 := solve1(registers, program, true)
	dur := time.Since(t1)
	fmt.Println("Solution part 1 =", sol1, "(ET =", dur, ")")

	// ---------------------------------------------
	// fmt.Println()
	// fmt.Println("=== Part 2 ===")
	// sol2_1_test := solve2(dataTest, startTest, endTest, true)
	// fmt.Println("Test solution 2_1 =", sol2_1_test, "->", checkSolution(sol2_1_test, testSolution2))

	// t1 = time.Now()
	// sol2 := solve2(data, start, end, false)
	// dur = time.Since(t1)
	// fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
