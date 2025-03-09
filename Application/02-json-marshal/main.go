package main

import (
	"encoding/json" // Provides functions for encoding and decoding JSON data.
	"fmt"
)

// Person represents an individual with basic attributes.
type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	// Initializing sample Person instances.
	p1 := person{
		First: "Bob",
		Last:  "Marley",
		Age:   36,
	}

	p2 := person{
		First: "Harry",
		Last:  "Potter",
		Age:   44,
	}

	// Creating a slice of Person objects.
	people := []person{p1, p2}
	fmt.Println(people)

	// Encoding the slice of People into JSON format.
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}

	// Printing the JSON-encoded result as a string.
	fmt.Println(string(bs))
}
