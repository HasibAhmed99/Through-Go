package main

import "fmt"

type person struct {
	first string
	last  string
	favIC []string
}

func main() {

	p1 := person{
		first: "Tom",
		last:  "Cruise",
		favIC: []string{"Strawberry", "Chocolate", "passion fruit with mango and guava"},
	}

	p2 := person{
		first: "Steve",
		last:  "Jobs",
		favIC: []string{"Vanilla", "Mint"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v)
		for _, v2 := range v.favIC {
			fmt.Println(v.first, v.last, v2)
		}
	}
}
