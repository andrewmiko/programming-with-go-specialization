package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var path string
	fmt.Print("Please, enter text file path: ")
	fmt.Scan(&path)

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	names := make([]Name, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		name := scanner.Text()
		nameSlice := strings.Split(name, " ")
		nameStruct := Name{nameSlice[0], nameSlice[1]}
		names = append(names, nameStruct)
	}

	for _, name := range names {
		fmt.Println(name.fname, name.lname)
	}
}
