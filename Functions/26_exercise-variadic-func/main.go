package main

import "fmt"

func foo(ii ...int) int {
	fmt.Printf("%T \n", ii)
	t := 0
	for _, v := range ii {
		t += v
	}
	return t
}

func bar(ii []int) int {
	fmt.Printf("%T \n", ii)
	t := 0
	for _, v := range ii {
		t += v
	}
	return t
}

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(foo(x...)) // Unfurls the slice into individual arguments for foo
	fmt.Println(bar(x))    // Passes the slice directly to bar

}
