package main

import "fmt"

func main() {
	x := doMath(42, 12, add)
	fmt.Println(x)

	y := doMath(42, 12, subtract)
	fmt.Println(y)

	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", subtract)
	fmt.Printf("%T\n", doMath)
}
func add(a, b int) int {
	return a + b
}
func subtract(a, b int) int {
	return a - b
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)

}