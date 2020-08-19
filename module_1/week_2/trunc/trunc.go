package main

import (
	"fmt"
)

func main() {
	var input float64
	fmt.Print("Please enter float number: ")
	fmt.Scan(&input)
	fmt.Println(int64(input))
}
