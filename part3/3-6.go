package main

import "fmt"

type Callback func(x,y int) int 

func main()  {
	x, y := 1, 2
	fmt.Println(test(x, y, add))
}

//提供一个接口，让外部去实现
func test(x,y int, callback Callback) int {
	return callback(x, y)
}

func add(x,y int) int  {
	return x + y
}