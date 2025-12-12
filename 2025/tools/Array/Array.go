package tools

// Functions with names starting with upper case are exported

func GetIndexInt(array []int, element int) int {
	for i := range array {
		if array[i] == element {
			return i
		}
	}
	return -1
}

func GetIndexInt2D(array [][]int, element []int) int {
	for i := range array {
		same := 0
		for j := range element {
			if array[i][j] == element[j] {
				same++
			}
		}
		if same == len(element) {
			return i
		}
	}
	return -1
}

func GetIndexString(array string, element string) int {
	for i := range array {
		if array[i] == element[0] {
			return i
		}
	}
	return -1
}

func ReverseString(inputStr string) string {

	reversed := ""

	for i := len(inputStr) - 1; i >= 0; i-- {
		reversed += string(inputStr[i])
	}

	return reversed
}

func InitArrayValuesInt(M int, N int, value int) [][] int {
	array := [][]int {}

	for range M {
		_line := []int {}
		for range N {
			_line = append(_line, value)
		}
		array = append(array, _line)
	}

	return array
}

func InitArrayValuesRune(M int, N int, value rune) [][] rune {
	array := [][]rune {}

	for range M {
		_line := []rune {}
		for range N {
			_line = append(_line, value)
		}
		array = append(array, _line)
	}

	return array
}

func InitArrayValuesInt4D(M int, N int) [][][][] int {
	array := [][][][]int {}

	for range M {
		_line := [][][]int {}
		for range N {
			__line := [][] int{}
			__line = append(__line, []int {})
			_line = append(_line, __line)
		}
		array = append(array, _line)
	}

	return array
}

func CopyRune2D(data [][]rune) [][]rune {

	duplicate := make([][]rune, len(data))
	for i := range data {
		duplicate[i] = make([]rune, len(data[i]))
		copy(duplicate[i], data[i])
	}
	return duplicate
}

func CopyString2D(data [][]string) [][]string {

	duplicate := make([][]string, len(data))
	for i := range data {
		duplicate[i] = make([]string, len(data[i]))
		copy(duplicate[i], data[i])
	}
	return duplicate
}

func CopyInt2D(data [][]int) [][]int {

	duplicate := make([][]int, len(data))
	for i := range data {
		duplicate[i] = make([]int, len(data[i]))
		copy(duplicate[i], data[i])
	}
	return duplicate
}

func CopyFloat642D(data [][]float64) [][]float64 {

	duplicate := make([][]float64, len(data))
	for i := range data {
		duplicate[i] = make([]float64, len(data[i]))
		copy(duplicate[i], data[i])
	}
	return duplicate
}