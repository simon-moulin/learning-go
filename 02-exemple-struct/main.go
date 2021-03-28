package main

import "fmt"

type User struct {
	Name  string
	Email string
}

type Admin struct {
	User
	Level int
}

func main() {
	u := User{
		Name:  "Bob",
		Email: "bob@gmail.com",
	}
	fmt.Printf("User= %v\n", u)

	a := Admin{
		Level: 2,
		User: User{ // Obligé de passe pas cette notation
			Name:  "Alice",
			Email: "alive@golang.org",
		},
	}

	fmt.Printf("User= %v\n", a)
	fmt.Printf("Admin nom=%v level=%v", a.Name, a.Level)

}
