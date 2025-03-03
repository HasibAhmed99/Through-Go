package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Generic Max function using constraints.Ordered
// Generic function to find the maximum of two values
func Max[T constraints.Ordered](a, b T) T { // constraints.Ordered builtin package
	if a > b {
		return a // if a greater than b return a
	}
	return b // else return b
	//  We don't need to write (else) here because Go understand it without writing it.
}

func main() {
	fmt.Println(Max(5, 11))             // works with int
	fmt.Println(Max(3.36, 2.96))        // works with float64
	fmt.Println(Max("apple", "banana")) // works with string
	fmt.Println(Max("cat", "dog"))
}

// Theory:
/*
How Does It Work?
T constraints.Ordered means that T must be one of the following types:
Numbers: int, float64, uint, etc.
Strings: string
These types support comparison using the operators >, <, >=, and <=, which is why they are considered ordered.
These types does't support arithmetic operations (+, -, *, /)
If you tried to pass a non-ordered type (like a struct or map), you would get a compile-time error because these types don’t support comparison with ordering operators.
*/

/*
✅ Solution: Enable Go Modules & Install constraints
Since golang.org/x/exp/constraints is not part of the Go standard library, You need to use Go modules to manage it.
Step 1: Initialize a Go Module
Run => go mod init myproject
This will create a go.mod file in your directory.

Step 2: Install the Constraints Package
Once inside a Go module, install the package properly:
Run => go get golang.org/x/exp/constraints

Step 3: Verify Installation
Run => go mod tidy
This cleans up unused dependencies and ensures golang.org/x/exp/constraints is installed.
*/
