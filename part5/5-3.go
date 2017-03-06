package main

import "fmt"

func main() {

	// 这里我们定义了一个可以存储字符串类型的带缓冲通道
	// 缓冲区大小为2
	messages := make(chan string, 2)

	// 因为messages是带缓冲的通道，我们可以同时发送两个数据
	// 而不用立刻需要去同步读取数据
	messages <- "buffered"
	messages <- "channel"

	// 然后我们和上面例子一样获取这两个数据
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}