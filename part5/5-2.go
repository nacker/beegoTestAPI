package main

import "fmt"

func main() {

	// 使用`make(chan 数据类型)`来创建一个Channel
	// Channel的类型就是它们所传递的数据的类型
	messages := make(chan string)

	// 使用`channel <-`语法来向一个Channel写入数据
	// 这里我们从一个新的协程向messages通道写入数据ping
	go func() { messages <- "ping" }()

	// 使用`<-channel`语法来从Channel读取数据
	// 这里我们从main函数所在的协程来读取刚刚写入
	// messages通道的数据
	msg := <-messages
	fmt.Println(msg)
}
