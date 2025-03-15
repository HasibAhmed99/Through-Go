package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 10, 6, 19, 17, 99, 42, 72}
	xs := []string{"Harry", "M", "s", "Dr.Kavin", "Potter"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println("-----------------")

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
