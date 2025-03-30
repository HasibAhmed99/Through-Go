package main

import "fmt"

// send-only
func foo(cs chan<- int) {
	cs <- 42
}

// receive-only
func bar(cr <-chan int) {
	fmt.Println(<-cr)
}

func main() {
	c := make(chan int)

	// send-only
	go foo(c)

	// receive-only
	bar(c)

	fmt.Println("about to exit")
}
