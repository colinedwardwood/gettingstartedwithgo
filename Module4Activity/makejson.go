package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {

	var usrnameinputstr string // allocate space for input as string
	var usraddrinputstr string // allocate space for input as string
	m := make(map[string]string)

	// get user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	usrnameinputstr, err1 := reader.ReadString('\n')
	usrnameinputstr = strings.TrimSuffix(usrnameinputstr, "\n")
	fmt.Print("Enter your address: ")
	usraddrinputstr, err2 := reader.ReadString('\n')
	usraddrinputstr = strings.TrimSuffix(usraddrinputstr, "\n")

	// check for input error
	if err1 != nil || err2 != nil {
		fmt.Println("Input error.")
		os.Exit(0)
	}

	// populate the map
	m["name"] = usrnameinputstr
	m["address"] = usraddrinputstr

	// marshall a json object from the map
	bytes, _ := json.Marshal(m)

	// output the json object
	fmt.Println(string(bytes))
}
