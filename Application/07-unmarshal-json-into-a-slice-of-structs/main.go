package main

import (
	"encoding/json"
	"fmt"
)

// Book represents a book with a title and author.
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	// JSON input as an array of objects
	jsonInput := `[{"title":"Go Programming","author":"John Doe"},{"title":"Advanced Go","author":"Jane Smith"}]`

	// Creating a slice to store the decoded data
	var bks []Book

	// Decoding JSON array into Go slice of structs
	err := json.Unmarshal([]byte(jsonInput), &bks)
	if err != nil {
		fmt.Println("Error unmarshaling", err)
		return
	}
	fmt.Println(bks)

	for _, v := range bks {
		fmt.Println("Title:", v.Title)
		fmt.Println("Author:", v.Author)
	}
}
