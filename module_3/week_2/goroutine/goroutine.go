package main

import (
	"fmt"
	"time"
)

func increase_number(number *int) {
	fmt.Println("increase_number")
	*number += 1
}

func double_number(number *int) {
	fmt.Println("double_number")
	*number *= 2
}

func main() {
	x := 1
	go increase_number(&x)
	go double_number(&x)
	time.Sleep(10 * time.Millisecond)

	// We assume that execution will be as:
	// 1 -> increase_number(1) => 2
	// 2 -> double_number(2) => 4
	// but, it can be as (sometimes):
	// 1 -> double_number(1) => 2
	// 2 -> increase_number(1) => 3
	// This is example of race condition, because:
	// 1) double_number executed first before increase_number
	// 2) we worked on same value (variable) which shows race condition conflict

	fmt.Println(x)
}
