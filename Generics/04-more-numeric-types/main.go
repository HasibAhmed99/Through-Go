package main

import "fmt"

type myNums interface {
	int | int8 | int16 | int32 | int64 |
		float32 | float64
}

func addT[T myNums](a, b T) T {
	return a + b
}

func main() {
	// Testing with different numeric types
	fmt.Println(addT(5, 5))                         // int
	fmt.Println(addT(int8(10), int8(20)))           // int8
	fmt.Println(addT(int16(100), int16(200)))       // int16
	fmt.Println(addT(int32(250), int32(300)))       // int32
	fmt.Println(addT(int64(1050), int64(950)))      // int64
	fmt.Println(addT(float32(3.3), float32(1.7)))   // float32
	fmt.Println(addT(float64(5.33), float64(4.67))) // float64

}
