package main

import "fmt"

func main() {

	func(a, b int) int {
		result := a / b
		fmt.Println("The equation is", result)
		return result
	}(9, 3)

	func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()
}
