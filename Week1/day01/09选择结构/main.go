package main

import "fmt"

// 选择结构
func main() {
	fmt.Println("请输入你的姓名")

	// 声明一个字符串变量name
	var name string
	// 接收用户输入, 存入name变量所在的地址
	fmt.Scan(&name)

	if name == "薛文朝" {
		fmt.Println("fuck off")
		
		// 结束当前函数的执行
		return
	}
	fmt.Println("你好")

	// 双分支
	if name == "杰哥" {
		fmt.Println("阿伟在床上等着你")
	} else {
		fmt.Println("你好")
	}
	fmt.Println("GAME OVER")

	// 多分支
	if name == "哈哈" {
		fmt.Println(10)
	} else if name == "呵呵" {
		fmt.Println(20)
	} else if name == "呼呼" {
		fmt.Println(30)
	} else {
		fmt.Println(40)
	}
	
	// switch这个不做描述
}

