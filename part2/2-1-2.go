package main

import "fmt"

/// 示例2：
/// 可以用new创建数组，并返回数组的指针

func main()  {
	var a = new([5]int)
	test(a)
	fmt.Println(a, len(a))
}

func test(a *[5]int)  {
	a[1] = 5
}