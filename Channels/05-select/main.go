package main

import "fmt"

// send distribution numbers into even, odd and quit channels.
func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i // send even numbers
		} else {
			o <- i // send odd numbers
		}
	}
	q <- 0 // send signal completion
}

// receive listens to the channels and processes values accordingly.
func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the even channel:", v)
		case v := <-o:
			fmt.Println("from the odd channel:", v)
		case v := <-q:
			fmt.Println("from the quit channel:", v)
			return // exit function when the quit signal is received
		}
	}
}

func main() {
	even := make(chan int) // channel for even numbers
	odd := make(chan int)  // channel for odd numbers
	quit := make(chan int) // channel for quit signal

	// send
	go send(even, odd, quit) // start send function in a Goroutine

	// receive
	receive(even, odd, quit) // process received values

	fmt.Println("about to exit")

}
