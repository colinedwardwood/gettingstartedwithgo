package main

import (
	"fmt"
	"strings"
)

func main() {
	var rawinputstr string
	var lowerinputstr string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&rawinputstr)
	lowerinputstr = strings.ToLower(rawinputstr)
	if string(lowerinputstr[0]) == "i" && string(lowerinputstr[len(lowerinputstr)-1]) == "n" && strings.Contains(lowerinputstr, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
