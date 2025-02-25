package main

import "fmt"

func something() func() int {
	return func() int {
		return 42
	}
}

func main() {
	st := something()
	fmt.Println(st())

}
