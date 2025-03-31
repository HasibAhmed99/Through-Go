package main

import "fmt"

func main() {
	c := make(chan int) // unbuffered channel

	// send

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	// receive

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}
