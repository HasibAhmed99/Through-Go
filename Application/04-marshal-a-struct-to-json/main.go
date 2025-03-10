package main

import (
	"encoding/json"
	"fmt"
)

// Product represents an item with a name and price.
type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	// Creating a product instance
	pd := Product{
		Name:  "Laptop",
		Price: 999.99,
	}

	// Converting the Go struct to JSON Format
	jsonData, err := json.Marshal(pd)
	if err != nil {
		fmt.Println("Error marshaling", err)
		return // Stop execution if an error occurs.
	}

	// Printing the JSON output as a string.
	fmt.Println(string(jsonData))
}
