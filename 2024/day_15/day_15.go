package main

import (
	Array "aoc_2024/tools/Array"
	// "bufio"
	// "os"
	"strings"

	// Math "aoc_2024/tools/Math"
	rw "aoc_2024/tools/rw"
	"time"

	Printer "aoc_2024/tools/Printer"
	"fmt"
	"strconv"
)

// var must be used for global variables
var testData1_1 = []string {
	"########",
	"#..O.O.#",
	"##@.O..#",
	"#...O..#",
	"#.#.O..#",
	"#...O..#",
	"#......#",
	"########",
	"",
	"<^^>>>vv<v>>v<<",
}

var testData1_2 = []string {
	"##########",
	"#..O..O.O#",
	"#......O.#",
	"#.OO..O.O#",
	"#..O@..O.#",
	"#O#..O...#",
	"#O..O..O.#",
	"#.OO.O.OO#",
	"#....O...#",
	"##########",
	"",
	"<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^",
	"vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v",
	"><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<",
	"<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^",
	"^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><",
	"^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^",
	">^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^",
	"<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>",
	"^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>",
	"v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^",
}

var testData2_0 = []string {
	"#######",
	"#...#.#",
	"#.....#",
	"#..OO@#",
	"#..O..#",
	"#.....#",
	"#######",
	"",
	"<vv<<^^<<^^",
}

var testSolution1_1, testSolution1_2, testSolution2 = 2028, 10092, 9021

//------------------------------------------------------

func checkSolution(testValue int, solValue int) string {
	if testValue == solValue {
		return "passed   :)"
	} else {
		return "failed   :(   (should be '"+strconv.Itoa(solValue)+"')"
	}
}

//------------------------------------------------------

func initData(fileLines []string, startChar string) ([][]rune, string, [2]int) {

	data := [][]rune {}
	moves := ""
	startInd := [2]int {-1,-1}

	firstPart := true
	for i := range fileLines {
		line := fileLines[i]

		if len(line) == 0 {
			firstPart = false
			continue
		}

		if firstPart {
			data = append(data, []rune(line))

			// get starting positions
			if strings.Contains(line, startChar) {
				startInd[0] = i
				startInd[1] = Array.GetIndexString(line, startChar)
			}
		} else {
			moves = moves + line
		}

	}

	return data, moves, startInd
}

func moveLeft1(data [][]rune, x int, y int) [2]int {

	if data[y][x-1] == '.' {
		data[y][x] = '.'
		data[y][x-1] = '@'
	} else if data[y][x-1] == 'O' {
		// move the boxes to next position
		// Get amount to move
		moveAmount := 0
		for k := range len(data[y]){
			if data[y][x-(k+2)] == '#' {
				break
			} else if data[y][x-(k+2)] == '.' {
				moveAmount = k+2
				break
			}
		}

		// Actually move the boxes
		if moveAmount > 0 {
			for k := 1; k <=moveAmount; k++ {
				data[y][x-k] = 'O'
			}
			data[y][x] = '.'
			data[y][x-1] = '@'
		} else {
			return [2]int {y, x}
		}
	}

	return [2]int {y, x-1}
}

func moveUp1(data [][]rune, x int, y int) [2]int {

	if data[y-1][x] == '.' {
		data[y][x] = '.'
		data[y-1][x] = '@'
	} else if data[y-1][x] == 'O' {
		// move the boxes to next position
		// Get amount to move
		moveAmount := 0
		for k := range len(data){
			if data[y-(k+2)][x] == '#' {
				break
			} else if data[y-(k+2)][x] == '.' {
				moveAmount = k+2
				break
			}
		}

		// Actually move the boxes
		if moveAmount > 0 {
			for k := 1; k <=moveAmount; k++ {
				data[y-k][x] = 'O'
			}
			data[y][x] = '.'
			data[y-1][x] = '@'
		} else {
			return [2]int {y, x}
		}
	}

	return [2]int {y-1, x}
}

func moveRight1(data [][]rune, x int, y int) [2]int {

	if data[y][x+1] == '.' {
		data[y][x] = '.'
		data[y][x+1] = '@'
	} else if data[y][x+1] == 'O' {
		// move the boxes to next position
		// Get amount to move
		moveAmount := 0
		for k := range len(data[y]){
			if data[y][x+k+2] == '#' {
				break
			} else if data[y][x+k+2] == '.' {
				moveAmount = k+2
				break
			}
		}

		// Actually move the boxes
		if moveAmount > 0 {
			for k := 1; k <=moveAmount; k++ {
				data[y][x+k] = 'O'
			}
			data[y][x] = '.'
			data[y][x+1] = '@'
		} else {
			return [2]int {y, x}
		}
	}

	return [2]int {y, x+1}
}

func moveDown1(data [][]rune, x int, y int) [2]int {

	if data[y+1][x] == '.' {
		data[y][x] = '.'
		data[y+1][x] = '@'
	} else if data[y+1][x] == 'O' {
		// move the boxes to next position
		// Get amount to move
		moveAmount := 0
		for k := range len(data){
			if data[y+k+2][x] == '#' {
				break
			} else if data[y+k+2][x] == '.' {
				moveAmount = k+2
				break
			}
		}

		// Actually move the boxes
		if moveAmount > 0 {
			for k := 1; k <=moveAmount; k++ {
				data[y+k][x] = 'O'
			}
			data[y][x] = '.'
			data[y+1][x] = '@'
		} else {
			return [2]int {y, x}
		}
	}
	
	return [2]int {y+1, x}
}

func makeOneMove1(data [][]rune, currPos [2]int, move rune) [2]int {

	y,x := currPos[0], currPos[1]
	newPos := [2]int {-1, -1}

	if move == '<' && data[y][x-1] != '#' {
		newPos = moveLeft1(data, x,y)
	} else if move == '>' && data[y][x+1] != '#' {
		newPos = moveRight1(data, x,y)
	} else if move == 'v' && data[y+1][x] != '#' {
		newPos = moveDown1(data, x,y)
	} else if move == '^' && data[y-1][x] != '#' {
		newPos = moveUp1(data, x,y)
	}

	if newPos[0] == -1 || newPos[1] == -1 {
		newPos[0] = y
		newPos[1] = x
	}

	return newPos
}

func solve1(data [][]rune, moves string, start [2]int, printout bool) int {

	if printout {
		Printer.PrintGridRune(data, "")
		// fmt.Println(moves)
		fmt.Println(start)
	}

	currPos := [2]int {start[0], start[1]}
	dataCopy := Array.CopyRune2D(data)

	for k := range moves {
		currPos = makeOneMove1(dataCopy, currPos, rune(moves[k]))

		// fmt.Print("Move",k, "Press 'Enter' to continue...")
  		// bufio.NewReader(os.Stdin).ReadBytes('\n') 

		// if k > 7 {
		// 	break
		// }
	}

	if printout {
		fmt.Println("\n---------")
		Printer.PrintGridRune(dataCopy, "")
	}

	sum := 0
	for i := range dataCopy {
		for j := range dataCopy {
			if dataCopy[j][i] == 'O' {
				sum = sum + 100*j + i
			}
		}
	}

	return sum
}

//----------------------------------------

func expandData(data [][]rune) ([][]rune, [2]int) {

	start := [2]int {-1, -1}
	expanded := [][]rune {}

	for i := range data {
		lineData := []rune {}
		for j := range data[i] {
			if data[i][j] == '@' {
				start[0] = i
				start[1] = len(lineData)
				lineData = append(lineData, []rune {'@','.'}...)
			} else if data[i][j] == 'O' {
				lineData = append(lineData, []rune {'[',']'}...)
			} else {
				lineData = append(lineData, []rune {data[i][j], data[i][j]}...)
			}
		}
		expanded = append(expanded, lineData)
	}

	return expanded, start
}

func moveLeft2(data [][]rune, x int, y int) [2]int {
	// fmt.Println("move left")

	if data[y][x-1] == '.' {
		data[y][x] = '.'
		data[y][x-1] = '@'
	} else if data[y][x-1] == ']' {
		// move the boxes to next position
		// fmt.Println("move box(s)")

		// Get amount to move
		moveAmount := 0
		for k := range len(data[y]){
			if data[y][x-(k+2)] == '#' {
				break
			} else if data[y][x-(k+2)] == '.' {
				moveAmount = k+2
				break
			}
		}

		// Actually move the boxes
		if moveAmount > 0 {
			// fmt.Println("move amount:", moveAmount)
			for k := x-moveAmount+1; k <= x; k++ {
				data[y][k-1] = data[y][k]
			}
			data[y][x] = '.'
		} else {
			// fmt.Println("Can't move boxes. No empty space")
			return [2]int {y, x}
		}
	}

	return [2]int {y, x-1}
}

func moveRight2(data [][]rune, x int, y int) [2]int {
	// fmt.Println("move right")

	if data[y][x+1] == '.' {
		data[y][x] = '.'
		data[y][x+1] = '@'
	} else if data[y][x+1] == '[' {
		// move the boxes to next position
		// fmt.Println("move box(s)")

		// Get amount to move
		moveAmount := 0
		for k := range len(data[y]){
			if data[y][x+(k+2)] == '#' {
				break
			} else if data[y][x+(k+2)] == '.' {
				moveAmount = k+2
				break
			}
		}

		// Actually move the boxes
		if moveAmount > 0 {
			// fmt.Println("move amount:", moveAmount)
			for k := x+moveAmount-1; k >= x; k-- {
				data[y][k+1] = data[y][k]
				// fmt.Println("move", k, "->", k+1)
			}
			data[y][x] = '.'
		} else {
			// fmt.Println("Can't move boxes. No empty space")
			return [2]int {y, x}
		}
	}

	return [2]int {y, x+1}
}

func moveUp2(data [][]rune, x int, y int) [2]int {
	// fmt.Println("move up")

	if data[y-1][x] == '.' {
		data[y][x] = '.'
		data[y-1][x] = '@'
	} else if data[y-1][x] == '[' || data[y-1][x] == ']' {
		// move the boxes to next position
		// fmt.Println("move box(s)")

		// Get amount to move
		wantToMove := [][2]int {}
		todo := [][2]int {}

		// first box
		if data[y-1][x] == ']' {
			wantToMove = append(wantToMove, [2]int {y-1, x-1})
			todo = append(todo, [2]int {y-1, x-1})
		} else if data[y-1][x] == '[' {
			wantToMove = append(wantToMove, [2]int {y-1, x})
			todo = append(todo, [2]int {y-1, x})
		}

		popIndex := 0
		canMove := true
		// look for other boxes above
		for range 999999999 {

			i,j := todo[popIndex][0], todo[popIndex][1]
			todo = append(todo[:popIndex], todo[popIndex+1:]...)

			// look up from left half of the box
			if data[i-1][j] == ']' {
				wantToMove = append(wantToMove, [2]int {i-1, j-1})
				todo = append(todo, [2]int {i-1, j-1})
			} else if data[i-1][j] == '[' {
				wantToMove = append(wantToMove, [2]int {i-1, j})
				todo = append(todo, [2]int {i-1, j})
			} else if data[i-1][j] == '#' {
				canMove = false
				// fmt.Println("can not move (left)", k, i,j)
				break
			}

			// look up from right half of the box
			if data[i-1][j+1] == ']' {
				wantToMove = append(wantToMove, [2]int {i-1, j})
				todo = append(todo, [2]int {i-1, j})
			} else if data[i-1][j+1] == '[' {
				wantToMove = append(wantToMove, [2]int {i-1, j+1})
				todo = append(todo, [2]int {i-1, j+1})
			} else if data[i-1][j+1] == '#' {
				canMove = false
				// fmt.Println("can not move (right)", k, i,j)
				break
			}

			if len(todo) == 0 {
				break
			}
		}

		// fmt.Println("want to move:", wantToMove)
		// fmt.Println("can move:", canMove)

		// Actually move the boxes
		if canMove {
			// replace boxes with '.'
			for k := range wantToMove {
				i,j := wantToMove[k][0], wantToMove[k][1]
				data[i][j] = '.'
				data[i][j+1] = '.'
				// fmt.Println("delete box", i,j, "|", i,j+1)
			}
			// give back boxes at one level higher
			for k := range wantToMove {
				i,j := wantToMove[k][0], wantToMove[k][1]
				data[i-1][j] = '['
				data[i-1][j+1] = ']'
			}

			data[y][x] = '.'
			data[y-1][x] = '@'
		} else {
			// fmt.Println("Can't move boxes. No empty space")
			return [2]int {y, x}
		}
	}

	return [2]int {y-1, x}
}

func moveDown2(data [][]rune, x int, y int) [2]int {
	// fmt.Println("move down")

	if data[y+1][x] == '.' {
		data[y][x] = '.'
		data[y+1][x] = '@'
	} else if data[y+1][x] == '[' || data[y+1][x] == ']' {
		// move the boxes to next position
		// fmt.Println("move box(es)")

		// Get boxes to move
		wantToMove := [][2]int {}
		todo := [][2]int {}
		visited := Array.InitArrayValuesInt(len(data), len(data[0]), 0)

		// first box
		if data[y+1][x] == ']' {
			wantToMove = append(wantToMove, [2]int {y+1, x-1})
			todo = append(todo, [2]int {y+1, x-1})
			visited[y+1][x-1]++
		} else if data[y+1][x] == '[' {
			wantToMove = append(wantToMove, [2]int {y+1, x})
			todo = append(todo, [2]int {y+1, x})
			visited[y+1][x]++
		}

		popIndex := 0
		canMove := true
		// look for other boxes above
		for range 999999999 {

			i,j := todo[popIndex][0], todo[popIndex][1]
			todo = append(todo[:popIndex], todo[popIndex+1:]...)
			// fmt.Println("todo", len(todo), i,j)

			// look down from left half of the box
			if data[i+1][j] == ']' && visited[i+1][j] == 0 {
				wantToMove = append(wantToMove, [2]int {i+1, j-1})
				todo = append(todo, [2]int {i+1, j-1})
				visited[i+1][j]++
			} else if data[i+1][j] == '[' && visited[i+1][j] == 0 {
				wantToMove = append(wantToMove, [2]int {i+1, j})
				todo = append(todo, [2]int {i+1, j})
				visited[i+1][j]++
			} else if data[i+1][j] == '#' {
				canMove = false
				// fmt.Println("can not move (left)", k, i,j)
				break
			}

			// // look up from left half of the box
			// if data[i-1][j] == ']' {
			// 	wantToMove = append(wantToMove, [2]int {i-1, j-1})
			// 	todo = append(todo, [2]int {i-1, j-1})
			// } else if data[i-1][j] == '[' {
			// 	wantToMove = append(wantToMove, [2]int {i-1, j})
			// 	todo = append(todo, [2]int {i-1, j})
			// } else if data[i-1][j] == '#' {
			// 	canMove = false
			// 	// fmt.Println("can not move (left)", k, i,j)
			// 	break
			// }

			// look down from right half of the box
			if data[i+1][j+1] == ']' && visited[i+1][j+1] == 0 {
				wantToMove = append(wantToMove, [2]int {i+1, j})
				todo = append(todo, [2]int {i+1, j})
				visited[i+1][j+1]++
			} else if data[i+1][j+1] == '[' && visited[i+1][j+1] == 0 {
				wantToMove = append(wantToMove, [2]int {i+1, j+1})
				todo = append(todo, [2]int {i+1, j+1})
				visited[i+1][j+1]++
			} else if data[i+1][j+1] == '#' {
				canMove = false
				// fmt.Println("can not move (right)", k, i,j)
				break
			}

			// // look up from right half of the box
			// if data[i-1][j+1] == ']' {
			// 	wantToMove = append(wantToMove, [2]int {i-1, j})
			// 	todo = append(todo, [2]int {i-1, j})
			// } else if data[i-1][j+1] == '[' {
			// 	wantToMove = append(wantToMove, [2]int {i-1, j+1})
			// 	todo = append(todo, [2]int {i-1, j+1})
			// } else if data[i-1][j+1] == '#' {
			// 	canMove = false
			// 	// fmt.Println("can not move (right)", k, i,j)
			// 	break
			// }

			if len(todo) == 0 {
				break
			}
		}

		// fmt.Println("want to move:", wantToMove)
		// fmt.Println("can move:", canMove)

		// Actually move the boxes
		if canMove {
			// replace boxes with '.'
			for k := range wantToMove {
				i,j := wantToMove[k][0], wantToMove[k][1]
				data[i][j] = '.'
				data[i][j+1] = '.'
				// fmt.Println("delete box", i,j, "|", i,j+1)
			}
			// give back boxes at one level higher
			for k := range wantToMove {
				i,j := wantToMove[k][0], wantToMove[k][1]
				data[i+1][j] = '['
				data[i+1][j+1] = ']'
			}

			data[y][x] = '.'
			data[y+1][x] = '@'
		} else {
			// fmt.Println("Can't move boxes. No empty space")
			return [2]int {y, x}
		}
	}
	
	return [2]int {y+1, x}
}

func makeOneMove2(data [][]rune, currPos [2]int, move rune) [2]int {

	// fmt.Println("position:", currPos, "move:", string(move))

	y,x := currPos[0], currPos[1]
	newPos := [2]int {-1, -1}

	if move == '<' && data[y][x-1] != '#' {
		newPos = moveLeft2(data, x,y)
	} else if move == '>' && data[y][x+1] != '#' {
		newPos = moveRight2(data, x,y)
	} else if move == 'v' && data[y+1][x] != '#' {
		newPos = moveDown2(data, x,y)
	} else if move == '^' && data[y-1][x] != '#' {
		newPos = moveUp2(data, x,y)
	}

	if newPos[0] == -1 || newPos[1] == -1 {
		newPos[0] = y
		newPos[1] = x
	}

	// fmt.Println("New state")
	// Printer.PrintGridRune(data, "")
	// fmt.Println()

	return newPos
}

func solve2(data [][]rune, moves string, printout bool) int {

	expanded, start := expandData(data)

	if printout {
		Printer.PrintGridRune(expanded, "")
		fmt.Println(moves)
		fmt.Println(start)
	}

	currPos := [2]int {start[0], start[1]}
	dataCopy := Array.CopyRune2D(expanded)

	for k := range moves {
		currPos = makeOneMove2(dataCopy, currPos, rune(moves[k]))

		// fmt.Print("Move ",k, " - Press 'Enter' to continue...")
  		// bufio.NewReader(os.Stdin).ReadBytes('\n') 

		// fmt.Println("move", k, len(moves))
		// if k > 10 {
		// 	break
		// }
	}

	if printout {
		fmt.Println("\n--------- dataCopy")
		Printer.PrintGridRune(dataCopy, "")
	}

	sum := 0
	for i := range dataCopy {
		for j := range dataCopy[i] {
			if dataCopy[i][j] == '[' {
				sum = sum + j + 100*i
			}
		}
	}
	
	return sum
}

func main() {

	// data gathering and parsing
	startChar := "@"
	dataTest1_1, movesTest1_1, startTest1_1 := initData(testData1_1, startChar)
	dataTest1_2, movesTest1_2, startTest1_2 := initData(testData1_2, startChar)
	dataTest2_0, movesTest2_0, _ := initData(testData2_0, startChar)

	fileName := "day_15_data.txt"
	fileData := rw.ReadFile(fileName)
	data, moves, start := initData(fileData, startChar)

	// ---------------------------------------------
	fmt.Println("=== Part 1 ===")
	sol1_1_test := solve1(dataTest1_1, movesTest1_1, startTest1_1, false)
	fmt.Println("Test solution 1_1 =", sol1_1_test, "->", checkSolution(sol1_1_test, testSolution1_1))
	sol1_2_test := solve1(dataTest1_2, movesTest1_2, startTest1_2, false)
	fmt.Println("Test solution 1_2 =", sol1_2_test, "->", checkSolution(sol1_2_test, testSolution1_2))

	t1 := time.Now()
	sol1 := solve1(data, moves, start, false)
	dur := time.Since(t1)
	fmt.Println("Solution part 1 =", sol1, "(ET =", dur, ")")

	// ---------------------------------------------
	fmt.Println()
	fmt.Println("=== Part 2 ===")
	solve2(dataTest2_0, movesTest2_0, true)
	sol2_test := solve2(dataTest1_2, movesTest1_2, true)
	fmt.Println("Test solution 2 =", sol2_test, "->", checkSolution(sol2_test, testSolution2))

	t1 = time.Now()
	sol2 := solve2(data, moves, false)
	dur = time.Since(t1)
	fmt.Println("Solution part 2 =", sol2, "(ET =", dur, ")")
}
