package main

import "fmt"

func main() {
	ab := map[string]int{
		"Henry": 23,
		"Alex":  25,
	}
	fmt.Println("Alex will turn", ab["Alex"], "this year")
	fmt.Println(ab)
	fmt.Printf("%#v\n", ab)

	fmt.Println("----------------")
	ef := make(map[string]int)
	ef["Gibson"] = 19
	ef["Ash"] = 20
	fmt.Println(ef)  
	fmt.Printf("%#v\n", ef)
	fmt.Println(len(ef))  

	
}
