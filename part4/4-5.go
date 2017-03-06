package main

import "errors"
import "fmt"

// Go语言里面约定错误代码是函数的最后一个返回值，
// 并且类型是error，这是一个内置的接口

func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New使用错误信息作为参数，构建一个基本的错误
		return -1, errors.New("can't work with 42")
	}
	// 返回错误为nil表示没有错误
	return arg + 3, nil
}

// 你可以通过实现error接口的方法Error()来自定义错误
// 下面我们自定义一个错误类型来表示上面例子中的参数错误
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {

		// 这里我们使用&argError语法来创建一个新的结构体对象，
		// 并且给它的成员赋值
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// 下面的两个循环例子用来测试我们的带有错误返回值的函数
	// 在for循环语句里面，使用了if来判断函数返回值是否为nil是
	// Go语言里面的一种约定做法。
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// 如果你需要使用自定义错误类型返回的错误数据，你需要使用类型断言
	// 来获得一个自定义错误类型的实例才行。
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
