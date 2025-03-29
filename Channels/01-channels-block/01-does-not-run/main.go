package main

import "fmt"

// This function wont run because channel isn't sending and receiving at the same time and It will be block. Think it like a relay race.[This apply only buffered channels]
func main() {
	ch := make(chan int)

	ch <- 42 // blocks

	fmt.Println(<-ch)
}
