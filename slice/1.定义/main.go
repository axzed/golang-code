package main

import "fmt"

func main() {
	intArr := [4]int{1, 2, 3, 4}
	// 自动推导类型创建slice
	s1 := []int{1, 2, 3, 4}
	fmt.Println(s1)
	fmt.Printf("%T:\n", intArr)
	fmt.Printf("%T:\n", s1)
	fmt.Printf("%p:\n", &s1)
	for _, v := range s1 {
		fmt.Printf("%p\n", &v)
	}

	// make创建slice
	s2 := make([]int, 5, 10) // make(切片类型, 长度, 容量)
	// make 没有指定容量的时候,默认 长度 == 容量
	s2[4] = 7
	fmt.Println(len(s2))
}

