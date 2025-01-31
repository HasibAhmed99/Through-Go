package main

import "fmt"

func main() {
	 xi := make(map[string]int)
	 xi["Henry"] = 32
	 xi["Mike"] = 33 
	 xi["Smith"] = 34
	fmt.Println(xi)
	fmt.Printf("%#v\n", xi)
	fmt.Println(len(xi))
	
	fmt.Println("-------------------")
	delete(xi, "Mike")
	fmt.Println(xi)
	fmt.Printf("%#v\n", xi)
	fmt.Println(len(xi))

	fmt.Println("--- accessing keys that dont exist")
	delete(xi, "Cooper") // wont panic
	fmt.Println(xi["Nike"]) // wont panic 
}
