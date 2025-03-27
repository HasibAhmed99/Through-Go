package main

import "fmt"

type person struct {
	first string
}

func (p *person) speak() {
	fmt.Println("nothing to say")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "James",
	}

	saySomething(&p1)

	// does not work because it will only take pointer
	// saySomething(p1)

	p1.speak()
}
