package main

import "fmt"

func main() {
	ab := map[string]int{
		"Brook": 45,
		"Rob":   15,
		"Bob":   18,
	}
	fmt.Println("The age of Brook is", ab["Brook"])

	fmt.Println(ab)
	fmt.Printf("%#v\n", ab)

    fmt.Println("-------------")
	cd := make(map[string]int)
	cd["Franky"] = 27
	cd["Alice"]  = 31
	cd["Mars"]   = 42
	fmt.Println(cd) 
	fmt.Printf("%#v\n", cd)
	fmt.Println(len(cd)) 
	
	//for range over a map
	for key, value := range cd{
		fmt.Println(key, value)
	}
	
    // only value print
	for _, value := range cd{  
		fmt.Println(value)
	}
	
	// if take one variable it will print out only the key
	for key := range cd{
		fmt.Println(key)
	}
	
	fmt.Println("-------------")
	//for range with a SLICE
	xi := []int{42, 43, 44}

	for i, v := range xi{
		fmt.Println(i, v)
	} 

	for _, v := range xi{
		fmt.Println(v)
	}

	//It will print out only index
	for i := range xi{
		fmt.Println(i)
	}
	}
