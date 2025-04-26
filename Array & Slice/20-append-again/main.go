package main

import "fmt"

func main() {
	var x []int

	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)
	y := x

	x = append(x, 4)
	y = append(y, 5)

	x[0] = 10

	fmt.Println(x) // output: [10, 2, 3, 5] . The output of x wont be [10, 2, 3, 4]. Because 4 was over written by 5 in the RAM.
	fmt.Println(y) // output: [10, 2, 3, 5]
}
