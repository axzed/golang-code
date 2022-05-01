package main

import "fmt"

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

func main() {
	// 指针取值
	a := 10
	b := &a // 取变量a的地址, 将指针保存到b中
	fmt.Printf("type of b : %T\n", b)
	c := *b
	fmt.Printf("type of c : %T\n", c)
	fmt.Printf("value of c : %v\n", c)

	modify1(a)
	fmt.Println(a)

	modify2(&a)
	fmt.Println(a)
}
