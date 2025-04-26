package main

import "fmt"

func main() {
	xi := []int{42, 43, 44}
	fmt.Println(xi)
	fmt.Println("----------")

	// variadic parameter
	xi = append(xi, 45, 46, 47, 48, 99, 888)
	fmt.Println(xi)
	fmt.Println("----------")

	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	x = append(x, 11, 111, 22, 222, 1122)
	fmt.Println(x)
	fmt.Println("----------")

	y := []int{123, 345, 456, 678, 890}
	z := []int{111, 256}
	z = append(z, x...)
	z = append(z, y...)
	fmt.Println(z)
	x = append(x, y...) // I can add y with x
	fmt.Println(x)
}
