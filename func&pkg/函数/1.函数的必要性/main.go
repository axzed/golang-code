package main

import (
	"fmt"
	"os/exec"
)

func main011() {
	fmt.Println("notepad")
	exec.Command("notepad").Run()

	fmt.Println("mspaint")
	exec.Command("mspaint").Run()

	fmt.Println("calc")
	exec.Command("calc").Run()
}

func main() {
	GoCmd("notepad")
}

// 函数将一定的代码抽象后复用
// 将重复的代码写到一起去
// 修改简洁
func GoCmd(cmd string) {
	fmt.Println(cmd) 
	exec.Command(cmd).Run()
}
