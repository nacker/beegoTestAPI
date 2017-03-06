package main

import "time"
import "fmt"

func main() {
	// Timer 代表了未来的一个事件，你告诉timer需要等待多久，然后
	// 计时器提供了一个通道，这个通道将在等待的时间结束后得到通知，
	// 这里的timer将等待2秒
	timer1 := time.NewTimer(time.Second * 2)

	// 这里`<-timer1.C`在timer的通道`C`上面阻塞等待，直到有个值发送给该
	// 通道，通知通道计时器已经等待完成。
	// timer.NewTimer方法获取的timer1的结构体定义为
	// type Ticket struct{
	//  C <-chan Time
	//}
	<-timer1.C
	fmt.Println("Timer 1 expired")

	// 如果你仅仅需要等待的话，你可以使用`time.Sleep`，而timer的
	// 独特之处在于你可以在timer等待完成之前取消等待。
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}