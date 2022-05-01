package main

import "fmt"

func fire() {
	fmt.Println("fire")
}

func main() {
	// go中函数也是一种变量 声明f为一种函数变量
	var f func()
	// 将变量fire()函数作为值,付给函数变量f,此时f的值为fire函数
	f = fire
	// 对函数变量f进行函数调用,实际上调用的是fire()函数
	f()
}
