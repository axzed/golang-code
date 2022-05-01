package main

import "fmt"

func lastPass() {
	fmt.Println("hello world")
}

func main121() {
	defer lastPass()
	// 配合匿名函数使用
	defer func() {
		fmt.Println("匿名函数")
		fmt.Println("结构上要成一个整体,但不必写成函数")
	}()
	fmt.Println("薛文朝")
}

func main122() {
	fmt.Println("开门")
	defer fmt.Println("关门")
	fmt.Println("开灯")
	defer fmt.Println("关灯")
	// 为了不忘记,所以现将操作挂起

	fmt.Println("上课")
	fmt.Println("写代码")
	fmt.Println("做练习")

	// fmt.Println("关门")
	// fmt.Println("关灯")
}

func main() {
	fmt.Println("打开网络")
	defer fmt.Println("关闭网络")

	fmt.Println("打开文件")
	defer fmt.Println("关闭文件")

	fmt.Println("打开数据库")
	fmt.Println("关闭数据库")
}