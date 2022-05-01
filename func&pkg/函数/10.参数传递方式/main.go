package main

import "fmt"

// 用指针接收地址
func test(n1 *int) {
	*n1 = *n1 + 10
	fmt.Println("test() n1 = ", *n1) // 30
}

func main() {
	num := 20
	test(&num)
	fmt.Println("main() num = ", num) // 30
}
