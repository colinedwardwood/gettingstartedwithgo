package main

import (
	"fmt"
	"strconv"
)

func main() {
	var vals string
	var valf float64
	fmt.Print("Enter a floating point number: ")
	fmt.Scanf("%s", &vals)
	valf, _ = strconv.ParseFloat(vals, 64)
	fmt.Printf("%f as an integer is %d", valf, int(valf))
}
