package main

import "fmt"

// Define type constraint for numeric types
type myNums interface {
	int | int8 | int16 | int32 | int64 |
		float32 | float64
}

// Define a generic struct
type Number[T myNums] struct {
	value T
}

// Method to add a number to the stored value
func (n *Number[T]) Add(val T) {
	n.value += val
}

// Method to get the stored value
func (n Number[T]) GetValue() T {
	return n.value
}

// func (n *Number[T]) Add(val T) T {
// 	n.value += val
// 	return n.value
// }

func main() {
	// Creating integer number struct
	intNum := Number[int]{value: 10}
	intNum.Add(5)
	fmt.Println("Integer value:", intNum.GetValue())

	// Creating float number struct
	floatNum := Number[float64]{value: 2.55}
	floatNum.Add(2.45)
	fmt.Println("Float value:", floatNum.GetValue())
}
