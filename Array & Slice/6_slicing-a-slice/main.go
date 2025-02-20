package main

import "fmt"

func main() {
	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("---------")

	// [inclusive:exclusive]
	fmt.Printf("xi - %#v\n", xi[0:4])
	fmt.Println("---------")
	
	// [:excluesive]
	fmt.Printf("xi - %#v\n", xi[:7])
	fmt.Println("---------")
	
	//[inclusive:]
	fmt.Printf("xi - %#v\n", xi[4:])
	fmt.Println("---------")
	
	//[:]
	fmt.Printf("xi - %#v\n", xi[:])
	fmt.Println("---------")

	xii := []int{4, 5, 7, 8, 42}
	fmt.Println(xii[1])
	fmt.Println(xii)
	fmt.Println(xii[1:])
	fmt.Println(xii[1:3])

	for i, v := range xii {
		fmt.Printf("index: %v \tvalue: %v\n", i, v)
	}

	for i := 0; i <= 4; i++ {
		fmt.Println(i, xii[1])
	}
	
}