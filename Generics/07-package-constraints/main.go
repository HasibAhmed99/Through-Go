package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// addI add two integer values and returns the result.
func addI(a, b int) int {
	return a + b
}

// addF add two floating-point values and returns the result.
func addF(a, b float64) float64 {
	return a + b
}

// myNums is a type constraint that allows only integers and floating-point numbers.
// This ensures that generics only work with numeric types.
type myNums interface {
	constraints.Integer | constraints.Float
}

// addT is a generic function that adds two number of type T.
// T must satisfy the myNums constraint, meaning it can be an integer or a float.
func addT[T myNums](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(addI(5, 5))
	fmt.Println(addF(4.05, 0.95))

	fmt.Println(addT(5, 5))
	fmt.Println(addT(4.05, 0.95))
}
