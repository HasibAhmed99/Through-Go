package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walking")
}

func (d dog) run() {
	fmt.Println("My name is", d.first, "and I'm running")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.walk()
	y.run()
}

func main() {
	d1 := dog{"Copper"} // create a new dog instance with a name "Copper"
	d1.walk()           // calls the walk() method
	d1.run()            // calls the run () method
	youngRun(d1)        // calls youngin(), which calls walk() and run()

	d2 := dog{"Teddy"}          // create a new dog instance with a name "Teddy"
	youngRun(d2)                // calls youngin() with "Teddy"
	d2 = d2.changeName("Buddy") // change the name from to "Teddy" to "Buddy"
	youngRun(d2)                // calls youngin with "Buddy"
}

func (d dog) changeName(s string) dog {
	d.first = s
	return d
}
