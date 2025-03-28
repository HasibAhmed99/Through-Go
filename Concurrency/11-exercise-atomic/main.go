package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup

	var incrementor int64
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementor, 1)
			//atomic.LoadInt64(&incrementor)
			r := atomic.LoadInt64(&incrementor)
			fmt.Println(r)
			//fmt.Println(incrementor)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", incrementor)
}
