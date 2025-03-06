package main

import "fmt"

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

type myNums interface {
	~int | ~float64
}

func addT[T myNums](a, b T) T {
	return a + b
}

type myalias int

func main() {
	var n myalias = 42

	fmt.Println(addI(5, 5))
	fmt.Println(addF(3.32, 1.68))

	fmt.Println(addT(n, 8))
	fmt.Println(addT(4.50, 5.50))
}
