package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func sortSlice(slice []int, c chan []int) {
	fmt.Printf("[GORUTINE] Unsorted slice: %v\n", slice)
	sort.Ints(slice)
	fmt.Printf("[GORUTINE] Sorted slice: %v\n", slice)
	c <- slice
}

func main() {
	var input string
	const goroutinesCount = 4

	fmt.Print("Please, enter sequence of integers (e.g. 1,2,3,n): ")
	fmt.Scanf("%s", &input)

	// Preparing slice of integers
	integersSequence := make([]int, 0)
	integers := strings.Split(input, ",")
	for _, n := range integers {
		number, _ := strconv.Atoi(n)
		integersSequence = append(integersSequence, number)
	}

	c := make(chan []int)
	slices := make([]int, 0)
	sliceSize := len(integersSequence) / goroutinesCount
	sliceSizeRest := len(integersSequence) % goroutinesCount
	for i := 0; i <= goroutinesCount-1; i++ {
		unorderedSlice := make([]int, 0)

		// If we cant split slices equally on goroutinesCount slices,
		// then on last iteration we add last element to last slice
		if i == (goroutinesCount-1) && sliceSizeRest > 0 {
			unorderedSlice = integersSequence[0 : sliceSize+1]
		} else {
			unorderedSlice = integersSequence[0:sliceSize]
		}
		integersSequence = integersSequence[sliceSize:]

		// Sort in alternative goroutine
		// and sync between channels
		go sortSlice(unorderedSlice, c)
		orderedSlice := <-c

		// Appending ordered slice to final slice
		slices = append(slices, orderedSlice...)
	}

	// IMPORTANT!
	// I don't know, do I need to sort final slice here or no,
	// but it looks better when it's sorted
	sort.Ints(slices)
	fmt.Printf("[MAIN] Sorted list: %v", slices)
}
