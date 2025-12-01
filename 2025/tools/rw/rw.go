package tools

import (
	"bufio"
	"fmt"
	"os"
)

func checkError(e error) {
    if e != nil {
        panic(e)
    }
}

// Functions with names starting with upper case are exported

func ReadFile(fileName string) []string {

	fmt.Println("Reading file '" + fileName + "'")

	// open the file
	f, err := os.Open(fileName)
	checkError(err)
	defer func() {
        if err := f.Close(); err != nil {
            panic(err)
        }
    }()

	// Read line by line
	scanner := bufio.NewScanner(f)
	var fileLines []string
	for scanner.Scan() {
		line := scanner.Text()
		fileLines = append(fileLines, line)
	}

	return fileLines

}

func WriteFile(fileName string, mode rune, content string) {

	var (
		f *os.File
		err error
	)

	if mode == 'a' {
		f, err = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	} else if mode == 'w' {
		f, err = os.Create(fileName)
	} else {
		panic("Wrong mode was given. Expecting [w, a]")
	}

	checkError(err)
    defer func() {
        if err := f.Close(); err != nil {
            panic(err)
        }
    }()	

    w := bufio.NewWriter(f)
    n4, err := w.WriteString(content)
    checkError(err)
    fmt.Printf("wrote %d bytes\n", n4)

    if err = w.Flush(); err != nil {
        panic(err)
    }
}

func WriteFileSlice(fileName string, mode rune, content []string) {

	var (
		f *os.File
		err error
	)

	if mode == 'a' {
		f, err = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	} else if mode == 'w' {
		f, err = os.Create(fileName)
	} else {
		panic("Wrong mode was given. Expecting [w, a]")
	}

	checkError(err)
    defer func() {
        if err := f.Close(); err != nil {
            panic(err)
        }
    }()	

    w := bufio.NewWriter(f)
	for i := range content {
    	// n4, err := w.WriteString(content[i])
		_, err := w.WriteString(content[i]+"\n")
		checkError(err)

		// fmt.Printf("wrote %d bytes\n", n4)
	}

	if err = w.Flush(); err != nil {
		panic(err)
	}
}