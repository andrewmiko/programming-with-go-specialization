package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

type Cow struct{}
type Bird struct{}
type Snake struct{}

func (c Cow) Eat() string {
	return "grass"
}

func (c Cow) Move() string {
	return "walk"
}

func (c Cow) Speak() string {
	return "moo"
}

func (b Bird) Eat() string {
	return "worms"
}

func (b Bird) Move() string {
	return "fly"
}

func (b Bird) Speak() string {
	return "peep"
}

func (s Snake) Eat() string {
	return "mice"
}

func (s Snake) Move() string {
	return "slither"
}

func (s Snake) Speak() string {
	return "hsss"
}

func cleanText(s string) string {
	text := strings.Trim(s, "\n")
	text = strings.Trim(text, "\r")
	text = strings.ToLower(text)
	return text
}

func createAnimal() (string, Animal, error) {
	var animal Animal
	var animalName, animalType string

	fmt.Println("Please enter new animal details: ")
	fmt.Print("Name (any): ")
	fmt.Scanf("%s", &animalName)
	animalName = cleanText(animalName)

	fmt.Print("Type (cow, bird, snake): ")
	fmt.Scanf("%s", &animalType)
	animalType = cleanText(animalType)

	switch animalType {
	case "cow":
		animal = Cow{}
	case "bird":
		animal = Bird{}
	case "snake":
		animal = Snake{}
	default:
		msg := fmt.Sprintf("Cannot set animal type as %s", animalType)
		return "", nil, errors.New(msg)
	}

	fmt.Println("Animal created!")
	return animalName, animal, nil
}

func queryAnimal(animals map[string]Animal) (string, error) {
	var result string
	var animal Animal
	var animalName, animalAction string

	fmt.Println("Please enter animal details: ")
	fmt.Print("Name (any): ")
	fmt.Scanf("%s", &animalName)
	animalName = cleanText(animalName)

	if val, ok := animals[animalName]; ok {
		animal = val
	} else {
		msg := fmt.Sprintf("Cannot find animal by name: %s", animalName)
		return "", errors.New(msg)
	}

	fmt.Print("Action (eat, move, speak): ")
	fmt.Scanf("%s", &animalAction)
	animalAction = cleanText(animalAction)

	switch animalAction {
	case "eat":
		result = animal.Eat()
	case "move":
		result = animal.Move()
	case "speak":
		result = animal.Speak()
	default:
		msg := fmt.Sprintf("Cannot find animal action: %s", animalAction)
		return "", errors.New(msg)
	}

	return result, nil
}

func main() {
	animals := make(map[string]Animal)
	fmt.Println("Welcome to Animal app!")

	for {
		fmt.Println("Please choose action: 'newanimal', 'query' or 'exit'")
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		command := cleanText(input)
		if command == "newanimal" {
			name, animal, err := createAnimal()
			if err != nil {
				fmt.Println(err)
			} else {
				animals[name] = animal
				fmt.Println(animals)
			}
		} else if command == "query" {
			result, err := queryAnimal(animals)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}
		} else if command == "exit" {
			os.Exit(0)
		} else {
			msg := fmt.Sprintf("Cannot understand command: %s", command)
			fmt.Println(msg)
		}
	}
}
