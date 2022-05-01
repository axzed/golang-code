package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6}
	var max int = a[0]
	var min int = a[0]
	var sum int
	for i := 0; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
		if a[i] < min {
			min = a[i]
		}
		sum += a[i]
	}
	fmt.Println(max, min, sum, sum / len(a))

	// 字符串拼接
	names := [...]string{"老王", "薛文朝", "王哥"}
	var str string
	for i := 0; i < len(names); i ++ {
		str += names[i] + "|"
	}
	fmt.Println(str)
}
