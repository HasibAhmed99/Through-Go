package main

import "fmt"

// the reason why I didn't use short declaration here because [:=] short declaration don't work at package level.
var x = func() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func main() {

	x() // func expression

	y := func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}
	y() // func expression

}
