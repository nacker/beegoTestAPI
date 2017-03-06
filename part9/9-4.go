package main

import "os"
import "fmt"

func main() {

	// `os.Args`提供了对命令行参数的访问，注意该
	// 切片的第一个元素是该程序的运行路径，而
	// `os.Args[1:]`则包含了该程序的所有参数
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// 你可以使用索引的方式来获取单个参数
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
