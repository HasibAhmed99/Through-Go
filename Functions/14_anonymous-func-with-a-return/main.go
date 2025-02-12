package main

import "fmt"

func main() {
	foo()
	
	func() {
		fmt.Println("This is an anonymous func with no parameters")
}()

	func (s string) {
		fmt.Println(s)
	}("Hello World")

 // Function with a return
  s := func(s string) string {
	fmt.Println("This function has a return")
	return fmt.Sprint("This book is ", s)
  }("The Great Gatsby")
  fmt.Println(s)

}

func foo() {
	fmt.Println("Hey I am foo")
}