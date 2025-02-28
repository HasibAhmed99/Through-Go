package main

import "fmt"

// Dog represents a dog that stores a retrievable data
type dog struct {
	first string
}

// Value receiver
func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walking.")
}

// Pointer receiver
func (d *dog) run() {
	d.first = "Copper"
	fmt.Println("My name is", d.first, "and I'm running")
}

func main() {
	// d1 := dog{
	// 	first: "Teddy",
	// }
	d1 := dog{"Teddy"} // a dog instance

	d1.walk() // I have to put the value receiver first to get walk method properly
	d1.run()  // If put this before walk method then it will print "copper" in both func expression

	d2 := &dog{"Buddy"} // a dog instance
	d2.walk()
	d2.run()
}
