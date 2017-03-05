package main

import "fmt"

func main()  {
	// 这里我们创建了一个具有5个元素的整型数组
	// 元素的数据类型和数组长度都是数组的一部分
	// 默认情况下，数组元素都是零值
	// 对于整数，零值就是0
	var a [5]int
	fmt.Println("emp:", a)

	// 我们可以使用索引来设置数组元素的值，就像这样
	// "array[index] = value"  或者使用索引来获取元素值，
	// 就像这样"array[index]"
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// 内置的len函数返回数组的长度
	fmt.Println("len:",len(a))

	// 这种方法可以同时定义和初始化一个数组
	b := [5]int{1,2,3,4,5}
	fmt.Println("dcl",b)

	// 数组都是一维的，但是你可以把数组的元素定义为一个数组
	// 来获取多维数组结构
	var twoD [2][3]int
	for i := 0; i< 2 ; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
