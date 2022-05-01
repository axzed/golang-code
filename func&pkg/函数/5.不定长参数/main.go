package main

import "fmt"

// ...int代表可以接收0个或者多个参数值
func Test(args ...int) {
	for i := 0; i < len(args); i++ {
		fmt.Print(args[i])
	}
	fmt.Println("")
}

func Test2(args ...int) {
	for i, data := range args {
		fmt.Println("编号为:", i)
		fmt.Println("值为:", data)
	}
}

/*
	这种定义参数列表的情况要注意下面的情况
	1.不定参数列表一定要在最后面
	2.确定的参数一定要传入值
*/
func Test3(a int, b int, args ...int) {
	for i, data := range args {
		fmt.Println("编号为:", i)
		fmt.Println("值为:", data)
	}
	fmt.Println(a, b)
}

/*
	for range循环中省略下标参数
*/
func sum(args ...int) {
	var sum int
	for _, data := range args {
		sum += data
	}
	fmt.Println(sum)
}

func main() {
	Test(1)
	Test(1, 2)
	Test(1, 2, 3)

	// Test2(1, 2, 3, 4)

	// Test3(1, 2, 3, 4, 5, 6)

	sum(1, 2, 3, 4, 5)
}
