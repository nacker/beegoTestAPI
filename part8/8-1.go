package main

import "fmt"
import "time"

func main() {

	// 从获取当前时间开始
	var now = time.Now()
	fmt.Println(now)

	// 你可以提供年，月，日等来创建一个时间。当然时间
	// 总是会和地区联系在一起，也就是时区
	var then = time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	fmt.Println(then)

	// 你可以获取时间的各个组成部分
	fmt.Println(then.Year())
	fmt.Println(then.Month())
	fmt.Println(then.Day())
	fmt.Println(then.Hour())
	fmt.Println(then.Minute())
	fmt.Println(then.Second())
	fmt.Println(then.Nanosecond())
	fmt.Println(then.Location())

	// 输出当天是周几，Monday-Sunday中的一个
	fmt.Println(then.Weekday())

	// 下面的几个方法判断两个时间的顺序，精确到秒
	fmt.Println(then.Before(now))
	fmt.Println(then.After(now))
	fmt.Println(then.Equal(now))

	// Sub方法返回两个时间的间隔(Duration)
	var diff = now.Sub(then)
	fmt.Println(diff)

	// 可以以不同的单位来计算间隔的大小
	fmt.Println(diff.Hours())
	fmt.Println(diff.Minutes())
	fmt.Println(diff.Seconds())
	fmt.Println(diff.Nanoseconds())

	// 你可以使用Add方法来为时间增加一个间隔
	// 使用负号表示时间向前推移一个时间间隔
	fmt.Println(then.Add(diff))
	fmt.Println(then.Add(-diff))
}
