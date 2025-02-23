package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2) //  calculates x raised to the power of y and here 4 is exponent and computes
}

func info(s shape) float64 {
	return s.area()
}
func main() {
	s1 := square{
		length: 5,
		width:  8,
	}

	c1 := circle{
		radius: 4,
	}
	fmt.Println(info(s1))
	fmt.Println(info(c1))
}
