package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter text: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	cleanText := strings.Trim(input, "\n")
	cleanText = strings.Trim(cleanText, "\r")
	cleanText = strings.ToLower(cleanText)

	foundI := strings.HasPrefix(cleanText, "i")
	foundN := strings.HasSuffix(cleanText, "n")
	foundA := strings.Index(cleanText, "a")

	if foundI && foundN && foundA >= 0 {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
