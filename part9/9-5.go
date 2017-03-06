package main

import "flag"
import "fmt"

func main() {

	// 基础的标记声明适用于string，integer和bool型选项。
	// 这里我们定义了一个标记`word`，默认值为`foo`和一
	// 个简短的描述。`flag.String`函数返回一个字符串指
	// 针（而不是一个字符串值），我们下面将演示如何使
	// 用这个指针
	wordPtr := flag.String("word", "foo", "a string")

	// 这里定义了两个标记，一个`numb`，另一个是`fork`，
	// 使用和上面定义`word`标记相似的方法
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// 你也可以程序中任意地方定义的变量来定义选项，只
	// 需要把该变量的地址传递给flag声明函数即可
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// 当所有的flag声明完成后，使用`flag.Parse()`来分
	// 解命令行选项
	flag.Parse()

	// 这里我们仅仅输出解析后的选项和任何紧跟着的位置
	// 参数，注意我们需要使用`*wordPtr`的方式来获取最
	// 后的选项值
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
