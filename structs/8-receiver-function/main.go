package main

import "fmt"

type User struct {
	name string
	age  int
}

func printUserDetails(usr User) {
	fmt.Println(usr.name)
	fmt.Println(usr.age)
}

// receiver function
func (usr User) printUser() {
	fmt.Println(usr.name)
	fmt.Println(usr.age)
}

func (u User) call(s string) {
	fmt.Println(u.name)
	fmt.Println(u.age)
	fmt.Println(s) // was just wondering
}

func main() {
	user1 := User{
		name: "Rob",
		age:  30,
	}
	printUserDetails(user1)
	user1.printUser() // proper way and more readable
	user1.call("Robert")

	fmt.Println("-----------")

	user2 := User{
		name: "Smith",
		age:  26,
	}
	printUserDetails(user2)
	user2.printUser() // proper way and more readable
}
