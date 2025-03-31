package main

import "fmt"

// send channel
func send(cs chan<- int) {
	cs <- 42
}

// receive channel
func receive(cr <-chan int) {
	fmt.Println(<-cr)
}

func main() {
	c := make(chan int)

	go send(c)

	receive(c)

	close(c) // We closed the channel

	fmt.Println("about to exit")
}
