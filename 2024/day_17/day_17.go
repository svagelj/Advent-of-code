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
var testData1 = []string {
	"Register A: 729",
	"Register B: 0",
	"Register C: 0",
	"",
	"Program: 0,1,5,4,3,0",
}

var testData2 = []string {
	"Register A: 2024",
	"Register B: 0",
	"Register C: 0",
	"",
	"Program: 0,3,5,4,3,0",
}

var testSolution1, testSolution2 = "4,6,3,5,6,3,5,2,1,0", 117440

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

	registersCopy := append([]int{}, registers...)
	programCopy := append([]int{}, program...)

	programPointer := 0
	_res := []string {}

	for programPointer < len(programCopy) {
		if programCopy[programPointer] == 3 {
			if registersCopy[0] != 0 {
				programPointer = programCopy[programPointer+1]
			} else {
				programPointer = programPointer + 2
			}
		} else {
			out := instructions(programCopy[programPointer], programCopy[programPointer+1], registersCopy, false)
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

func instructions2(opcode int, input int, registers []int, printout bool) int {

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
		value := combos(input, registers) % 8
		return value
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

	return -1
}

func solve2(registers []int, program []int, printout bool) int {

	if printout {
		fmt.Println("registers:", registers)
		fmt.Println("program:  ", program)
	}

	registersCopy := append([]int{}, registers...)
	programCopy := append([]int{}, program...)
	delta := 1000000

	k := 0
	for k < 999999999999999999 {
		programPointer := 0
		res := []int {}

		registersCopy[0] = k
		if k % delta == 0 {
			fmt.Println("\nguess:", k)
			// fmt.Println("reg", registersCopy)
			// fmt.Println("prog", programCopy)
		}

		for programPointer < len(programCopy) {
			if programCopy[programPointer] == 3 {
				if registersCopy[0] != 0 {
					programPointer = programCopy[programPointer+1]
				} else {
					programPointer = programPointer + 2
				}
			} else {
				out := instructions2(programCopy[programPointer], programCopy[programPointer+1], registersCopy, false)
				if out != -1 {
					res = append(res, out)
				}

				programPointer = programPointer + 2
			}
		}

		if k % delta == 0 {
			fmt.Println(programCopy)
			fmt.Println(res)
		}
		// if k > 100 {
		// 	break
		// }

		if len(res) != len(programCopy) {
			k++
			continue
		}

		same := true
		for i := range res {
			if res[i] != programCopy[i] {
				same = false
				break
			}
		}
		if same {
			// fmt.Println("Output is the same as program! regA =", k, registers[0])
			return k
		}

		k++
	}

	sum := 0
	return sum
}

func main() {

	// data gathering and parsing
	registersTest1, programTest1 := initData(testData1)
	// registersTest2, programTest2 := initData(testData2)

	fileName := "day_17_data.txt"
	fileData := rw.ReadFile(fileName)
	registers, program := initData(fileData)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	testSolve()
	sol1_1_test := solve1(registersTest1, programTest1, true)
	fmt.Println("Test solution 1 =", sol1_1_test, "->", checkSolutionStr(sol1_1_test, testSolution1))

	t1 := time.Now()
	sol1 := solve1(registers, program, true)
	dur := time.Since(t1)
	fmt.Println("Solution part 1 =", sol1, "(ET =", dur, ")")

	// ---------------------------------------------
	fmt.Println()
	fmt.Println("=== Part 2 ===")
	// sol2_1_test := solve2(registersTest2, programTest2, true)
	// fmt.Println("Test solution 2 =", sol2_1_test, "->", checkSolution(sol2_1_test, testSolution2))

	t1 = time.Now()
	sol2 := solve2(registers, program, false)
	dur = time.Since(t1)
	fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
