package main

import "fmt"

func main() {
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice := arr[1:4]
	// 取下标循环, 记得一定要用len,因为是已经初始化的长度
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v] = %v", i, slice[i])
		fmt.Println()
	}

	fmt.Println()

	// for range 方式遍历切片
	for i, v := range slice {
		fmt.Println(i, v)
	}

}
