package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range xi {
		fmt.Printf("%v \t %T \t %v\n", i, v, v)
	}
	fmt.Printf("%T \t %#v\n", xi, xi)
	fmt.Printf("%T \t %v\n", xi, xi)

	//[inclusive:exclusive]
 	fmt.Println(xi[:5])
 	fmt.Println(xi[5:])
	fmt.Println(xi[2:7])
	fmt.Println(xi[1:6])

	}