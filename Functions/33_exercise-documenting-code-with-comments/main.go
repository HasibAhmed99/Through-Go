package main

import "fmt"

// Landscape prints out the user's input of landscape as a sentence
func Landscape(loc string) string {
	return fmt.Sprint("My idea of landscape is ", loc)
}

func main() {
	fmt.Println(Landscape("Maldives"))
}
