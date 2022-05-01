package main

import "fmt"

func main() {
	var m map[int]string
	fmt.Println(m)
	fmt.Println(len(m))

	m2 := make(map[int]string)
	fmt.Println("m2 = ", m2)
	fmt.Println("len = ", len(m2))

	m2[1] = "张三"
	m2[2] = "李四"
	m2[3] = "王五"
	m2[4] = "赵六"

	fmt.Println("m2 = ", m2)
	fmt.Println("len = ", len(m2))

	// 直接初始化
	m3 := map[int]string{1: "java", 2: "go", 3: "c++"}
	fmt.Println("m3 = ", m3)

	var ch rune
	fmt.Println(ch)
}
