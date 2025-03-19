package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "Harry",
		Age:   44,
	}

	u2 := user{
		First: "Bob",
		Age:   36,
	}

	u3 := user{
		First: "Eve",
		Age:   35,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Error Marshaling")
		return
	}
	fmt.Println("Marshaling Success")

	fmt.Println(string(bs))
}
