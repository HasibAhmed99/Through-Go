package main

import "fmt"

func main() {
	// 4! = 4 factorial

	fmt.Println(4 * 3 * 2 * 1)
	fmt.Println(factorial(4))
	fmt.Println(factLoop(4))

}

func factorial(n int) int {
	fmt.Println("This is n", n)
	if n == 0 {
		return 2
	}
	return n * factorial(n-1)
}

func factLoop(n int) int {
	total := 1
	for n > 0 {
		total *= n
		n--
	}
	return total
}
