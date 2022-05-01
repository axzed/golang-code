package main

import "fmt"

// 常量 变量 表达式
func main021() {
	// 常量定义
	const pi = 3.14

	// 变量定义
	var r = 1.0
	area := pi * r * r
	fmt.Println("圆的面积是:", area)

	r = 2.0
	area = pi * r * r
	fmt.Println("圆的面积是:", area)
}

// 一次性声明多个变量和常量
func main022() {
	const (
		lightSpeed = 3e5
		earthCircle = 40000
	)
	fmt.Println("光速常量", lightSpeed, "地球周长是", earthCircle)

	var (
		age = 20
		height = 180
	)
	fmt.Println(age, height)
}
