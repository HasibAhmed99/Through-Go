package main

import "fmt"

func foo() int {
	return 42
}

func bar() (int, string) {
	return 43, "Harry Potter"
}

func main() {
	// fmt.Println(foo())
	// fmt.Println(bar())
	x := foo()
	fmt.Println(x)

	y, z := bar()
	fmt.Println(y, z)

}
