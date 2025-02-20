package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(len(x))
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])
	// fmt.Println(x[5]) // this wont work because This index out of range which causes a runtime error
	fmt.Println("------------")
	
	xs := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)"}
	fmt.Println(xs)
	fmt.Printf("%T\n", xs)
	
	// blank identifer to not use a returned value
	for _, value := range xs {
		fmt.Printf("value: %v\n", value)
	}
	fmt.Println("------------")
	fmt.Println(xs[0])
	fmt.Println(xs[1])
	fmt.Println(xs[2])
	fmt.Println(xs[3])
	
	
	for i := 0; i < len(xs); i++ {
		fmt.Println(xs[i])
	}
}