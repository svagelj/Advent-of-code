package main

import (
	// Math "aoc_2024/tools/Math"
	rw "aoc_2024/tools/rw"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// var must be used for global variables
var testData1 = []string {
	"47|53",
	"97|13",
	"97|61",
	"97|47",
	"75|29",
	"61|13",
	"75|53",
	"29|13",
	"97|29",
	"53|29",
	"61|53",
	"97|53",
	"61|29",
	"47|13",
	"75|47",
	"97|75",
	"47|61",
	"75|61",
	"47|29",
	"75|13",
	"53|13",
	"",
	"75,47,61,53,29",
	"97,61,53,29,13",
	"75,29,13",
	"75,97,47,61,53",
	"61,13,29",
	"97,13,75,29,47",
}

var testSolution1, testSolution2 = 143, 123

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed"
	} else {
		return "failed"
	}
}

func getIndexInt(array []int, element int) int {
	for i := range array {
		if array[i] == element {
			return i
		}
	}
	return -1
}

func initData(fileLiens []string, printout bool) ([][] int, [][] int) {

	rules := [][] int{}
	updates := [][] int{}

	isFirstPart := true
	var lineData []int
	i := 0
	for i < len(fileLiens) {
		line := fileLiens[i]

		if len(line) == 0 {
			isFirstPart = false
			i++
			continue
		}

		if isFirstPart {
			// Get array from string, separator is empty space
			// _line := strings.Fields(line)
			_line := strings.Split(line, "|")
			
			// Go trough _data an convert each element to int
			lineData = []int{}
			for j := range _line {
				_int, err := strconv.Atoi(_line[j])
				if err == nil {
					lineData = append(lineData, _int)
				} else {
					panic(err)
				}
			}

			if printout == true {
				fmt.Printf("%2d: '%s' => %v\n", i+1, line, lineData)
			}

			rules = append(rules, lineData)
		} else {
			// Get array from string, separator is empty space
			// _line := strings.Fields(line)
			_line := strings.Split(line, ",")
			
			// Go trough _data an convert each element to int
			lineData = []int{}
			for j := range _line {
				_int, err := strconv.Atoi(_line[j])
				if err == nil {
					lineData = append(lineData, _int)
				} else {
					panic(err)
				}
			}

			if printout == true {
				fmt.Printf("%2d: '%s' => %v\n", i+1, line, lineData)
			}

			updates = append(updates, lineData)
		}

		i++
	}

	if printout == true {
		fmt.Println("rules", rules)
		fmt.Println("updates", updates)
	}

	return rules, updates

}

func isUpdateCorrect(update []int, rules [][]int) bool {

	// fmt.Println("update:", update)
	for i := range rules {

		// fmt.Println("  rule:", rules[i])

		firstFound := false
		secondFound := false

		for j := range update {

			// fmt.Println("\tupdate:", update[j])

			if slices.Contains(update, rules[i][0]) && slices.Contains(update, rules[i][1]) {
				if update[j] == rules[i][0] {
					firstFound = true
					// fmt.Println("\t  found first")
				}
				if update[j] == rules[i][1] {
					secondFound = true
					// fmt.Println("\t  found second")
				}

				if secondFound == true && firstFound == false {
					return false
				} else if secondFound && firstFound {
					continue
				}
			}
		}
	}

	return true
}

func solve1(rules [][]int, updates [][]int, printout bool) int {

	sum := 0
	for i := range updates {
		// fmt.Println("yay:", data[i])

		isCorrect := isUpdateCorrect(updates[i], rules)

		if printout {
			fmt.Println(i, updates[i], "=>", isCorrect)
		}

		if isCorrect {
			middleInd := len(updates[i]) / 2	// this should be int floor

			if printout {
				fmt.Println("\t", middleInd, len(updates[i]), "->", updates[i][middleInd])
			}

			sum = sum + updates[i][middleInd]
		}
	}

	return sum
}

//----------------------------------------

func orderAnUpdate(update []int, rules [][]int, maxIter int) []int {

	orderedUpdate := append([]int {}, update...)
	// fmt.Println("update:", update)

	for k := range maxIter {
		// fmt.Println("iter", k)
		for i := range rules {

			ind1 := getIndexInt(orderedUpdate, rules[i][0])
			ind2 := getIndexInt(orderedUpdate, rules[i][1])

			if ind1 != -1 && ind2 != -1 {
				if ind1 > ind2 {
					// fmt.Println("  broken rule:", rules[i], "switch:", ind1, ind2, "->", orderedUpdate[ind1], orderedUpdate[ind2])
					
					// switch these two elements
					tmp1 := orderedUpdate[ind1]
					orderedUpdate[ind1] = orderedUpdate[ind2]
					orderedUpdate[ind2] = tmp1
					// fmt.Println("\t", orderedUpdate)
				}
			}
		}

		if isUpdateCorrect(orderedUpdate, rules) {
			// fmt.Println("  update is correctly ordered")
			break
		}

		if k == maxIter-1 {
			fmt.Println("Max iteration reached:", maxIter)
		}
	}

	return orderedUpdate
}

func solve2(rules [][]int, updates [][]int, printout bool) int {

	maxIter := 10
	sum := 0
	for i := range updates {
		// fmt.Println("yay:", data[i])

		isCorrect := isUpdateCorrect(updates[i], rules)

		if isCorrect == false {

			orderedUpdate := orderAnUpdate(updates[i], rules, maxIter)

			middleInd := len(orderedUpdate) / 2			// this should be int floor
			sum = sum + orderedUpdate[middleInd]

			if printout {
				fmt.Println(i, updates[i], "=>", orderedUpdate)
				fmt.Println("\t", middleInd, len(orderedUpdate), "->", orderedUpdate[middleInd])
			}
		}
	}

	return sum
}

func main() {

	// data gathering and parsing
	testRules, testUpdates := initData(testData1, false)

	fileName := "day_05_data.txt"
	fmt.Println("Reading file '" + fileName + "'")
	fileData := rw.ReadFile(fileName)
	rules, updates := initData(fileData, false)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_test := solve1(testRules, testUpdates, true)
	fmt.Println("Test solution 1 =", sol1_test, "->", checkSolution(sol1_test, testSolution1))

	sol1 := solve1(rules, updates, false)
	fmt.Println("Solution part 1 =", sol1)

	// ---------------------------------------------
	fmt.Println()
	fmt.Println("=== Part 2 ===")
	sol2_test := solve2(testRules, testUpdates, true)
	fmt.Println("Test solution 2 =", sol2_test, "->", checkSolution(sol2_test, testSolution2))

	sol2 := solve2(rules, updates, false)
	fmt.Println("Solution part 2 =", sol2)

	// fmt.Println("yolo", 4/2, 3/2)
}
