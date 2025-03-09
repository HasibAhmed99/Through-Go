package main

import (
	"encoding/json" // For JSON encoding/decoding.
	"fmt"
)

// Person represents an individual with basic attributes.
type person struct {
	First string `json:"First"` // First name mapped to "First" in JSON.
	Last  string `json:"Last"`  // Last name mapped to "Last" in JSON.
	Age   int    `json:"Age"`   // Age mapped to "Age" in JSON
}

func main() {
	// Json string representing an array of person objects.
	s := `[{"First":"Bob", "Last":"Marley", "Age": 36}, {"First":"Harry", "Last":"Potter", "Age":44}]`
	// Convert string to byte slice for unmarshaling
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	// Slice to hold unmarshal data.
	//people := []person{}
	var people []person

	// Unmarshal the JSON into the "people" slice.
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	// Print all unmarshalled data.
	fmt.Println("\nall of the data", people)

	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
