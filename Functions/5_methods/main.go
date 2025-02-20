package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Println("I am", p.first)
	fmt.Println("My age is", p.age)
}

func main() {
	p1 := person{
		first: "Morty",
		age:   37,
	}

	p2 := person{
		first: "Kevin",
		age:   34,
	}
	p1.speak()
	p2.speak()

}
