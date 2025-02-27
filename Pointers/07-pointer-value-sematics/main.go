package main

import "fmt"

// pointer semantics
func addP(v *int) {
	*v += 1
}

// value semantics
func add(v int) int {
	return v + 1
}

func main() {
	//pointer semantics
	b := 1
	fmt.Println(b)
	addP(&b)
	fmt.Println(b) // changed the original value

	// value semantics
	a := 1
	fmt.Println(a)      // 1
	fmt.Println(add(a)) // 2 // its just copy the value but don't change the original
	fmt.Println(a)      // 1
}
