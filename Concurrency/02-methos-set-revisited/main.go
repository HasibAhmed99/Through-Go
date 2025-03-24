package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// func area(c circle) float64 {
// 	return math.Pi * c.radius * c.radius
// }

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	// c := &circle{5}
	// info(c)

	// c := circle{5}
	// // info(c)
	// fmt.Println(area(c))

	c := circle{5}
	fmt.Println(c.area())
}
