package main

import "fmt"

func main() {
	 xi := make(map[string]int)
	 xi["Henry"] = 32
	 xi["Mike"] = 33 
	 xi["Smith"] = 34
	fmt.Println(xi)
	fmt.Printf("%#v\n", xi)
	fmt.Println(len(xi))

	// v, ok := xi["Mike"]
	// if ok {
	// 	fmt.Println("the value prints", v)
	// } else {
	// 	fmt.Println("The key doesnt exist")
	// }
     
	// statement-statement idiom 
	//if v, alright := xi["Alex"]; !alright {
	if v, ok := xi["Alex"]; !ok {
     fmt.Println("key doesnt exist")
	} else {
	fmt.Println("The value prints", v)
	}
}
	