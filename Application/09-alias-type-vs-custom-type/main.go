package main

import "fmt"

// Alias Type: This does NOT create a new type, it just renames int
type MyAliasInt = int

// Custom Type: This creates a NEW distinct type
type MyCustomInt int

func main() {
	// Using Alias Type
	var a MyAliasInt = 100
	var b int = 200

	// Since MyAliasInt is just another name for int, they are interchangeable
	a = b                                 // ✅ Allowed, since MyAliasInt is an alias for int
	fmt.Println("Alias Type Example:", a) // Output: 200

	// Using Custom Type
	var x MyCustomInt = 300
	var y int = 400

	// x = y // ❌ Error: cannot use y (type int) as type MyCustomInt
	x = MyCustomInt(y)                     // ✅ Explicit conversion required
	fmt.Println("Custom Type Example:", x) // Output: 400

	// Printing types
	fmt.Printf("Type of a: %T\n", a) // Output: int
	fmt.Printf("Type of x: %T\n", x) // Output: main.MyCustomInt (distinct type)
}
