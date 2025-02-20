package main

import "fmt"

func main() {
	xp := [][]string{
	{"Rob", "Mike", "Hank"},
	{"Bob", "Alice", "Jenny"},
	{"Grace", "Eve", "David"},
	}
	fmt.Println(xp)
	
	fmt.Println("-------------")
	xp = append(xp, []string{"Frank", "Bruno"})
	fmt.Println(xp)
	
	
	fmt.Println("-------------")
	xp[1][1] = "Brook"
	xp[2][0] = "Franky"
	fmt.Println(xp)
	
	fmt.Println("-------------")
	for i, v := range xp {
		//fmt.Println(xp[2][1], i)
		fmt.Printf("%v -> %v\n", i, v)
		for a, b := range v{
			fmt.Printf("%v - %v\n", a, b)
			}
	}
	
	fmt.Println("-------------")
	for i := 0; i < 4; i++{
		fmt.Println(xp[i], i)
	}
	 
}
