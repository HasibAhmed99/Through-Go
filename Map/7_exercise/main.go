package main

import "fmt"

func main() {
	xi := make(map[string][]string) 

	xi[`mike_pal`] =  []string{`shaken, not stirred`, `martnis`, `fast cars`}
	xi[`ron_steves`] = []string{`intelligence`, `literature`, `computer science`}
	xi[`keanu_reeves`] = []string{`dog`, `car`, `sunsets`}
	
	fmt.Printf("%#v\n", xi)

	for k, v :=  range xi {
		fmt.Println(k, v)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}

}