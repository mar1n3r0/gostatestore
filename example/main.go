package main

import (
	"fmt"
	"time"

	store "github.com/mar1n3r0/gostatestore"
)

type User struct {
	Name     string
	Username string
}

func main() {
	user := User{}
	fmt.Println("User.Name before: " + user.Name)
	fmt.Println("User.Username before: " + user.Username)
	store.NewStore()
	store.Reader(&user)
	user2 := User{
		Name:     "test2",
		Username: "tester2",
	}
	store.Writer(&user2)
	user3 := User{
		Name:     "test3",
		Username: "tester3",
	}
	store.Writer(&user3)

	time.Sleep(time.Second * 1)
	fmt.Println("User.Name after: " + user.Name)
	fmt.Println("User.Username after: " + user.Username)
}
