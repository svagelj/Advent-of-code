package main

import (
	Array "aoc_2024/tools/Array"
	"strings"
	// Math "aoc_2024/tools/Math"
	rw "aoc_2024/tools/rw"
	"time"
	// Printer "aoc_2024/tools/Printer"
	"fmt"
	"strconv"
)

// var must be used for global variables
var testData1 = []string {
	"p=0,4 v=3,-3",
	"p=6,3 v=-1,-3",
	"p=10,3 v=-1,2",
	"p=2,0 v=2,-1",
	"p=0,0 v=1,3",
	"p=3,0 v=-2,-2",
	"p=7,6 v=-1,-3",
	"p=3,0 v=-1,-2",
	"p=9,3 v=2,3",
	"p=7,3 v=-1,2",
	"p=2,4 v=2,-3",
	"p=9,5 v=-3,-3",
}

var testSolution1, testSolution2 = 12, -1

//------------------------------------------------------

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed   :)"
	} else {
		return "failed   :(   (should be '"+strconv.Itoa(solValue)+"')"
	}
}

func PrintGridInt(grid [][]int){

	for i := range grid {
		for j:= range grid[i] {
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Printf("\n")
	}
}

//------------------------------------------------------

func initData(fileLines []string) [][4]int {

	data := [][4]int {}

	for i := range fileLines {
		line := fileLines[i]

		_line := strings.Split(line, " ")
		_lineData := append(strings.Split(_line[0][2:], ","), strings.Split(_line[1][2:], ",")...)

		lineData := []int {}
		for k := range _lineData {
			valInt, err := strconv.Atoi(_lineData[k])
			if err != nil {
				panic(err)
			} 
			
			lineData = append(lineData, valInt)
		}
		data = append(data, [4]int(lineData))
	}

	return data
}

func simulateOneRobot(robotData [4]int, Nx int, Ny int, nSteps int) (int, int) {
	
	x0, y0 := robotData[0], robotData[1]
	vx, vy := robotData[2], robotData[3]

	var (
		x1 int
		y1 int
	)
	
	if vx >= 0 {
		x1 = (x0 + vx*nSteps) % Nx
	} else {
		x1 = ((x0 + vx*nSteps) % Nx + Nx)%Nx
	}

	if vy >= 0 {
		y1 = (y0 + vy*nSteps) % Ny
	} else {
		y1 = ((y0 + vy*nSteps) % Ny + Ny) % Ny
	}

	return x1,y1
}

func solve1(data [][4]int, Nx int, Ny int, nSteps int, printout bool) int {

	visited := Array.InitArrayValuesInt(Ny,Nx, 0)

	xHalf, yHalf := Nx / 2, Ny / 2
	fmt.Println(Nx, xHalf)
	fmt.Println(Ny, yHalf)

	quadrants := [4]int {}
	// quadrants numbers
	// 0 1
	// 2 3

	for k :=  range data {

		x, y := simulateOneRobot(data[k], Nx, Ny, nSteps)
		visited[y][x]++

		if printout {
			fmt.Println(k, "=>", x,y)
		}

		if x < xHalf && y < yHalf {
			quadrants[0]++
		} else if x > xHalf && y < yHalf {
			quadrants[1]++
		} else if x < xHalf && y > yHalf {
			quadrants[2]++
		}else if x > xHalf && y > yHalf {
			quadrants[3]++
		}
	}

	sum := 1
	for i := range quadrants {
		sum = sum * quadrants[i]
		if printout {
			fmt.Println("quadrant", i, "=>", quadrants[i])
		}
	}

	return sum
}

//----------------------------------------

func solve2(data [][][2]int, printout bool) int {

	sum := 0
	
	return sum
}

func main() {

	// data gathering and parsing
	dataTest := initData(testData1)
	NxTest, NyTest := 11, 7
	Nx, Ny := 101, 103
	nSteps := 100

	fileName := "day_14_data.txt"
	fileData := rw.ReadFile(fileName)
	data := initData(fileData)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_test := solve1(dataTest, NxTest,NyTest, nSteps, true)
	fmt.Println("Test solution 1 =", sol1_test, "->", checkSolution(sol1_test, testSolution1))

	t1 := time.Now()
	sol1 := solve1(data, Nx, Ny, nSteps, false)
	dur := time.Since(t1)
	fmt.Println("Solution part 1 =", sol1, "(ET =", dur, ")")

	// ---------------------------------------------
	// fmt.Println()
	// fmt.Println("=== Part 2 ===")
	// sol2_test := solve2(dataTest2, true)
	// fmt.Println("Test solution 2 =", sol2_test, "->", checkSolution(sol2_test, testSolution2))

	// t1 = time.Now()
	// sol2 := solve2(data2, false)
	// dur = time.Since(t1)
	// fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
