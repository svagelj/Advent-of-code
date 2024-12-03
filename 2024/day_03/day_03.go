package main

import (
	rw "aoc_2024/tools"
	"fmt"
	"strconv"
	"strings"
)

// var must be used for global variables
var testData1 = []string{
	"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
}

var testData2 = [] string {
	"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
}

var testSolution1, testSolution2 = 161, 48

// ----------------------------------------------------------------------

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed"
	} else {
		return "failed"
	}
}

func absInt(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func minInt(a int, b int) int {
	if a > b {
		return b
	} 

	return a
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	} 

	return b
}

func stringToInt(str string) int {

	_int, err := strconv.Atoi(str)
	if err == nil {
		return _int
	} else {
		panic(err)
	}
}

// ----------------------------------------------------------------------

// func initData(fileLiens []string, printout bool) [][] int {

// 	data := [][] int{}

// 	var n int
// 	var lineData []int
// 	i := 0
// 	for i < len(fileLiens) {
// 		line := fileLiens[i]

// 		// Get array from string, separator is empty space
// 		_data := strings.Fields(line)

// 		if printout == true {
// 			fmt.Printf("%2d: %s => %s\n", i+1, line, _data)
// 		}
		
// 		// Go trough _data an convert each element to int
// 		lineData = []int{}
// 		n = len(_data)
// 		for j := range n {
// 			_int, err := strconv.Atoi(_data[j])
// 			if err == nil {
// 				lineData = append(lineData, _int)
// 			} else {
// 				panic(err)
// 			}
// 		}
// 		data = append(data, lineData)

// 		i++
// 	}

// 	if printout == true {
// 		fmt.Println("yay", data)
// 	}

// 	return data

// }

func findCommaAndEnd(substring string) (int, int) {

	// fmt.Println("test", substring)

	// find a comma
	commaInd := -1
	endInd := -1
	for i := range substring {
		if substring[i] == ')'{
			// fmt.Println("found ')'")
			endInd = i
			break
		} else if !strings.Contains("1234567890,", substring[i:i+1]) {
			// fmt.Println("found illegal argument")
			return -1, -1
		} else if substring[i] == ','{
			commaInd = i
		}
	}

	if commaInd != -1 && endInd != -1 {
		return commaInd, endInd
	} else {
		return -1, -1
	}
}

func solve1(data []string, printout bool) int {

	var mul int
	sum := 0
	for i := range data {

		mul = 0
		for j := 0; j<len(data[i])-5; j++ {
			substringStart := data[i][j:j+4]
			containsBeginning := strings.Contains(substringStart, "mul(")

			if containsBeginning == true {
				maxInd := j+12
				if maxInd > len(data[i]) {
					maxInd = len(data[i])
				}
				maxSubstring := data[i][j:maxInd]

				if printout {
					fmt.Println("try", maxSubstring)
				}

				commaInd, endInd := findCommaAndEnd(maxSubstring[4:])
				
				if commaInd != -1 && endInd != -1 {
					firstNumber := stringToInt(maxSubstring[4:4+commaInd])
					secondNumber := stringToInt(maxSubstring[4+commaInd+1:4+endInd])
					mul = firstNumber*secondNumber

					if printout {
						fmt.Println("\tFirst number ", firstNumber)
						fmt.Println("\tSecond number", secondNumber)
						fmt.Println("\tyay", substringStart, "->", maxSubstring, "=>", mul)
					}

					sum = sum + mul
				}
			}
		}
	}

	return sum
}

// ----------------------------------------------------------------------

func tryToMultiply(dataLine string, j int, printout bool) int {

	substringStart := dataLine[j:j+4]
			containsBeginning := strings.Contains(substringStart, "mul(")

	if containsBeginning == true {
		maxInd := j+12
		if maxInd > len(dataLine) {
			maxInd = len(dataLine)
		}
		maxSubstring := dataLine[j:maxInd]

		if printout {
			fmt.Println("try", maxSubstring)
		}

		commaInd, endInd := findCommaAndEnd(maxSubstring[4:])
		
		if commaInd != -1 && endInd != -1 {
			firstNumber := stringToInt(maxSubstring[4:4+commaInd])
			secondNumber := stringToInt(maxSubstring[4+commaInd+1:4+endInd])
			mul := firstNumber*secondNumber

			if printout {
				fmt.Println("\tFirst number ", firstNumber)
				fmt.Println("\tSecond number", secondNumber)
				fmt.Println("\tyay", substringStart, "->", maxSubstring, "=>", mul)
			}

			return mul
		}
	}

	return 0
}

func solve2(data []string, printout bool) int {

	allowToMultiply := true
	var mul int
	sum := 0
	for i := range data {

		mul = 0
		for j := 0; j<len(data[i])-3; j++ {

			maxIndDo := minInt(j + 4, len(data[i]))
			maxIndDont := minInt(j + 7, len(data[i]))

			foundDo := strings.Contains(data[i][j:maxIndDo], "do()")
			foundDont := strings.Contains(data[i][j:maxIndDont], "don't()")
			
			if foundDo {
				// fmt.Println(data[i][j:maxInd1])
				if printout {
					fmt.Println("Found 'do()'")
				}
				allowToMultiply = true
			}
			if foundDont {
				// fmt.Println(data[i][j:maxInd2])
				if printout {
					fmt.Println("Found 'don't()'")
				}
				allowToMultiply = false
			}

			if allowToMultiply {
				mul = tryToMultiply(data[i], j, printout)
				if mul != 0 {
					sum = sum + mul
				}
			}
			
		}
	}

	return sum
}

// ----------------------------------------------------------------------

func main() {

	// data gathering and parsing
	// testData := initData(testData1, true)
	testData := append([]string{}, testData1...)

	fileName := "day_03_data.txt"
	lines := rw.ReadFile(fileName)
	// fileData := initData(lines, false)
	fileData := append([]string{}, lines...)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_test := solve1(testData, true)
	fmt.Println("Test solution 1 =", sol1_test, "->", checkSolution(sol1_test, testSolution1))

	sol1 := solve1(fileData, false)
	fmt.Println("Solution part 1 =", sol1)

	// // ---------------------------------------------
	fmt.Println("")
	fmt.Println("=== Part 2 ===")
	sol2_test := solve2(testData2, false)
	fmt.Println("Test solution 2 =", sol2_test, "->", checkSolution(sol2_test, testSolution2))

	sol2 := solve2(fileData, false)
	fmt.Println("Solution part 2 =", sol2)

	// a := "0123456"
	// fmt.Println(a[2:3])
}
