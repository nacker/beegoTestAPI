package main

import "syscall"
import "os"
import "os/exec"

func main() {

	// 本例中，我们使用`ls`来演示。Go需要一个该命令
	// 的完整路径，所以我们使用`exec.LookPath`函数来
	// 找到它
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}
	// `Exec`函数需要一个切片参数，我们给ls命令一些
	// 常见的参数。注意，第一个参数必须是程序名称
	args := []string{"ls", "-a", "-l", "-h"}

	// `Exec`还需要一些环境变量，这里我们提供当前的
	// 系统环境
	env := os.Environ()

	// 这里是`os.Exec`调用。如果一切顺利，我们的原
	// 进程将终止，然后启动一个新的ls进程。如果有
	// 错误发生，我们将获得一个返回值
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
