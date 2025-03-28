package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age   int
}

type ByName []Person

func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].First < bn[j].First }

func main() {
	p1 := Person{"Harry", 32}
	p2 := Person{"Bob", 36}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 30}

	people := []Person{p1, p2, p3, p4}
	fmt.Println(people)

	sort.Sort(ByName(people)) // ByName(people) temporarily treats people as ByName so it can use Len(), Swap(), and Less().
	fmt.Println(people)
}
