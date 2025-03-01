package main

import "fmt"

type speaker interface {
	speak()
}

type dog struct {
	name string
}

func (d dog) speak() {
	fmt.Println(d.name, "says Woof!")
}

type cat struct {
	name string
}

func (c cat) speak() {
	fmt.Println(c.name, "says Meow!")
}

func main() {
	var x speaker
	x = dog{"Buddy"}
	//fmt.Println(x)

	x.speak()

	x = cat{"Luna"}
	x.speak()
}
