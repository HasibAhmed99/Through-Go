package main

import "fmt"

func main() {
	ch := make(chan<- int, 2) // Send-only channel

	ch <- 42
	ch <- 43

	fmt.Println(<-ch) // Won't work because the channels are send-only. So it can't be receive.[And it will give compile error]
	fmt.Println(<-ch) // Won't work because the channels are send-only. So it can't be receive.[And it will give compile error]
	fmt.Println("-------")
	fmt.Println("%T\n", ch)
}
