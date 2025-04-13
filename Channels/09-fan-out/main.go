package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Create input and output channels
	c1 := make(chan int)
	c2 := make(chan int)

	// Start producer to fill c1 with numbers
	go populate(c1)

	// Start fan-out/in processor to handle for parallel work
	go fanOutIn(c1, c2)

	// Read and Print the results from output channel c2
	for v := range c2 {
		fmt.Println(v)
	}
	fmt.Println("about to exit")
}

// populate sends integer 0-99 into the channel and then closes it
func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c) // Signals no more data will be sent
}

// fanOutIn reads from c1, processes data concurrently, send results into c2
func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup

	// Launce a goroutine for each input value
	for v := range c1 {
		wg.Add(1)
		go func(v2 int) {
			c2 <- timeConsumingWork(v2) // send processed result to c2
			wg.Done()
		}(v)
	}
	// Wait for all processing goroutines to finish
	wg.Wait()
	// Close output channel after all the results are sent
	close(c2)
}

// timeConsumingWork simulates a random number delay and returns a modified value
func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500))) // Simulated work
	return n + rand.Intn(1000)                                   // return transformed result
}
