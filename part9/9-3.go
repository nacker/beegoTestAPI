package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// 使用缓冲scanner来包裹无缓冲的`os.Stdin`可以让我们
	// 方便地使用`Scan`方法，这个方法会将scanner定位到下
	// 一行的位置
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// `Text`方法从输入中返回当前行
		ucl := strings.ToUpper(scanner.Text())

		// 输出转换为大写的行
		fmt.Println(ucl)
	}

	// 在`Scan`过程中，检查错误。文件结束不会被当作一个错误
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
