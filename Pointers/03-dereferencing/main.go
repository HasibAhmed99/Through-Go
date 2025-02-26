package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)

	y := &x // assigning var x memory address into var y // we can also say it this way that y stores x's memory address
	fmt.Printf("%v\t%T\n", y, y)
	fmt.Println(*y)
	fmt.Println(*&x) // simplified : It's the same as *y

	*y = 43
	fmt.Println(x) // It will print out 43 because *y changes the value of variable of x // its not changes temporary but permanently
	fmt.Println(*y)

}

/* func main() {
	x := 15
	fmt.Println(x)
	fmt.Println(&x)

	y := &x
	fmt.Println(y)
	fmt.Printf("%v\t%T\n", y, y)

	*y = 18
	fmt.Println(x)
	fmt.Println(*y)
} */
