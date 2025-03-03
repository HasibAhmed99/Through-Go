package main

import "fmt"

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

func addT[T int | float64](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(addI(5, 5))
	fmt.Println(addF(2.63, 2.37))

	fmt.Println(addT(5, 5))
	fmt.Println(addT(2.63, 2.37))
}
