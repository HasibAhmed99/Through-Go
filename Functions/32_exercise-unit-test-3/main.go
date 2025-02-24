package main

import "fmt"

// landscape returns a string
func landscape(loc string) string {
	return fmt.Sprint("My idea of landscape is ", loc)
}

func main() {
	fmt.Println(landscape("Maldives"))
}
