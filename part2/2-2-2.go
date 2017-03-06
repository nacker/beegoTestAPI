package main

import "fmt"

func main()  {
	s1 := make([]int, 0)
	test(s1)
	fmt.Println(s1)
}

func test(s []int) {
	s = append(s, 3)
	//因为原来分配的空间不够，所以在另外一个地址又重新分配了空间，所以原始地址的数据没有变
}