package main

import "fmt"

/// 指向同一底层数组的slice之间copy时，允许存在重叠。copy数组时，受限于src和dst数组的长度最小值。

func main()  {
	s1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := make([]int, 3, 20)
	var n int
	n = copy(s2, s1)
	fmt.Println(n, s2, len(s2), cap(s2))

	s3 := s1[4:6]
	fmt.Println(n, s3, len(s3), cap(s3))

	n = copy(s3, s1[1:5])
	fmt.Println(n, s3, len(s3), cap(s3))
}
