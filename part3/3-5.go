package main

import "fmt"

/// Go语言支持递归函数，这里是一个经典的斐波拉切数列的列子。

// fact函数不断地调用自身，直到达到基本状态fact(0)
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}
