package main

import "fmt"
import "time"

// 这个worker函数将以协程的方式运行
// 通道`done`被用来通知另外一个协程这个worker函数已经执行完成
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 向通道发送一个数据，表示worker函数已经执行完成
	done <- true
}

func main() {

	// 使用协程来调用worker函数，同时将通道`done`传递给协程
	// 以使得协程可以通知别的协程自己已经执行完成
	done := make(chan bool, 1)
	go worker(done)

	// 一直阻塞，直到从worker所在协程获得一个worker执行完成的数据
	<-done
}
