package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// main function demonstrate concurrent programming in Go
// by launching multiple goroutines that safely increment a shared counter
// using atomic operations to prevent race condition
func main() {
	// Print the number of CPU cores available for execution.
	fmt.Println("CPUs", runtime.NumCPU())
	// Print the current number of active goroutines (should be 1 initially).
	fmt.Println("Goroutines", runtime.NumGoroutine())

	// counter is a shared variable that will be incremented by multiple goroutines concurrently.
	var counter int64

	// gs (goroutine size) defines the number of goroutines to be launched.
	const gs = 100
	// wg (WaitGroup) ensures that all goroutines complete execution before the main exits.
	var wg sync.WaitGroup
	// Set the WaitGroup counter to match the number of goroutines.
	wg.Add(gs)

	// Launce 'gs' (100) goroutines.
	for i := 0; i < gs; i++ {
		go func() {
			// Automatically increment the shared counter by 1 to avoid race condition
			atomic.AddInt64(&counter, 1)
			// Yield execution to allow other goroutines to run.
			runtime.Gosched()
			// Print the current value of the counter using an atomic read operation.
			fmt.Println("counter\t", atomic.LoadInt64(&counter))
			// Mark this goroutine as done in the WaitGroup.
			wg.Done()
		}()
		//Print the current number of active goroutines
		fmt.Println("Goroutines", runtime.NumGoroutine())
	}
	// Wait for all the goroutines to finish before proceeding
	wg.Wait()
	// Print the current number of active goroutines (should be 1, as only main remains).
	fmt.Println("Goroutines", runtime.NumGoroutine())
	// Print the final value of the counter after all increments are done.
	fmt.Println("count:", counter)
}
