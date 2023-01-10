package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {

	var filename string
	var filehandle *os.File

	// get the file name
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your file name: ")
	filename, err1 := reader.ReadString('\n')
	filename = strings.TrimSuffix(filename, "\n")
	persons := make([]Person, 0, 1)

	// catch input error
	if err1 != nil {
		fmt.Println("Input error.")
		os.Exit(0)
	}

	// create a file handle
	for {
		fh, err2 := os.Open(filename)
		if err2 != nil {
			fmt.Println("Error opening file. Try another file name.")
			fmt.Print("Enter your file name: ")
			filename, _ := reader.ReadString('\n')
			filename = strings.TrimSuffix(filename, "\n")
		} else {
			filehandle = fh
			break
		}
	}
	filescanner := bufio.NewScanner(filehandle)
	var lines []string

	for filescanner.Scan() {
		lines = strings.Split(filescanner.Text(), " ")
		persons = append(persons, Person{lines[0], lines[1]})
	}

	filehandle.Close()

	for _, person := range persons {
		fmt.Print("%v - %v\n", person.fname, person.lname)
	}
}
