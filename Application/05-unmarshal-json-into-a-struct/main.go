package main

import (
	"encoding/json"
	"fmt"
)

// User represents a user with a name and age.
type User struct {
	Name string `json:"username"`
	Age  int    `json:"age"`
}

func main() {
	// JSON string input (received from API, file, etc.)
	jsonInput := `{"username":"JohnDoe","age":30}`

	// Creating a variable to store the decoded data
	var ur User

	// Decoding JSON data into Go struct
	err := json.Unmarshal([]byte(jsonInput), &ur)
	if err != nil {
		fmt.Println("Error unmarshaling:", err)
		return
	}
	// Printing the decoded struct
	fmt.Println(ur)
	fmt.Println(ur.Name)
	fmt.Println(ur.Age)
}
