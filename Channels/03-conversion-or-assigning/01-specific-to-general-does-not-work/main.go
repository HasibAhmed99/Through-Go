package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int) // receive-only
	cs := make(chan<- int) // sent-only

	fmt.Println("----------")
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

	// specific to general doest assign
	c = cr
	c = cs
}
