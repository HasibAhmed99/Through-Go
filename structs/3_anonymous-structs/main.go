package main

import (
	"fmt"
)

type car struct {
	name  string
	model string
	speed int
}

func main() {  
	x := struct{
		first string
		last string
		age int
	}{
		first: "Rick",
		last: "Morty",
		age: 24,
	}

	c := car{
		name: "BMW",
		model: "iX5",
		speed: 185,
	}

	fmt.Printf("%#v \t %T\n", x, x)
	fmt.Printf("%#v \t %T\n", c, c)
}