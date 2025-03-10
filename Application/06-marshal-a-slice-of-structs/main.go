package main

import (
	"encoding/json"
	"fmt"
)

// Car represents a vehicle with a model and price.
type Car struct {
	Model string `json:"model"`
	Price int    `json:"price"`
}

func main() {
	// Creating a slice of Car structs
	crs := []Car{
		{"Tesla Model S", 79999},
		{"Ford Mustang", 55999},
	}

	// Encoding slice of structs to JSON
	jsonData, err := json.Marshal(crs)
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}
	fmt.Println(string(jsonData))
}
