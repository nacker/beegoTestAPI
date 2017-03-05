package main

import "fmt"

type User struct {
	Id int
	Name string
}

func main()  {
	a := [...]User{
		{0, "User0"},
		{8, "User8"},
	}

	b := [...]*User{
		{0, "User0"},
		{8, "User8"},
	}

	fmt.Println(a, len(a))
	fmt.Println(b, len(b))
}
