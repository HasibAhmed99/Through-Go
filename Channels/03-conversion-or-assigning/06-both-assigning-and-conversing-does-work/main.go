package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int)
	cs := make(chan<- int)

	fmt.Println("---------")
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

	// assigning general to specific
	cr = c
	cs = c

	// converting specific to general [but it's need explicit]
	fmt.Printf("c\t%T\n", (chan int)(cr))
	fmt.Printf("c\t%T\n", (chan int)(cs))
}
