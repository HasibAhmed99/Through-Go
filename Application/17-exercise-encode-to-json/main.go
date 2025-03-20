package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// user represents an individual with a First name, Last name, Age and collections of sayings.
type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {

	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   33,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Harry",
		Last:  "Potter",
		Age:   44,
		Sayings: []string{"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "Bob",
		Last:  "Marley",
		Age:   36,
		Sayings: []string{"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	enc := json.NewEncoder(os.Stdout)
	err := enc.Encode(users)
	if err != nil {
		fmt.Println("Something went wrong!")
	}

	// err := json.NewEncoder(os.Stdout).Encode(users)
	// if err != nil {
	// 	fmt.Println("Gone something wrong and here is the error:", err)
	// }
}
