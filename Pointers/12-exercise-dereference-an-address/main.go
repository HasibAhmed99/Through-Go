package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Actions speak louder than words"
	q := "Practice makes perfect"
	r := "Knowledge is power"
	n := 42

	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	fmt.Printf("%v\t%T\n", a, a)
	fmt.Printf("%v\t%T\n", b, b)
	fmt.Printf("%v\t%T\n", c, c)
	fmt.Printf("%v\t%T\n", d, d)

	fmt.Println(*a)
	fmt.Println(*b)
	fmt.Println(*c)
	fmt.Println(*d)

}
