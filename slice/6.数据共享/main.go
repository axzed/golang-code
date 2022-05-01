package main

import "fmt"

func main() {
	slice1 := make([]int, 4)
	fmt.Println(slice1)
	slice2 := slice1[1:3]
	// append之后为slice1重新分配了内存,slice2不再和新的slice1共享数据
	slice1 = append(slice1, 0)
	slice1[1] = 2
	slice2[1] = 6
	fmt.Println(slice1)
	fmt.Println(slice2)
}
