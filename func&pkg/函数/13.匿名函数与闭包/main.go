package main

import "fmt"

// 全局匿名函数,将匿名函数赋值给一个全局变量

// 闭包: 就是一个函数和与其相关的引用环境作何的一个整体(实体)
// 累加器
/*
	返回的是一个匿名函数,  但是这个匿名函数引用到函数外的 n ,因此这个匿名函数就和 n 形成一 个整体，构成闭包。
	1) AddUpper  是一个函数，返回的数据类型是 fun (int) int
	2) 闭包的说明: 返回的是一个匿名函数,  但是这个匿名函数引用到函数外的 n ,因此这个匿名函数就和 n 形成一 个整体，构成闭包。
	3) 大家可以这样理解: 闭包是类,  函数是操作，n 是字段。函数和它使用到 n 构成闭包。
	4) 当我们反复的调用 f 函数时，因为 n 是初始化一次，因此每调用一次就进行累计。
	5) 我们要搞清楚闭包的关键，就是要分析出返回的函数它使用(引用)到哪些变量，因为函数和它引 用到的变量共同构成闭包。
*/
func AddUpper() func(int) int {
	var n int = 10
	var str = "hello"
	return func(x int) int {
		n = n + x
		str += string(36)
		fmt.Println("str = ", str)
		return n
	}
}

func main() {
	// 在定义匿名函数的时候就直接调用,这种方式匿名函数只能调用一次
	res := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res = ", res)

	// 赋值给一个函数变量通过函数变量进行调用
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	res2 := a(10, 30)
	fmt.Println("res2 = ", res2)
	res3 := a(90, 30)
	fmt.Println("res3 = ", res3)

	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}
