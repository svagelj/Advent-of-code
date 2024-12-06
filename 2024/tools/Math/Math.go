package tools

import (
	"strconv"
)

// Functions with names starting with upper case are exported

func AbsInt(a int) int {

	if a < 0 {
		return -a
	} else {
		return a
	}
}

func MinInt(a int, b int) int {

	if a > b {
		return b
	}

	return a
}

func MaxInt(a int, b int) int {
	
	if a > b {
		return a
	}

	return b
}

func StringToInt(str string) int {

	_int, err := strconv.Atoi(str)
	if err == nil {
		return _int
	} else {
		panic(err)
	}
}

func MatrixDotVector(matrix [][]int, vector []int) []float32 {

	result := []float32 {}

	for i := range matrix {

		line := float32(0)
		for j := range vector {
			line = line + float32(matrix[i][j]) * float32(vector[j])
		}
		result = append(result, line)
	}

	return result
}