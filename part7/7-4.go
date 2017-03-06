package main

import "os"

func main() {

	// 我们使用panic来检查预期不到的错误
	panic("a problem")

	// Panic的通常使用方法就是如果一个函数
	// 返回一个我们不知道怎么处理的错误的
	// 时候，直接终止执行。
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
