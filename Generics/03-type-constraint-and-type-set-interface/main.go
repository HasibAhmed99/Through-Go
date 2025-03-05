package main

import "fmt"

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

type myNums interface {
	int | float64
}

func addT[T myNums](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(addI(5, 5))
	fmt.Println(addF(3.67, 1.33))

	fmt.Println(addT(5, 5))
	fmt.Println(addT(3.67, 1.33))
}
