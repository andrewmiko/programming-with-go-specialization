package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	fmt.Println("Welcome to Animal app!")
	fmt.Println("Please choose animal and information that you want to receive.")
	fmt.Println("Animal options: cow, bird, snake")
	fmt.Println("Information options: eat, move, speak")
	fmt.Println("Example: cow eat")

	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	for {
		fmt.Print(">")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		cleanText := strings.Trim(input, "\n")
		cleanText = strings.Trim(cleanText, "\r")
		cleanText = strings.ToLower(cleanText)

		commandSlice := strings.Split(cleanText, " ")
		animal, action := commandSlice[0], commandSlice[1]

		var _animal Animal
		switch animal {
		case "cow":
			_animal = cow
		case "bird":
			_animal = bird
		case "snake":
			_animal = snake
		default:
			fmt.Println("There is no such animal")
		}

		switch action {
		case "eat":
			_animal.Eat()
		case "move":
			_animal.Move()
		case "speak":
			_animal.Speak()
		default:
			fmt.Println("There is no such action")
		}
	}
}
