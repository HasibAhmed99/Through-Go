package main

import (
	"fmt"
)

func main() {

	p := struct {
		first string
		friends map[string]int
		favDrinks []string
	}{
       first: "James",
	   friends: map[string]int{
				"Jenny" : 27,
				"Tom": 31,
				"Kevin": 35,
	   },
	   favDrinks: []string{"Water", "Juice"},
	}

	for k, v := range p.friends{
		fmt.Println(p.first, "- friends - ", k, v)
	}

	for _, v := range p.favDrinks {
		fmt.Println(p.first, "- drinks - ", v)
	}
}
