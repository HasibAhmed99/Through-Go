package main

import "fmt"

func main() {
	
	var d [5]int
	fmt.Println(d)
	d[3] = 42
	fmt.Println(d)
	fmt.Println(len(d))
     
	fmt.Println("---------")
	a := [3]int{41, 42, 43}
	fmt.Println(a)	
	fmt.Println("---------")
	
	b := [...]string{"Hello", "World"}
	fmt.Println(b)
	fmt.Println("---------")
	
	var c [2]int
	c[0] = 7
	c[1] = 10
	fmt.Println(c)
	
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", c)
	fmt.Println("---------")
	
   var x [7]int
   x[0] = 42
   fmt.Printf("%#v \t\t\t %T\n",x, x)

   y := [4]int{50, 51, 52, 53}
   fmt.Printf("%#v \t\t\t\t %T\n", y, y)

   z := [...]string{"Banana", "Cherry", "Apple"}
   fmt.Printf("%#v \t\t %T\n", z, z)
}