package main

import (
	"fmt"
	"log"
)

func main() {
	// wrapping divide function
	divideWithLog := logWrapper(divide)

	// calling the wrapped function
	result := divideWithLog(10, 2)
	fmt.Println(result)
}

// original function
func divide(a, b float64) float64 {
	return a / b
}

// wrapper function

func logWrapper(f func(float64, float64) float64) func(float64, float64) float64 {
	return func(a, b float64) float64 {
		log.Println("Function is called with values: ", a, b)
		return f(a, b)
	}
}
