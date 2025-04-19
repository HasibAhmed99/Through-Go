package main

import "fmt"

func main() {
	rp := []string{"Rob", "Pike", "Water", "Vanila"}
	kt := []string{"ken", "Thompson", "Water", "Chocolate"}
	fmt.Println(rp)
	fmt.Println(kt)

	mds := [][]string{rp, kt}
	fmt.Println(mds)
	fmt.Println("-------------")

	bm := []string{"Bruno", "Mars", "Strawberry", "Water"}
	fmt.Println(bm)
	lg := []string{"Lady", "Gaga", "Butterscotch", "Water"}
	fmt.Println(lg)

	mds2 := [][]string{bm, lg}
	fmt.Println(mds2)

}
