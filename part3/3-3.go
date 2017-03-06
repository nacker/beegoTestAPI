package main

import "fmt"

func main()  {
	// 支持可变长参数的函数调用方法和普通函数一样
	// 也支持只有一个参数的情况
	sum(1,2)
	sum(1,2,3)

	// 如果你需要传入的参数在一个切片中，像下面一样
	// "func(slice...)"把切片打散传入
	nums := []int{1,2,3,4}
	sum(nums...)
}

func sum(nums ...int)  {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
