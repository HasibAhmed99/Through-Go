package main

import (
	"encoding/json"
	"fmt"
)

// person represents an individual with First name, Last name, Age and collection of sayings.
type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	// Simulated JSON data, typically received from an API response or external source.
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`

	// Printing raw JSON data for reference
	fmt.Println(s)

	// Declaring a variable to store the parsed JSON data as a slice of person structs.
	var people []person

	// Unmarshaling(decoding) the JSON data into the 'people' slice
	err := json.Unmarshal([]byte(s), &people)
	if err != nil {
		fmt.Println("Unmarshaling error")
		return // Exiting if there's an error to prevent further execution
	}

	// Indicating that JSON parsing was successful
	fmt.Println("Unmarshaling Success")

	// Printing the entire parsed data structure
	fmt.Println(people)

	// Iterating over the list of people to display their details in a structured format
	for i, value := range people {
		fmt.Println("Person #", i, value)
		fmt.Println("\t", value.First, value.Last, value.Age)
		// Iterating over and displaying each person's sayings
		for _, saying := range value.Sayings {
			fmt.Println("\t\t", saying)
		}
	}
}
