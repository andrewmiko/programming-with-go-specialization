package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func chunkSlice(slice []int, chunkSize int) [][]int {
	var chunks [][]int

	for {
		if len(slice) == 0 {
			break
		}

		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}

		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}

	return chunks
}

func sortSlice(slice []int, c chan []int) {
	fmt.Printf("[GORUTINE] Unsorted slice: %v\n", slice)
	sort.Ints(slice)
	fmt.Printf("[GORUTINE] Sorted slice: %v\n", slice)
	c <- slice
}

func main() {
	var input string

	fmt.Print("Please, enter sequence of integers (e.g. 1,2,3,n): ")
	fmt.Scanf("%s", &input)

	integersSequence := make([]int, 0)
	integers := strings.Split(input, ",")
	for _, n := range integers {
		number, _ := strconv.Atoi(n)
		integersSequence = append(integersSequence, number)
	}

	chunksCount := 4
	chunkSize := len(integersSequence) / chunksCount
	chunkSizeRest := len(integersSequence) % chunksCount
	chunks := chunkSlice(integersSequence, chunkSize)
	//for i := 1; i <= chunksCount; i++ {
	//	if i == chunksCount && chunkSizeRest != 0 {
	//		limit := chunkSize
	//	} else {}
	//
	//	slice = integersSequence
	//
	//}

	fmt.Println(integersSequence)
	fmt.Println(chunkSize)
	fmt.Println(chunkSizeRest)
	fmt.Println(chunks)

	c := make(chan []int)
	sortedSlices := make([]int, 0)
	for _, slice := range chunks {
		go sortSlice(slice, c)
		sortedSlice := <-c
		sortedSlices = append(sortedSlices, sortedSlice...)
	}

	// TODO
	sort.Ints(sortedSlices)
	fmt.Printf("[MAIN] Sorted list: %v", sortedSlices)
}
