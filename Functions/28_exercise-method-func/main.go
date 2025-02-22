package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	// fmt.Println("Hello, I am", p.first)
	// fmt.Println("And my age is", p.age)
	fmt.Println("Hello, I am", p.first, "and my age is", p.age) // I can also do this way
}

func main() {
	p1 := person{
		first: "Harry",
		age:   24,
	}
	//fmt.Println(p1)
	p1.speak()

}
