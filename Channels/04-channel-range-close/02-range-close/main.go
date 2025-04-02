package main

import "fmt"

func main() {
	// Creating an unbuffered channel of type int
	// This channel will be used to send or receive integer values
	c := make(chan int) // unbuffered channel

	// send
	// Starting a goroutine to send values into channel
	go func() {
		for i := 0; i < 5; i++ {
			c <- i // sending data to the channel
		}
		close(c) // Closing the channel sending all the values
	}()

	// receive
	// Receiving value from the channel
	// Using range automatically stops the channel when the channel is closed
	for v := range c {
		fmt.Println(v) // Printing received values
	}

	fmt.Println("about to exit") // Indicating program termination
}
