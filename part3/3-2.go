package main

import "fmt"

func main()  {
	// 获取函数的两个返回值
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// 如果你只对多个返回值里面的几个感兴趣
	// 可以使用下划线(_)来忽略其他的返回值
	_, c := vals()
	fmt.Println(c)
}

// 这个函数的返回值为两个int
func vals() (int , int)  {
	return 3, 7
}