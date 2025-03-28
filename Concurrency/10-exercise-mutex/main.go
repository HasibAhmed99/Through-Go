package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	var mu sync.Mutex

	incrementor := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := incrementor
			// runtime.Gosched()
			v++
			incrementor = v
			fmt.Println(incrementor)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", incrementor)
}
