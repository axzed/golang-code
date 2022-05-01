package main

import "fmt"

func main() {
	i := sum(3, 4)
	fmt.Println(i)
	i2, i3, i4, i5 := Calc(3, 4)
	fmt.Println(i2, i3, i4, i5)
}

func sum(a, b int) int {
	res := a + b
	return res
}

// 给返回值一个名字 
func GetSum(a, b int) (sum int) {
	res := a + b
	return res
}

// 多返回值 记得预定义函数返回值的名字 	
func Calc(a, b int) (int, int, int, int) {
	sum := a + b
	sub := a - b
	pro := a * b
	div := a / b
	return sum, sub, pro, div
}

// 预定义好了的多返回值的名称 这样可以望文生义 容易观看
func Calc2(a, b int) (sum, sub, pro, div int) {
	sum = a + b
	sub = a - b
	pro = a * b
	div = a / b
	return sum, sub, pro, div
}