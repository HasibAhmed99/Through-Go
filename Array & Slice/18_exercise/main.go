package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	jm := []string{"Miss", "Moneypenny", "I'm 008."}

	xp := [][]string{jb, jm}

	for i, v := range xp {
		fmt.Println(i, v)
		for a, b := range v{
			fmt.Println(a, b)
		}
	}
}