package main

import "fmt"

func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

func main() {
	// 匿名函数直接调用
	func(data int) {
		fmt.Println(data)
	}(100)

	// 匿名函数赋值给一个函数变量
	f := func(data int) {
		fmt.Println("hello", data)
	}
	f(100)

	// 使用匿名函数打印切片内容, 匿名函数做回调函数
	// 回调函数:自己定义一个函数作为别的函数的参数,并被别的函数使用
	visit([]int{1, 2, 3, 4}, func(v int) {
		fmt.Println(v)
	})
}
