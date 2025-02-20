package main

import "fmt"

func main() {
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
    fmt.Println("---------")
	x[0] = 42
	x[9] = 999
	
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println("---------")
	
	x = append(x, 2025)
	
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println("---------")
	
	x = append(x, 2026)
	
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println("---------")
	
	x = append(x, 2027)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}