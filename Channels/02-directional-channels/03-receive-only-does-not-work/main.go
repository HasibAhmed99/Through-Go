package main

import "fmt"

func main() {
	ch := make(<-chan int, 2) // receive-only

	ch <- 42 // Won't work because the channels are Receive-only. So it can't be send.[And it will give compile-time error]
	ch <- 43 // Won't work because the channels are Receive-only. So it can't be send.[And it will give compile-time error]

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println("-------")
	fmt.Println("%T\n", ch)
}
