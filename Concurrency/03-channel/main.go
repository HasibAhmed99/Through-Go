package main

import "fmt"

func doSometing(x int) int {
	return x * 5
}

func main() {
	ch := make(chan int)
	go func() {
		ch <- doSometing(5)
	}()
	fmt.Println(<-ch)
}
