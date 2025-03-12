package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//JSON input with dynamic keys
	jsonInput := `{"name":"Alice","age":26,"city":"Berlin"}`

	// Creating a map to hold JSON key-value pairs
	var result map[string]interface{}

	// Decoding JSON into the map
	err := json.Unmarshal([]byte(jsonInput), &result)
	if err != nil {
		fmt.Println("Error Unmarshaling", err)
		return
	}
	// Printing the decoded map
	fmt.Println(result)

	fmt.Println(result["name"])
	fmt.Println(result["age"])
	fmt.Println(result["city"])

}
