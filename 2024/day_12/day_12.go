package main

import (
	Array "aoc_2024/tools/Array"
	// "sort"
	// Math "aoc_2024/tools/Math"
	rw "aoc_2024/tools/rw"
	"time"
	Printer "aoc_2024/tools/Printer"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// var must be used for global variables
var testData0 = []string {
	"AAAA",
	"BBCD",
	"BBCC",
	"EEEC",
}
// var testData0 = []string {
// 	"OOOOO",
// 	"OXOXO",
// 	"OOOOO",
// 	"OXOXO",
// 	"OOOOO",
// }
var testData1 = []string {
	"RRRRIICCFF",
	"RRRRIICCCF",
	"VVRRRCCFFF",
	"VVRCCCJFFF",
	"VVVVCJJCFE",
	"VVIVCCJJEE",
	"VVIIICJJEE",
	"MIIIIIJJEE",
	"MIIISIJEEE",
	"MMMISSJEEE",
}

var testSolution1, testSolution2 = 1930, 1206

//------------------------------------------------------

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed   :)"
	} else {
		return "failed   :(   (should be '"+strconv.Itoa(solValue)+"')"
	}
}

//------------------------------------------------------

func initData(fileLines []string) [][]rune {

	data := [][]rune{}

	for i := range fileLines {
		line := fileLines[i]
		lineData := []rune(line)
		data = append(data, lineData)
	}

	return data
}

func floodFillQ(data [][]rune, plant rune, i int, j int) [][2]int {

	nodes := [][2] int {}
	queue := [][2]int{{i,j}}
	visited := Array.InitArrayValuesInt(len(data), len(data[0]), 0)

	k := 0
	for {

		curLoc := queue[0]
		y,x := curLoc[0], curLoc[1]
		queue = queue[1:]

		nodes = append(nodes, curLoc)
		visited[y][x] = 1

		if x+1 < len(data[0]) && data[y][x+1] == plant && visited[y][x+1] == 0 && !slices.Contains(queue, [2]int {y,x+1}) {
			queue = append(queue, [2]int{y,x+1})
		}
		if x > 0 && data[y][x-1] == plant && visited[y][x-1] == 0 && !slices.Contains(queue, [2]int {y,x-1}) {
			queue = append(queue, [2]int{y,x-1})
		}

		if y+1 < len(data) && data[y+1][x] == plant && visited[y+1][x] == 0 && !slices.Contains(queue, [2]int {y+1,x}) {
			queue = append(queue, [2]int{y+1,x})
		}
		if y > 0 && data[y-1][x] == plant && visited[y-1][x] == 0 && !slices.Contains(queue, [2]int {y-1,x}) {
			queue = append(queue, [2]int{y-1,x})
		}

		if len(queue) == 0 {
			break
		}
		k++
	}

	return nodes
}

func printRegion(nodes [][2]int, M int, N int) {

	area := Array.InitArrayValuesRune(M,N, '.')
	for k := range nodes {
		i,j := nodes[k][0], nodes[k][1]
		area[i][j] = '#'
	}
	Printer.PrintGridRune(area)

}

func getNewRegionName(regions map[string][][2]int, plant rune) string {

	plantStr := string(plant)

	n := 0
	for key := range regions {
		if strings.Contains(key, plantStr) {
			n++
		}
	}

	if n == 0 {
		return plantStr + "-1"
	} else {
		return plantStr + "-"+strconv.Itoa(n+1)
	}
}

func makeRegions(data [][]rune) map[string][][2]int {

	regions := make(map[string][][2]int)
	visited := Array.InitArrayValuesInt(len(data), len(data[0]), 0)

	for i := range data {
		for j := range data[i] {
			if visited[i][j] == 0 {
				// we are in new region -> flood fill
				plant := data[i][j]
				region := floodFillQ(data, plant, i,j)

				for k := range len(region) {
					n,m := region[k][0], region[k][1]
					visited[n][m] = 1
				}

				newRegionName := getNewRegionName(regions, plant)
				regions[newRegionName] = region
			}
		}
	}

	return regions
}

func getPlantPerimeter(data [][]rune, i int, j int, M int, N int) int {

	perimeter := 0
	plant := data[i][j]

	if j == N-1 || data[i][j+1] != plant {
		perimeter++
	}
	if j == 0 || data[i][j-1] != plant {
		perimeter++
	}

	if i == M-1 || data[i+1][j] != plant {
		perimeter++
	}
	if i == 0 || data[i-1][j] != plant {
		perimeter++
	}

	return perimeter
}

func solve1(data [][]rune, printout bool) int {

	if printout {
		Printer.PrintGridRune(data)
	}

	regions := makeRegions(data)

	// if printout {
	// 	for regionName, region := range regions {
	// 		fmt.Println("--------")
	// 		fmt.Println(regionName)
	// 		printRegion(region, len(data), len(data[0]))
	// 	}
	// }

	M,N := len(data), len(data[0])
	area := make(map[string]int)
	perimeter := make(map[string]int)

	i:=0
	for rName, region := range regions {
		// fmt.Println("  ", i, "/", len(regions))
		for k := range len(region) {
			i,j := region[k][0], region[k][1]
			
			// area
			_, found := area[rName]
			if !found {
				area[rName] = 1
			} else {
				area[rName]++
			}

			// perimeter
			// take into account different regions
			perim := getPlantPerimeter(data, i,j, M,N)
			_, found = perimeter[rName]
			if !found {
				perimeter[rName] = perim
			} else {
				perimeter[rName] = perimeter[rName] + perim
			}
		}
		i++
	}

	if printout {
		for key := range area {
			fmt.Printf("key %s: a=%d, p=%d => %d \n", key, area[key], perimeter[key], area[key]*perimeter[key])
			// fmt.Println("   ", regions[key], len(regions[key]))
		}
	}

	sum := 0
	for key := range area {
		sum = sum + area[key]*perimeter[key]
	}
	
	return sum
}

//----------------------------------------

func getRegionSides(data [][]rune, region [][2]int) int {

	plant := data[region[0][0]][region[0][1]]
	M,N := len(data), len(data[0])
	visitedAbove := [][2]int {}
	visitedBellow := [][2]int {}
	visitedRight := [][2]int {}
	visitedLeft := [][2]int {}
	sides := 0

	for k :=  range region {
		y,x := region[k][0], region[k][1]

		// above
		if !slices.Contains(visitedAbove, [2]int{y,x}) && (y == 0 || data[y-1][x] != plant) {
			sides++
			leftStop, rightStop := false, false
			// go right and left for all points applicable
			for i := range N {
				// look right
				if x+i == N {
					rightStop = true
				}
				if !rightStop && data[y][x+i] == plant && (y == 0 || data[y-1][x+i] != data[y][x+i]) {
					if !slices.Contains(visitedAbove, [2]int{y,x+i}) {
						visitedAbove = append(visitedAbove, [2]int {y,x+i})
					}
				} else {
					rightStop = true
				}

				// look left
				if x-i < 0 {
					leftStop = true
				}
				if !leftStop && data[y][x-i] == plant && (y == 0 || data[y-1][x-i] != data[y][x-i]) {
					if !slices.Contains(visitedAbove, [2]int{y,x-i}) {
						visitedAbove = append(visitedAbove, [2]int {y,x-i})
					}
				} else {
					leftStop = true
				}

				if rightStop && leftStop {
					break
				}
			}
		}

		// bellow
		if !slices.Contains(visitedBellow, [2]int{y,x}) && (y == M-1 || data[y+1][x] != plant) {
			sides++
			leftStop, rightStop := false, false
			// go right and left for all points applicable
			for i := range M {
				// look right
				if x + i == N {
					rightStop = true
				}
				if !rightStop && data[y][x+i] == plant && (y == M-1 || data[y+1][x+i] != data[y][x+i]) {
					if !slices.Contains(visitedBellow, [2]int{y,x+i}) {
						visitedBellow = append(visitedBellow, [2]int {y,x+i})
					}
				} else {
					rightStop = true
				}

				// look left
				if x - i < 0 {
					leftStop = true
				}
				if !leftStop && data[y][x-i] == plant && (y == M-1 || data[y+1][x-i] != data[y][x-i]) {
					if !slices.Contains(visitedBellow, [2]int{y,x-i}) {
						visitedBellow = append(visitedBellow, [2]int {y,x-i})
					}
				} else {
					leftStop = true
				}

				if rightStop && leftStop {
					break
				}
			}
		}

		// right side
		if !slices.Contains(visitedRight, [2]int{y,x}) && (x == M-1 || data[y][x+1] != plant) {
			sides++
			upStop, downStop := false, false
			// go up and down for all points applicable
			for i := range N {
				// look down
				if y+i == M {
					downStop = true
				}
				if !downStop && data[y+i][x] == plant && (x == M-1 || data[y+i][x+1] != data[y+i][x]) {
					if !slices.Contains(visitedRight, [2]int{y+i,x}) {
						visitedRight = append(visitedRight, [2]int {y+i,x})
					}
				} else {
					downStop = true
				}

				// look up
				if y-i < 0 {
					upStop = true
				}
				if !upStop && data[y-i][x] == plant && (x == M-1 || data[y-i][x+1] != data[y-i][x]) {
					if !slices.Contains(visitedRight, [2]int{y-i,x}) {
						visitedRight = append(visitedRight, [2]int {y-i,x})
					}
				} else {
					upStop = true
				}

				if upStop && downStop {
					break
				}
			}
		}

		// left side
		if !slices.Contains(visitedLeft, [2]int{y,x}) && (x == 0 || data[y][x-1] != plant) {
			sides++
			upStop, downStop := false, false
			// go up and down for all points applicable
			for i := range N {
				// look down
				if y+i == M {
					downStop = true
				}
				if !downStop && data[y+i][x] == plant && (x == 0 || data[y+i][x-1] != data[y+i][x]) {
					if !slices.Contains(visitedLeft, [2]int{y+i,x}) {
						visitedLeft = append(visitedLeft, [2]int {y+i,x})
					}
				} else {
					downStop = true
				}

				// look up
				if y-i < 0 {
					upStop = true
				}
				if !upStop && data[y-i][x] == plant && (x == 0 || data[y-i][x-1] != data[y-i][x]) {
					if !slices.Contains(visitedLeft, [2]int{y-i,x}) {
						visitedLeft = append(visitedLeft, [2]int {y-i,x})
					}
				} else {
					upStop = true
				}

				if upStop && downStop {
					break
				}
			}
		}
	}

	return sides
}

func solve2(data [][]rune, printout bool) int {

	if printout {
		Printer.PrintGridRune(data)
	}

	area := make(map[string]int)
	sides := make(map[string]int)
	regions := makeRegions(data)

	for rName, region := range regions {
		// fmt.Println("  ", i, "/", len(regions))
		for range len(region) {
			// area
			_, found := area[rName]
			if !found {
				area[rName] = 1
			} else {
				area[rName]++
			}

			sides[rName] = getRegionSides(data, regions[rName])
		}
	}

	if printout {
		for key := range area {
			fmt.Printf("key %s: a=%d, p=%d => %d \n", key, area[key], sides[key], area[key]*sides[key])
			// fmt.Println("   ", regions[key], len(regions[key]))
		}
	}

	sum := 0
	for key := range area {
		sum = sum + area[key]*sides[key]
	}
	
	return sum
}

func main() {

	// data gathering and parsing
	dataTest := initData(testData1)

	fileName := "day_12_data.txt"
	fileData := rw.ReadFile(fileName)
	data := initData(fileData)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_test := solve1(dataTest, true)
	fmt.Println("Test solution 1 =", sol1_test, "->", checkSolution(sol1_test, testSolution1))

	t1 := time.Now()
	sol1 := solve1(data, false)
	dur := time.Since(t1)
	fmt.Println("Solution part 1 =", sol1, "(ET =", dur, ")")

	// ---------------------------------------------
	fmt.Println()
	fmt.Println("=== Part 2 ===")
	sol2_test := solve2(dataTest, true)
	fmt.Println("Test solution 2 =", sol2_test, "->", checkSolution(sol2_test, testSolution2))

	t1 = time.Now()
	sol2 := solve2(data, false)
	dur = time.Since(t1)
	fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
