package main

import "fmt"

func main()  {
	// 和数组不同的是，切片的长度是可变的。
	// 我们可以使用内置函数make来创建一个长度不为零的切片
	// 这里我们创建了一个长度为3，存储字符串的切片，切片元素
	// 默认为零值，对于字符串就是""。
	s := make([]string,3)
	fmt.Println("emp",s)
}
