package main

import "fmt"
import "os"

func main() {
	// 当使用`os.Exit`的时候defer操作不会被运行，
	// 所以这里的``fmt.Println`将不会被调用
	defer fmt.Println("!")

	// 退出程序并设置退出状态值
	os.Exit(3)
}
