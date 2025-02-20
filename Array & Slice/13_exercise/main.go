package main

import "fmt"

func main() {
	// creating an array
	x := [5]int{}

	// assigning values to each  index position
	for  i := 0; i < 5; i++ {
		x[i] = i
	}

	for i, v := range x {
		fmt.Printf("%v - %T - %v\n", v, v, i)
	}
}