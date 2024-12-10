package main

import (
	Array "aoc_2024/tools/Array"
	"sort"
	// Math "aoc_2024/tools/Math"
	rw "aoc_2024/tools/rw"
	// "time"
	Printer "aoc_2024/tools/Printer"
	"fmt"
	"strconv"
	// "strings"
)

// var must be used for global variables
var testData1 = []string {
	"89010123",
	"78121874",
	"87430965",
	"96549874",
	"45678903",
	"32019012",
	"01329801",
	"10456732",

}

var testSolution1, testSolution2 = 36, 81

//------------------------------------------------------

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed   :)"
	} else {
		return "failed   :(   (should be '"+strconv.Itoa(solValue)+"')"
	}
}

//------------------------------------------------------

func initData(fileLines []string) ([][]int, [][2]int) {

	data := [][]int{}
	starts := [][2]int{}

	for i := range fileLines {
		line := fileLines[i]

		lineData := []int{}
		for j := range line {
			_int, err := strconv.Atoi(string(line[j]))
			if err == nil {
				lineData = append(lineData, _int)
				if _int == 0 {
					starts = append(starts, [2]int{i,j})
				}
			} else {
				panic(err)
			}
		}
		data = append(data, lineData)
	}

	return data, starts
}

func getLargestTrailID(trailMap map[int]([][2]int)) int {

	maxID := 0
	for key,_ := range trailMap {
		if key > maxID {maxID = key}
	}
	return maxID
}

func CopyTrail(data [][2]int) [][2]int {

	duplicate := make([][2]int, len(data))
	for i := range data {
		duplicate[i] = [2]int {data[i][0],data[i][1]}
	}
	return duplicate
}

func doOneTrail(data [][]int, start [2]int, printout bool) [][]int {

	M, N := len(data), len(data[0])

	paths := make(map[int]([][2]int))
	paths[0] = [][2]int {{start[0], start[1]}}
	nextID := 1

	visited := Array.InitArrayValuesInt(M,N, 0)
	visited[start[0]][start[1]] = 1
	peaks := Array.InitArrayValuesInt(M,N, 0)

	maxSteps := 9999999
	for range maxSteps {

		// Copy paths so that we don't overwrite them
		copyPaths := make(map[int]([][2]int))
		for k,v := range paths {
			copyPaths[k] = v
		}

		// order keys (trailID to get the same order)
		keys := make([]int, 0)
		for k := range copyPaths {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		// Go trough all paths
		for trailID, trail := range copyPaths {
		// for _,k := range keys {
		// 	trailID, trail := k, copyPaths[k]

			lastInd := len(trail) - 1
			y, x := trail[lastInd][0], trail[lastInd][1]
			z := data[y][x]
			if printout {
				fmt.Println("curr pos:", y, x, z, "trail:", trailID)
			}

			// if reached the end - mark it and delete path
			if z == 9 {
				peaks[y][x]++

				if printout {
					// fmt.Println("Reached the end at", [2]int{y, x}, "| id =", trailID)
					fmt.Println("   the end", trailID, len(trail), trail)
					if len(trail) != 10 {
						fmt.Println("oops")
					}
				}
				delete(paths, trailID)
				continue
			}

			// check all neighbors
			foundNext := false
			if x + 1 < N && z + 1 == data[y][x+1] {
				// append to existing path if we were not yet here
				if true || visited[y][x+1] == 0 || data[y][x+1] == 9 {
					newPath := CopyTrail(trail)
					paths[trailID] = append(newPath, [2]int {y, x+1})
					visited[y][x+1] = 1
					foundNext = true
					// fmt.Println("  add new point (x+1):", [2]int {y, x+1}, "to id", trailID)
				} else {
					// fmt.Println("we were here before (x+1):", [2]int{y, x+1}, visited[y][x+1])
					delete(paths, trailID)
				}
			}

			if x - 1 >= 0 && z + 1 == data[y][x-1] {
				if true || visited[y][x-1] == 0 || data[y][x-1] == 9 {
					if !foundNext {
						newPath := CopyTrail(trail)
						paths[trailID] = append(newPath, [2]int {y, x-1})
						visited[y][x-1] = 1
						foundNext = true
						// fmt.Println("  add new point (x-1):", [2]int {y, x-1}, "to id", trailID)
					} else {
						// open a new path id
						newPath := CopyTrail(trail)
						paths[nextID] = append(newPath, [2]int {y, x-1})
						if printout {
							fmt.Println("  open new id (x-1)", nextID, [2]int {y, x-1})
						}
						nextID++
					}
				} else {
					// fmt.Println("  we were here before (x-1):", [2]int{y, x-1}, visited[y][x-1])
					delete(paths, trailID)
				}
			}

			if y + 1 < M && z + 1 == data[y+1][x] {
				if true || visited[y+1][x] == 0 || data[y+1][x] == 9 {
					if !foundNext {
						newPath := CopyTrail(trail)
						paths[trailID] = append(newPath, [2]int {y+1, x})
						visited[y+1][x] = 1
						foundNext = true
						// fmt.Println("  add new point (y+1):", [2]int {y+1, x}, "to id", trailID)
					} else {
						// open a new path id
						newPath := CopyTrail(trail)
						paths[nextID] = append(newPath, [2]int {y+1, x})
						if printout {
							fmt.Println("  open new id (y+1)", nextID, [2]int {y+1, x})
						}
						nextID++
					}
				} else {
					// fmt.Println("we were here before (y+1):", [2]int{y+1, x}, visited[y+1][x])
					delete(paths, trailID)
				}
			}

			if y - 1 >= 0 && z + 1 == data[y-1][x] {
				if true || visited[y-1][x] == 0 || data[y-1][x] == 9 {
					if !foundNext {
						newPath := CopyTrail(trail)
						paths[trailID] = append(newPath, [2]int {y-1, x})
						visited[y-1][x] = 1
						foundNext = true
						// fmt.Println("  add new point (y-1):", [2]int {y-1, x}, "to id", trailID)
					} else {
						// open a new path id
						newPath := CopyTrail(trail)
						paths[nextID] = append(newPath, [2]int {y-1, x})
						if printout {
							fmt.Println("  open new id (y-1)", nextID, [2]int {y-1, x})
						}
						nextID++
					}
				} else {
					// fmt.Println("we were here before (y-1):", [2]int{y-1, x}, visited[y-1][x])
					delete(paths, trailID)
				}
			}

			if !foundNext {
				// we are int a dead end - delete this path
				// fmt.Println("  none found")
				delete(paths, trailID)
			}

			// if printout {
			// 	fmt.Println("  ", trailID, trail)
			// }
		}

		// if i > 10 {break}
		if len(paths) == 0 {
			break
		}

		if printout {
			keys := make([]int, 0)
			for k := range paths {
				keys = append(keys, k)
			}
			sort.Ints(keys)
			fmt.Println("paths")
			for _,k := range keys {
				fmt.Println("  ", k, paths[k])
			}
			fmt.Println()
		}
	}

	// fmt.Println("\n--------")
	// for key, value := range peaks {
	// 	fmt.Println(key, value, "|", data[key[0]][key[1]])
	// }
	// fmt.Println("Yay", len(peaks))
	return peaks
}

func solve1(data [][]int, starts [][2]int, printout bool) int {

	if printout {
		Printer.PrintGridInt(data)
	}

	// i := 3
	// peaks := doOneTrail(data, starts[i], true)
	// fmt.Println("start from:", i, starts[i], "=>", len(peaks))

	sum := 0
	for k := range starts {
		// break

		peaks := doOneTrail(data, starts[k], false)

		// Count the number of peaks reached
		n := 0
		for i := range peaks {
			for j := range peaks[i] {
				if peaks[i][j] != 0 {
					n++
				}
			}
		}
		sum = sum + n

		if printout || false {
			fmt.Println("start from:", starts[k], "=>", n)
			// Printer.PrintGridInt(peaks)
		}

	}
	
	return sum
}

//----------------------------------------

func solve2(data []int, printout bool) int {

	sum := 0
	
	return sum
}

func main() {

	// data gathering and parsing
	dataTest, startsTest := initData(testData1)

	fileName := "day_10_data.txt"
	fileData := rw.ReadFile(fileName)
	data, starts := initData(fileData)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_test := solve1(dataTest, startsTest, true)
	fmt.Println("Test solution 1 =", sol1_test, "->", checkSolution(sol1_test, testSolution1))

	sol1 := solve1(data, starts, false)
	fmt.Println("Solution part 1 =", sol1)

	// ---------------------------------------------
	// fmt.Println()
	// fmt.Println("=== Part 2 ===")
	// sol2_test := solve2(dataTest, true)
	// fmt.Println("Test solution 2 =", sol2_test, "->", checkSolution(sol2_test, testSolution2))

	// t1 := time.Now()
	// sol2 := solve2(data, false)
	// dur := time.Since(t1)
	// fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
