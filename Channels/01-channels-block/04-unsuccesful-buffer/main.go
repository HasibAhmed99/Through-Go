package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 42
	ch <- 43 // blocks because we took only one space

	fmt.Println(<-ch)
}
