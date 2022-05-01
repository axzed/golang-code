package main

import "fmt"

func Test() func() int {
	var x int
	return func() int {
		x ++
		return x
	}
}

func main() {
	f := Test()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	// 这是地址
	// fmt.Println(Test())
	// fmt.Println(Test())
	// fmt.Println(Test())
}
