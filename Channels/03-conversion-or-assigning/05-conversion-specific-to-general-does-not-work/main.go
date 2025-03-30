package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int)
	cs := make(chan<- int)

	fmt.Println("---------")
	fmt.Printf("%T/n", c)
	fmt.Printf("%T/n", cr)
	fmt.Printf("%T/n", cs)

	fmt.Println("---------")
	fmt.Printf("c\t%T\n", (chan int)(cr))
	fmt.Printf("c\t%T\n", (chan int)(cs))
}
