package main

import "fmt" // Needed only if using fmt.Println()

func main() {
	// Built-in println() – no formatting, no import needed
	println("This is from built-in println (no formatting)")

	// fmt.Println() – requires import "fmt"
	fmt.Println("This is from fmt.Println (formatted and professional output)")
}

/*
 Benefit Summary:
Use fmt for any real application, logging, or formatted messages — it’s reliable, readable, and future-proof.
Use println() only for quick and dirty debug logs during early development or experimentation.
*/
