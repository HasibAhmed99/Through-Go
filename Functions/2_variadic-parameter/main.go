package main

import "fmt"

func main() {
 x := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
 fmt.Println("The sum is", x)
}
func sum(xi ...int) int {
	fmt.Println(xi)
	fmt.Printf("%T\n", xi)

	n := 0
	for _, v := range xi{
		n += v
	}
	return n 
}