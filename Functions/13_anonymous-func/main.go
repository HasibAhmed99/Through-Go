package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("anonymous func ran")
	}()

	func (s string) {
		fmt.Println("This is an anonymous func showing my name", s)
	}("Rick")
}

func foo(){
	fmt.Println("Foo ran")
}