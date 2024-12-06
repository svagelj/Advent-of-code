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

func GetIndexString(array string, element string) int {
	for i := range array {
		if array[i] == element[0] {
			return i
		}
	}
	return -1
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