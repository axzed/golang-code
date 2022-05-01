package main

import "fmt"

func main() {
	var a = []int{1, 2, 3}
	fmt.Printf("slice a : %v\n", a)
	var b = []int{4, 5, 6}
	fmt.Printf("slice b : %v\n", b)

	// append: 向slice尾部添加数据,返回新的slice对象
	c := append(a, b...)
	fmt.Printf("slice c : %v\n", c)
	d := append(c, 7)
	fmt.Printf("slice d : %v\n", d)
	e := append(d, 8, 9, 10)
	fmt.Printf("slice e : %v\n", e)
}
