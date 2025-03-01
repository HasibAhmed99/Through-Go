package main

import (
	"fmt"
	"time"
)

func doWork() {
	for i := 0; i < 100000; i++ {
		fmt.Println(i)
	}
}

func timeFunc(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func main() {
	timeFunc(doWork)
}
