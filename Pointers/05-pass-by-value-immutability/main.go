package main

import "fmt"

func modifyInt(n int) {
	n = 100                            // 'n' is a local copy inside the function
	fmt.Println("Inside function:", n) // Shows 100
}

func main() {
	x := 42
	modifyInt(x) // Passing by value (creates a copy)
	// Still 42
	fmt.Println("Outside function:", x)
}
