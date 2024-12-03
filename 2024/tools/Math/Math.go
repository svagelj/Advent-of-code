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
