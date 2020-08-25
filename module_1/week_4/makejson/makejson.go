package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, address string
	fmt.Print("Please, enter your name: ")
	fmt.Scan(&name)
	fmt.Print("Please, enter your address: ")
	fmt.Scan(&address)

	data := map[string]string{
		"name":    name,
		"address": address,
	}

	json, _ := json.Marshal(data)
	fmt.Println(string(json))
}
