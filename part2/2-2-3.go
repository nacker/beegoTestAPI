package main

import "fmt"

func main()  {
	s1 := make([]int, 0)
	s1 = test(s1)
	fmt.Println(s1)
}

func test(s []int) []int {
	s = append(s, 3)
	return s
}