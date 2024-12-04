package tools

import (
	"bufio"
	"fmt"
	"os"
)

// Functions with names starting with upper case are exported

func ReadFile(fileName string) []string {

	fmt.Println("Reading file '" + fileName + "'")

	// open the file
	f, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}

	// Read line by line
	scanner := bufio.NewScanner(f)
	var fileLines []string
	for scanner.Scan() {
		line := scanner.Text()
		fileLines = append(fileLines, line)
	}

	// Check for errors
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	f.Close()

	return fileLines

}

func ReverseString(inputStr string) string {

	reversed := ""

	for i := len(inputStr) - 1; i >= 0; i-- {
		reversed += string(inputStr[i])
	}

	return reversed
}