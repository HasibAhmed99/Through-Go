package main

import "fmt"

func main() {
	x := "Harry"
	fmt.Println(x)
	fmt.Println(&x)

	y := &x
	fmt.Printf("%v\t%T\n", y, y)

	*y = "Not Harry"
	fmt.Println(x)
	fmt.Println(*y)
}
