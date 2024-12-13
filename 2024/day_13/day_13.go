package main

import (
	Array "aoc_2024/tools/Array"
	"strings"
	// "sort"
	// Math "aoc_2024/tools/Math"
	rw "aoc_2024/tools/rw"
	"time"
	// Printer "aoc_2024/tools/Printer"
	"fmt"
	// "slices"
	"strconv"
)

// var must be used for global variables
var testData1 = []string {
	"Button A: X+94, Y+34",
	"Button B: X+22, Y+67",
	"Prize: X=8400, Y=5400",
	"",
	"Button A: X+26, Y+66",
	"Button B: X+67, Y+21",
	"Prize: X=12748, Y=12176",
	"",
	"Button A: X+17, Y+86",
	"Button B: X+84, Y+37",
	"Prize: X=7870, Y=6450",
	"",
	"Button A: X+69, Y+23",
	"Button B: X+27, Y+71",
	"Prize: X=18641, Y=10279",
}

var testSolution1, testSolution2 = 480, -1

//------------------------------------------------------

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed   :)"
	} else {
		return "failed   :(   (should be '"+strconv.Itoa(solValue)+"')"
	}
}

//------------------------------------------------------

func initData(fileLines []string, addPrizeValue int) [][][2]int {

	data := [][][2]int {}

	machineData := [][2]int {}
	for i := range fileLines {
		line := fileLines[i]

		if len(line) == 0 {
			data = append(data, machineData)
			machineData = [][2]int {}
		} else {
			xStart := Array.GetIndexString(line, "X")+1
			yStart := Array.GetIndexString(line, "Y")+1
			if !strings.Contains(line, "Button") {
				xStart++
				yStart++
			} 
			xEnd := Array.GetIndexString(line, ",")

			xValue, err := strconv.Atoi(line[xStart:xEnd])
			if err != nil {
				panic(err)
			}
			yValue, err := strconv.Atoi(line[yStart:])
			if err != nil {
				panic(err)
			}
			
			if !strings.Contains(line, "Button") {
				machineData = append(machineData, [2]int{xValue+addPrizeValue, yValue+addPrizeValue})
			} else {
				machineData = append(machineData, [2]int{xValue, yValue})
			}
		}
	}

	data = append(data, machineData)

	return data
}


func solveOneGameByHand(gameData [][2]int, costA int, costB int, maxButtonPress int) float64 {

	Ax, Ay := gameData[0][0], gameData[0][1]
	Bx, By := gameData[1][0], gameData[1][1]
	Px, Py := gameData[2][0], gameData[2][1]

	// solving equations -> solve matrix
	// Ax * Na + Bx * Nb = Px
	// Ay * Na + By * Nb = Py
	// 
	// matrix form: M*x = P, x=[Na,Nb] => x = M^{-1}*P

	// determinant = Ax*By - Bx*Ay > 0
	d := Ax*By - Bx*Ay

	if d == 0 {
		return 0
	} else {
		// inverse of M without determinant
		ax, bx := By, -Bx
		ay, by := -Ay, Ax

		Na := float64(ax*Px + bx*Py) / float64(d)
		Nb := float64(ay*Px + by*Py) / float64(d)

		if float64(int(Na)) - Na != 0 || float64(int(Nb)) - Nb != 0 {
			// fmt.Println("Number of presses is not int")
			return 0
		}

		if Na < 0 || Nb < 0 {
			// fmt.Println("Button press is negative")
			return 0
		}

		if maxButtonPress > 0 && (Na > float64(maxButtonPress) || Nb > float64(maxButtonPress)) {
			// fmt.Println("Button press over", maxButtonPress, "times")
			return 0
		}

		return float64(costA)*Na + float64(costB)*Nb
	}
}

func solve1(data [][][2]int, printout bool) float64 {

	costA, costB := 3,1
	maxButtonPress := 100

	sum := 0.0
	for k := range data {
		tokens := solveOneGameByHand(data[k], costA, costB, maxButtonPress)

		if printout {
			fmt.Println(k+1, "=>", tokens)
		}
		sum = sum + tokens
	}
	
	return sum
}

//----------------------------------------

func solve2(data [][][2]int, printout bool) float64 {

	costA, costB := 3,1
	maxButtonPress := -1

	sum := 0.0
	for k := range data {
		tokens := solveOneGameByHand(data[k], costA, costB, maxButtonPress)

		if printout {
			fmt.Println(k+1, "=>", tokens, "| integer:", int(tokens))
		}
		sum = sum + tokens
	}
	
	return sum
}

func main() {

	// data gathering and parsing
	dataTest := initData(testData1, 0)
	dataTest2 := initData(testData1, 10000000000000)

	fileName := "day_13_data.txt"
	fileData := rw.ReadFile(fileName)
	data := initData(fileData, 0)
	data2 := initData(fileData, 10000000000000)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_test := solve1(dataTest, true)
	fmt.Println("Test solution 1 =", sol1_test, "->", checkSolution(int(sol1_test), testSolution1))

	t1 := time.Now()
	sol1 := solve1(data, false)
	dur := time.Since(t1)
	fmt.Println("Solution part 1 =", sol1, "(ET =", dur, ")")

	// ---------------------------------------------
	fmt.Println()
	fmt.Println("=== Part 2 ===")
	sol2_test := solve2(dataTest2, true)
	fmt.Println("Test solution 2 =", sol2_test, "->", checkSolution(int(sol2_test), testSolution2))

	t1 = time.Now()
	sol2 := solve2(data2, false)
	dur = time.Since(t1)
	fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
	fmt.Printf("integer: %d", int(sol2))
}
