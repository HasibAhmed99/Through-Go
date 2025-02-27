package main

import "fmt"

func intDelta(n *int) {
	*n = 43
}

func sliceDelta(ii []int) {
	ii[0] = 99
}

func mapDelta(md map[string]int, s string) {
	md[s] = 30
}

func main() {
	a := 42
	fmt.Println(a)
	intDelta(&a)   // need pointer otherwise it's will just only copy the value
	fmt.Println(a) // if we don't use pointer the value wont change and it will remain like original value 42

	xi := []int{1, 2, 3, 4}
	fmt.Println(xi)
	sliceDelta(xi) // default don't need pointers
	fmt.Println(xi)

	m := make(map[string]int)
	m["Harry"] = 27
	//m["Tom"] = 29
	fmt.Println(m["Harry"])
	//mapDelta(m, "Tom")
	mapDelta(m, "Harry") // default don't need pointers
	fmt.Println(m["Harry"])
	//fmt.Println(m["Tom"])
}
