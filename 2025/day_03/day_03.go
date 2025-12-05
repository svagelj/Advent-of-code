package main

import (
	// Math "aoc_2025/tools/Math"
	// Array "aoc_2025/tools/Array"
	// Printer "aoc_2025/tools/Printer"
	rw "aoc_2025/tools/rw"
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
	"987654321111111",
	"811111111111119",
	"234234234234278",
	"818181911112111",
}

var testSolution1, testSolution2 = 357, 3121910778619

//------------------------------------------------------

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed   :)"
	} else {
		return "failed   :(   (should be '"+strconv.Itoa(solValue)+"')"
	}
}

//------------------------------------------------------

func initData(fileLines []string) ([]string) {

	batteryBanks := []string {}

	for i := range fileLines {
		line := fileLines[i]

		batteryBanks = append(batteryBanks, line)		
	}

	return batteryBanks
}

func getBiggestValue(batteryBank string) (int, int) {

	maxValue := -1
	maxInt := -1
	for i := range batteryBank {
		valueInt := int(batteryBank[i]) - '0'

		if valueInt > maxValue {
			maxValue = valueInt
			maxInt = i
		}
	}

	return maxInt, int(batteryBank[maxInt]) - '0'
}

func solve1(batteryBanks []string, printout bool) int {

	if printout {
		fmt.Println("battery bank:", batteryBanks)
	}

	sum := 0
	for i := range batteryBanks {

		firstNumberInd := -1
		firstNumberValue := -1

		currentBatteryBank := batteryBanks[i]
		n := len(currentBatteryBank)
		firstNumberInd, firstNumberValue = getBiggestValue(currentBatteryBank[:n-1])

		_, secondNumberValue := getBiggestValue(batteryBanks[i][firstNumberInd+1:])

		bestCombinationValue := firstNumberValue*10 + secondNumberValue
		if printout {
			fmt.Println(currentBatteryBank, "->", bestCombinationValue)
		}

		sum = sum + bestCombinationValue
	}

	return sum
}

//----------------------------------------

func solve2(batteryBanks []string, printout bool) int {

	if printout {
		fmt.Println("battery bank:", batteryBanks)
	}

	nDigits := 12

	sum := 0
	for i := range batteryBanks {

		chosenDigits := ""

		currentInd := 0
		for j := range nDigits {
			searchString := batteryBanks[i][currentInd:]
			n := len(searchString)
			searchString = searchString[:n-(nDigits-j-1)]
			
			foundInd, _ := getBiggestValue(searchString)
			chosenDigits = chosenDigits + string(searchString[foundInd])

			currentInd = currentInd + foundInd + 1
		}

		chosenDigitsInt, err := strconv.Atoi(chosenDigits)
		if err != nil {
			panic(err)
		}

		if printout {
			fmt.Println(batteryBanks[i], "->", chosenDigits, chosenDigitsInt)
		}

		sum = sum + chosenDigitsInt
	}

	return sum
}

func main() {

	// data gathering and parsing
	batteryBanksTest := initData(testData)

	fileName := "day_03_data.txt"
	fileData := rw.ReadFile(fileName)
	batteryBanks1 := initData(fileData)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_1_test := solve1(batteryBanksTest, true)
	fmt.Println("Test solution 1 =", sol1_1_test, "->", checkSolution(sol1_1_test, testSolution1))

	t1 := time.Now()
	sol1 := solve1(batteryBanks1, false)
	dur := time.Since(t1)
	fmt.Println("Solution part 1 =", sol1, "(ET =", dur, ")")

	// ---------------------------------------------
	fmt.Println()
	fmt.Println("=== Part 2 ===")
	sol2_1_test := solve2(batteryBanksTest, true)
	fmt.Println("Test solution 2 =", sol2_1_test, "->", checkSolution(sol2_1_test, testSolution2))

	t1 = time.Now()
	sol2 := solve2(batteryBanks1, false)
	dur = time.Since(t1)
	fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
