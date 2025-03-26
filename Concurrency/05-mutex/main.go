package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Display the number of CPUs available to the program
	fmt.Println("CPUs", runtime.NumCPU())
	// Display the number of currently running goroutines
	fmt.Println("Goroutines", runtime.NumGoroutine())

	// counter is a shared integer variable that will modified the concurrently
	counter := 0

	// gs (goroutine size) represents the number of goroutines to be created
	const gs = 100

	// wg (WaitGroup) is used to wait for all goroutines to complete execution
	var wg sync.WaitGroup

	// Add 'gs' (100) to the WaitGroup counter to track all goroutines
	wg.Add(gs)

	// mu (Mutex) ensures mutual exclusion, preventing race condition when modifying 'counter'
	var mu sync.Mutex

	// Loop to launch 'gs' (100) goroutines
	for i := 0; i < gs; i++ {
		go func() {
			// Locking the mutex to prevent race conditions when accessing 'counter'
			mu.Lock()

			// Read the current  value of 'counter'
			v := counter

			// Yield execution to allow other goroutines to run(does not prevent race condition alone)
			runtime.Gosched()

			// Increment the local copy of 'counter'
			v++

			// Write the updated value back to 'counter'
			counter = v

			// Unlock the mutex so other goroutines can access 'counter'
			mu.Unlock()

			// Mark this goroutine as done in the WaitGroup
			wg.Done()
		}()

		// Display the current number of active goroutines
		fmt.Println("Goroutines", runtime.NumGoroutine())
	}

	// Wait for all the goroutines to finish before proceeding
	wg.Wait()

	// Display the final number of running goroutines (should be 1, the main goroutine)
	fmt.Println("Goroutines", runtime.NumGoroutine())

	// Print the final value of the counter after all increments are done
	fmt.Println("count", counter)
}
