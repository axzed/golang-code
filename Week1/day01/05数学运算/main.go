package main

import (
	"fmt"
	"math"
)

// 基本数学运算
func main051() {
	var a int = 2
	var b = 3
	c := 4
	var d, e int = 5, 6
	var f, g = 7, 8
	fmt.Println(a, b, c, d, e, f, g)

	// 数学运算 + - * / % 不解释
}

// 四舍五入 乘方 开方 三角函数
func main052() {
	// 四舍五入
	fmt.Println(math.Round(3.4))
	fmt.Println(math.Round(3.5))
	// 先对绝对值四舍五入再加负号
	fmt.Println(math.Round(-3.4))

	// 绝对值
	fmt.Println(math.Abs(-3.4)) // 3.4

	// 乘方
	fmt.Println(math.Pow(2, 3)) // 2的三次方 --> 8

	// 开方
	fmt.Println(math.Sqrt(9)) // 3
}

// 整型和浮点型的互换
func main053() {
	// 把整数化为浮点数
	var a = 123
	b := float64(a)
	fmt.Printf("b的类型是%T,值是%v\n", b, b)

	// 浮点数化为整数
	var c = 123.45
	d := int64(c)
	fmt.Printf("d的类型是%T,值是%v", d, d)

	// 纯舍 纯入
	fmt.Println(math.Floor(3.99)) // 3
	fmt.Println(math.Ceil(3.01))  // 4
}
