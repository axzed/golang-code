package main

import "fmt"

func test(n1 int) {
	n1 = n1 + 1
	fmt.Println("test() n1 = ", n1)
}

// 一个函数getSum
func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	fmt.Println("getSum sum = ", sum) // 30
	// 当函数有return语句时,就是将结果返回调用者
	// 说调用我,我就返回给谁
	return sum
}

func main() {
	n1 := 10
	// 调用test
	test(n1)
	fmt.Println("main() n1 = ", n1)

	sum := getSum(10, 20) 
	fmt.Println("main sum = ", sum)
}
