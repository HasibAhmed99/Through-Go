package main

import "fmt"

type person struct {
	first string
}

func nameChange(p person, s string) person {
	p.first = s
	return p
}

// func nameChangePtr(p *person, s string) person {
// 	p.first = s
// 	return *p
// }

func nameChangePtr(p *person, s string) {
	p.first = s
}

func main() {
	p := person{"Harry"}
	fmt.Println(p)

	p = nameChange(p, "Bob")
	fmt.Println(p)

	nameChangePtr(&p, "Rob")
	fmt.Println(p)
}
