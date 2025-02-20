package main

import "fmt"

func main() {
	m := map[string]int{
		"Rob": 28,
		"Bob": 29,
	}
     // also can do like this
	// m := make(map[string]int)
	// m["Rob"] = 28
	// m["Bob"] = 29
	fmt.Println(m)
	
	for key := range m{
		fmt.Printf("just the keys: %s\n", key)
	}

	for key, value := range m{
		fmt.Printf("%s is %d years old\n", key, value)
	}

	for _, value := range m{
		fmt.Printf("just the values: %d\n", value)
	}

	// delete
	m["Newton"] = 1643
	fmt.Printf("%#v\n", m)
	delete(m, "Newton")
	fmt.Printf("%#v\n", m)
	delete(m, "Newton") // no panic

	fmt.Println("len: ", len(m))

	// -----------or----------
	/* var mi map[string]int
	fmt.Println(mi["Felix"])
	mi = make(map[string]int)
	mi["Zoro"] = 21
	mi["Hawk"] = 43
	fmt.Println(mi) */

} 