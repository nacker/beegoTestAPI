package main

import "fmt"

func main()  {
	// `var` 关键字用来定义一个或者多个变量
	var a string = "initial"
	fmt.Println(a)

	// 你一次可以定义多个变量
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go会推断出具有初始值的变量的类型
	var d = true
	fmt.Println(d)

	//定义变量时，没有给出初始值的变量被默认初始化为零值
	//整型的零值就是0
	var e int
	fmt.Println(e)

	//":=" 语法是同时定义和初始化变量的快捷方式
	f := "short"
	fmt.Println(f)

}
