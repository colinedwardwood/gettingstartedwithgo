package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	ints := make([]int, 0, 3) // create the empty slice
	var userinputstr string   // allocate space for input
	var userinputnum int      // allocate space for numeric version of input
	for {

		// get user input
		fmt.Print("Enter an integer: ")
		_, err := fmt.Scan(&userinputstr)

		// check for input error
		if err != nil {
			fmt.Println("Input error.")
			fmt.Println(err)
			os.Exit(0)
		}

		// check for exit code being received
		if strings.ToLower(userinputstr) == "x" {
			fmt.Println("Exit code received.")
			os.Exit(0)
		}

		// try to convert the string to integer
		userinputnum, err = strconv.Atoi(userinputstr)
		if err != nil {
			fmt.Println("Input must be limited to integers or x to exit.")
			fmt.Println(err)
			os.Exit(0)
		}

		ints = append(ints, userinputnum)
		sort.Ints(ints)
		fmt.Println(ints)
	}
}
