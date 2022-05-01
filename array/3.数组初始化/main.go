package main

import "fmt"

func main() {
	// 1.全部初始化
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("a = ", a)

	// 我喜欢这种初始化数组的方式
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// 部分初始化可以指定

	// 通过初始化确定数组长度 : 简易多用这种方式初始化数组
	c := [...]int{1, 2, 3}
	fmt.Println(c)
	fmt.Println(len(c))
}
