package main

import (
	"fmt"
)

type canBeAnything struct {
	first string
	last  string
	age   int
}
func main() {
  n1 := canBeAnything{
	first : "Steve",
	last : "jobs",
	age : 56,
  }

  n2 := canBeAnything{
	first : "Bill",
	last : "Gates",
	age : 69,
  }

  fmt.Println(n1)
  fmt.Println(n2)
  fmt.Printf("%T \t %#v\n", n1, n1)
  fmt.Printf("%T \t %#v\n", n2, n2)

  fmt.Println(n1.first, n1.last, n1.age)
  fmt.Println(n2.first, n2.last, n2.age)
}

		
