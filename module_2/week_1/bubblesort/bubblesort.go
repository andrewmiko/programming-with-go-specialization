package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Swap(slice []int, index int) {
	slice[index], slice[index+1] = slice[index+1], slice[index]
}

func BubbleSort(slice []int) {
	end := len(slice) - 1
	for {
		if end == 0 {
			break
		}

		for i := 0; i < len(slice)-1; i++ {
			if slice[i] > slice[i+1] {
				Swap(slice, i)
			}
		}

		end -= 1
	}
}

func main() {
	slice := make([]int, 0)

	var input string
	fmt.Println("Please, enter sequence of 10 integers: ")
	fmt.Scanf("%s", &input)

	numbers := strings.Split(input, ",")
	for index, number := range numbers {
		if index <= 9 {
			number, _ := strconv.Atoi(number)
			slice = append(slice, number)
		}
	}

	BubbleSort(slice)
	fmt.Println(slice)
}
