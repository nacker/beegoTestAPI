package main

import "fmt"

/// cap是slice的最大容量，append函数添加元素，如果超过原始slice的容量，会重新分配底层数组。

func main() {
	s1 := make([]int, 3, 6)
	fmt.Println("s1= ", s1, len(s1), cap(s1))
	s2 := append(s1, 1, 2, 3)
	fmt.Println("s1= ", s1, len(s1), cap(s1))
	fmt.Println("s2= ", s2, len(s2), cap(s2))
	s3 := append(s2, 4, 5, 6)
	fmt.Println("s1= ", s1, len(s1), cap(s1))
	fmt.Println("s2= ", s2, len(s2), cap(s2))
	fmt.Println("s3= ", s3, len(s3), cap(s3))
}