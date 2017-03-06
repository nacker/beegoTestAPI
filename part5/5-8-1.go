package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// 这里是一个非阻塞的从通道接收数据，如果messages通道有数据
	// 可以接收，那么select将运行`<-messages`这个case，否则的话
	// 程序立刻执行default选项后面的语句
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// 非阻塞通道发送数据也是一样的
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// 在default前面，我们可以有多个case选项，从而实现多通道
	// 非阻塞的选择，这里我们尝试从messages和signals接收数据
	// 如果有数据可以接收，那么执行对应case后面的逻辑，否则立刻
	// 执行default选项后面的逻辑
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
