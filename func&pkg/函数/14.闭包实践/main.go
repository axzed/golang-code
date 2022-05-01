package main

import (
	"fmt"
	"strings"
)

/*
	1) 编写一个函数 makeSuffix(suffix string)可以接收一个文件后缀名(比如.jpg)，并返回一个闭包
    2) 调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(比如.jpg) ,则返回 文件名.jpg , 如果已经有.jpg 后缀，则返回原文件名。
    3) 要求使用闭包的方式完成
	4) strings.HasSuffix , 该函数可以判断某个字符串是否有指定的后缀。
*/
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	f2 := makeSuffix(".jpg")
	fmt.Println(f2("winter"))
	fmt.Println(f2("bird.jpg"))
}
