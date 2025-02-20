package main

import "fmt"

func main(){
    foo()

	x := func (){
		fmt.Println("This is an anonymous func")
	}

	y := func (s string) {
		fmt.Println("This is an anonymous showing my name", s)
 	}

	x() // func expression
	y("Rick") // func expression

}
func foo() {
 fmt.Println("This is foo")
}