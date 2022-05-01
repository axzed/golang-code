package main

import (
	"fmt"
	"time"
)

// 匿名函数 : 微不足道没有复用的价值 但要进行封装(一次性塑料袋)
// 匿名典型场景一: 延时执行一段代码
func main031() {
	defer fmt.Println("我是倒数第一")

	// 打个包
	defer func() {
		fmt.Println("我是倒数第二")
		fmt.Println("我是倒数第三")
	}()
}

// 典型应用场景二: 并发执行一段代码
func main032() {
	// 以下代码跑在主协程
	fmt.Println("文能复习划重点")
	// 跑在子协程
	go func() {
		fmt.Println("我是废物")
		time.Sleep(1 * time.Second)
		fmt.Println("与上面区别")
		time.Sleep(1 * time.Second)
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("xuewenchao")
	time.Sleep(1 * time.Second)
	fmt.Println("hehe")
}

func them(name string, age int) (pow int) {
	fmt.Println(name)
	fmt.Println(age)
	pow = 100
	return pow
}

// 带有参数和返回值的匿名函数
func main() {
	defer them("xuewenchao", 19)
}
