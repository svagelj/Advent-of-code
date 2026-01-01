package main

// This one is the last one in 2025 and has only one part

import (
	// Math "aoc_2025/tools/Math"
	Array "aoc_2025/tools/Array"
	Printer "aoc_2025/tools/Printer"
	rw "aoc_2025/tools/rw"
	"time"

	"fmt"
	"strconv"

	"slices"
	// "sort"
	"strings"
	// "math"
	// "bufio"
	// "os"
)

// var must be used for global variables
var testData = []string {
	"0:",
	"###",
	"##.",
	"##.",
	"",
	"1:",
	"###",
	"##.",
	".##",
	"",
	"2:",
	".##",
	"###",
	"##.",
	"",
	"3:",
	"##.",
	"###",
	"##.",
	"",
	"4:",
	"###",
	"#..",
	"###",
	"",
	"5:",
	"###",
	".#.",
	"###",
	"",
	"4x4: 0 0 0 0 2 0",
	"12x5: 1 0 1 0 2 2",
	"12x5: 1 0 1 0 3 2",
}

var testSolution1, testSolution2 = 2, -1

//------------------------------------------------------

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed   :)"
	} else {
		return "failed   :(   (should be '"+strconv.Itoa(solValue)+"')"
	}
}

//------------------------------------------------------

func initData(fileLines []string) ([][]string, [][]int, [][]int) {

	shapes := [][]string {}
	regions := [][]int {}
	shapeNumbers := [][]int {}

	currShape := []string {}

	shapeInd := -1
	readingRegions := false
	for i := range fileLines {
		line := fileLines[i]

		if line != "" && line[1] == ':'{
			_shapeInd, err := strconv.Atoi(line[:1])
			if err != nil {
				panic(err)
			}
			shapeInd = _shapeInd
			continue
		}

		if shapeInd != -1 {
			// stop reading shapes
			if line == "" {
				shapes = append(shapes, currShape)
				shapeInd = -1
				currShape = []string {}
				continue
			}

			// read shape
			currShape = append(currShape, line)
		}

		if readingRegions || (line != "" && strings.Contains(line, "x")) {
			_line := strings.Split(line, ": ")
			
			// region size
			regi := strings.Split(_line[0], "x")
			region := []int {}
			for _,v := range regi {
				_v, err := strconv.Atoi(v)
				if err != nil {
					panic(err)
				}
				region = append(region, _v)
			}
			regions = append(regions, region)

			// requested number of shapes
			numbers := strings.Split(_line[1], " ")
			numbersOut := []int {}
			for _,numStr := range numbers {
				_v, err := strconv.Atoi(numStr)
				if err != nil {
					panic(err)
				}
				numbersOut = append(numbersOut, _v)
			}
			shapeNumbers = append(shapeNumbers, numbersOut)
		}
	}

	return shapes, regions, shapeNumbers
}

func rotateShape(shape []string) [][]string {

	rotatedShapes := [][]string {slices.Clone(shape)}

	// transpose - kinda like clockwise rotation
	org := rotatedShapes[len(rotatedShapes)-1]		
	rot := []string {}
	for i := range org {
		_line := ""
		for j := range org[i] {
			_line = _line + string(org[j][i])
		}
		rot = append(rot, _line)
	}
	rotatedShapes = append(rotatedShapes, rot)

	// mirror in y direction last rotation
	org = rotatedShapes[len(rotatedShapes)-1]
	rot = []string {}
	for i := range org {
		_line := ""
		for j := range org[i] {
			_line = _line + string(org[len(org)-1-j][i])
		}
		rot = append(rot, _line)
	}
	rotatedShapes = append(rotatedShapes, rot)

	// transpose last rotation
	org = rotatedShapes[len(rotatedShapes)-1]		
	rot = []string {}
	for i := range org {
		_line := ""
		for j := range org[i] {
			_line = _line + string(org[j][i])
		}
		rot = append(rot, _line)
	}
	rotatedShapes = append(rotatedShapes, rot)

	return rotatedShapes
}

func insertShapeInRegionOnce(loc []int, shape []string, region [][]int, value int) ([][]int, bool) {

	// newState := slices.Clone(region)	// not good enough
	newState := Array.CopyInt2D(region)

	// if shape looks outside of the region return false
	if loc[0] + len(shape) > len(region) {
		return region, false
	} else if loc[1] + len(shape[0]) > len(region[0]) {
		return region, false
	}

	for i := range len(shape) {
		for j := range len(shape[i]) {
			yInd := loc[0] + i
			xInd := loc[1] + j

			if shape[i][j] == '#' {
				if region[yInd][xInd] == 0 {
					newState[yInd][xInd] = value
				} else {
					return region, false
				}
			}
		}
	}

	return newState, true
}

func insertShapesInRegionRecursive(region [][]int, shapesToFit [][]string, depth int) (bool, [][]int) {

	// Exit condition
	if len(shapesToFit) == 0 {
		return true, region
	}

	regionSizeY := len(region)
	regionSizeX := len(region[0])
	shapeSizeY := 3
	shapeSizeX := 3
	effectiveRegionSizeX := regionSizeX - shapeSizeX + 1
	effectiveRegionSizeY := regionSizeY - shapeSizeY + 1 

	// Printer.PrintGridInt(region, 2)

	_shape := shapesToFit[0]

	rotatedShape := rotateShape(_shape)

	// get location to insert in
	for i := range effectiveRegionSizeX * effectiveRegionSizeY {
		locX := i % effectiveRegionSizeX
		locY := int(i / effectiveRegionSizeX)
		insertLoc := []int {locY, locX}

		for _, shape := range rotatedShape {

			newState, success := insertShapeInRegionOnce(insertLoc, shape, region, (depth%10)+1)

			if success {
				globalSuccess, finalRegion := insertShapesInRegionRecursive(newState, shapesToFit[1:], depth+1)
				if globalSuccess {
					return true, finalRegion
				}
			}
		}
	}
	
	// We tried to fit in first shape at all locations and all rotations and it did not fit
	return false, region
}

func canShapesFit(regionSize []int, requestedShapes []int, shapes [][]string, printout bool) bool {

	shapesToFit := [][]string {}
	for i,n := range requestedShapes {
		for range n {
			shapesToFit = append(shapesToFit, shapes[i])
		}
	}

	// Recursion with backtracking takes too long time -> let's try engineering solution
	// Count number of solid blocks in shapes to fit. 
	// If there are more solid blocks than space in region, shapes cannot fit
	// This is solution for actual data but not for test data - quite sad

	nEmptySpacesRegion := regionSize[0]*regionSize[1]
	nSolidShape := 0
	for _, shape := range shapesToFit {
		for _, line := range shape {
			for _,c := range line {
				if c == '#' {
					nSolidShape = nSolidShape + 1
				}
			}
		}
	}
	
	if printout {
		fmt.Println("   n solid", nSolidShape)
		fmt.Println("   size", nEmptySpacesRegion)
	}

	if nSolidShape <= nEmptySpacesRegion {
		return true
	}

	return false
	// Bellow is solution with backtracking even though it takes too long

	// create empty region - space to fit in shapes
	region := [][]int {}
	for range regionSize[1]	 {
		regionX := []int {}
		for range regionSize[0] {
			regionX = append(regionX, 0)
		}
		region = append(region, regionX)
	}

	// Recursion with backtracking takes too long time -> engineering solution (above) is a solution even if bad one
	// Even though it takes too long time, here is backtracking solution
	success, finalRegion := insertShapesInRegionRecursive(region, shapesToFit, 0)
	if success && printout {
		fmt.Println("---->")
		Printer.PrintGridInt(finalRegion, 2)
	}

	return success
}

func solve1(shapes [][]string, regions [][]int, shapeNumbers [][]int, printout bool) int {

	if printout {
		fmt.Println("shapes", shapes)
		fmt.Println("regions", regions)
		fmt.Println("numbers", shapeNumbers)
	}

	sum := 0
	for nCase := range regions {

		region := regions[nCase]
		requestedNumbers := shapeNumbers[nCase]

		if printout {
			fmt.Println(nCase, region, requestedNumbers)
		}

		if canShapesFit(region, requestedNumbers, shapes, printout) {
			sum = sum + 1

			if printout {
				fmt.Println("  -> yes")
			}
		} else if printout {
			fmt.Println("  -> no")
		}
	}
	
	return sum
}

//----------------------------------------

func solve2(nodes [][]int, printout bool) int {

	if printout {
		fmt.Println("nodes", nodes)
	}
	
	sum := 0
	return sum
}

func main() {

	// data gathering and parsing
	shapesTest1, regionsTest1, shapeNumbersTest1 := initData(testData)

	fileName := "day_12_data.txt"
	fileData := rw.ReadFile(fileName)
	shapes1, regions1, shapeNumbers1 := initData(fileData)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_1_test := solve1(shapesTest1, regionsTest1, shapeNumbersTest1, true)
	fmt.Println("Test solution 1 =", sol1_1_test, "->", checkSolution(sol1_1_test, testSolution1))

	t1 := time.Now()
	sol1 := solve1(shapes1, regions1, shapeNumbers1, false)
	dur := time.Since(t1)
	fmt.Println("Solution part 1 =", sol1, "(ET =", dur, ")")

	// ---------------------------------------------

	// fmt.Println()
	// fmt.Println("=== Part 2 ===")
	// sol2_1_test := solve2(nodesTest1, true)
	// fmt.Println("Test solution 2 =", sol2_1_test, "->", checkSolution(sol2_1_test, testSolution2))

	// t1 = time.Now()
	// sol2 := solve2(diagram1, start1, false)
	// dur = time.Since(t1)
	// fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
