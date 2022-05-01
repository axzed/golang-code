package main

import "fmt"

// 提供一个值(引用环境),每次调用函数回指定对应的值进行累加
func Accumulate(value int) func() int {
	// 返回一个闭包
	return func() int {
		// 累加
		value ++
		// 返回累加值
		return value
	}
}

// 闭包实现生成器
func playerGen(name string) func() (string, int) {
	hp := 150
	return func() (string, int) {
		return name, hp
	}
}

func main() {
	// 创建一个累加器,初始值为1
	accumulate := Accumulate(1)
	// 累加1并打印
	fmt.Println(accumulate())
	fmt.Println(accumulate())

	generator := playerGen("high noon")
	name, hp := generator()
	fmt.Println(name, hp)
}
