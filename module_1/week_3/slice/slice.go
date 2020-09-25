package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	slice := make([]int, 3)

	for {
		var input string
		fmt.Println("Please, enter integer number: ")
		fmt.Scanf("%s", &input)

		input = strings.ToLower(input)
		input = strings.Trim(input, "\n")
		input = strings.Trim(input, "\r")
		if input == "x" {
			fmt.Println("Goodbye!")
			break
		}

		number, _ := strconv.Atoi(input)
		slice = append(slice, number)
		sort.Ints(slice)
		fmt.Println(slice)
	}
}
