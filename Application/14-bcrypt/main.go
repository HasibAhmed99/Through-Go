package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pw := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
	if err != nil {
		fmt.Println("Error!", err)
	}
	fmt.Println(pw)
	fmt.Println(bs)

	loginPword := `password1234`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPword))
	if err != nil {
		fmt.Println("YOU CANT LOGIN")
		return
	}

	fmt.Println("YOU'RE LOGGED IN")
}
