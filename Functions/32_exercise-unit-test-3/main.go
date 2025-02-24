package main

import "fmt"

func landscape(loc string) string {
	return fmt.Sprint("My idea of landscape is ", loc)
}

func main() {
	fmt.Println(landscape("Maldives"))
}
