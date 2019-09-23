package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main()  {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
			fmt.Println(err)
	}  // END if
	fmt.Println(s)
	fmt.Println(bs)

	loginPword1 := `password123`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPword1))
	if err != nil {
		fmt.Println("You cannot login")
		return
	}  // END if
	fmt.Println("You're logged in now")

}  // END main
