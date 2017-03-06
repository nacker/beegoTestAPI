package main

import (
	"fmt"
)

type Person struct {
	Id   int
	Name string
}

func (this Person) test(x int) {
	fmt.Println("Id:", this.Id, "Name", this.Name)
	fmt.Println("x=", x)
}

func main()  {
	p := Person{2, "张三"}

	p.test(1)
	var f1 func(int)  = p.test

	f1(2)
	Person.test(p, 3)
	var f2 func(Person, int) = Person.test
	f2(p, 4)
}


