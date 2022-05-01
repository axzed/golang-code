package main

import "fmt"

func Test(a int, b int) (sum int) {
	sum = a + b
	return
}

// 定义函数类型
type FuncType func(a int, b int) int

func main() {
	var s int
	// 定义函数类型的变量
	var result FuncType
	result = Test
	s = result(3, 6)
	fmt.Println(s)
}
