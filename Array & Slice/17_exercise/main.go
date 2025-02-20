package main

import "fmt"

func main() {
	x1 := make([]int, 50)
	fmt.Println(x1)
	fmt.Println(len(x1))
	fmt.Println(cap(x1))
	
	x2 := make([]int, 0, 50)
	fmt.Println(x2)
	fmt.Println(len(x2))
	fmt.Println(cap(x2))

	fmt.Println("----------------------")
	// x1 = append(x1, 98)
	x1[0] = 98
	x2 = append(x2, 99)
	
	fmt.Println(x1)
	fmt.Println(len(x1))
	fmt.Println(cap(x1))
	
	fmt.Println(x2)
	fmt.Println(len(x2))
	fmt.Println(cap(x2))
	
	fmt.Println("---------------------")
	
	xs := make([]string, 0, 50)
	//fmt.Println(xs)
	fmt.Println(len(xs))
	fmt.Println(cap(xs))
	
	fmt.Println("---------------------")

	xs = append(xs,  ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	//fmt.Println(xs)
	fmt.Println(len(xs)) 
	fmt.Println(cap(xs))

  //without using range clause
	for i := 0; i < len(xs); i++ {
		fmt.Println(xs[i], i)
	}

}