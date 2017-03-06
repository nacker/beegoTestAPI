package main

import "time"
import "fmt"

func main() {

	// Ticker使用和Timer相似的机制，同样是使用一个通道来发送数据。
	// 这里我们使用range函数来遍历通道数据，这些数据每隔500毫秒被
	// 发送一次，这样我们就可以接收到
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// Ticker和Timer一样可以被停止。一旦Ticker停止后，通道将不再
	// 接收数据，这里我们将在1500毫秒之后停止
	time.Sleep(time.Millisecond * 1500)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}