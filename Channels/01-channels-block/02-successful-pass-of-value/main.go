package main

import "fmt"

// This will be run because the channel sending
func main() {
	ch := make(chan int)

	go func() {
		ch <- 42
	}()

	fmt.Println(<-ch)
}
