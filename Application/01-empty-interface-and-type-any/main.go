package main

import "fmt"

func main() {
	var anything interface{} // Empty interface(interface{}) can hold any type [Old].

	anything = "Hello"
	fmt.Println(anything)

	anything = 42
	fmt.Println(anything)

	anything = 42.55
	fmt.Println(anything)

	anything = true
	fmt.Println(anything)

	var anything2 any // This can also store any value [New].

	anything2 = "Hello"
	fmt.Println(anything2)

	anything2 = 42
	fmt.Println(anything2)

	anything2 = 42.55
	fmt.Println(anything2)

	anything2 = true
	fmt.Println(anything2)
}
