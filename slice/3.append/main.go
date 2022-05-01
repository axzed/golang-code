package main

import "fmt"

func main() {
	s := make([]int, 5, 8)
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))
	fmt.Println(s)

	// append追加数据, 默认向后压入, 和vector差不多
	s = append(s, 1)
	fmt.Println(s)

	// 切片可以按索引改变值
	s[0] = 1
	fmt.Println(s)

	s = append(s, 2, 3)
	fmt.Println(len(s), cap(s))

	// cap不够的时候, 会直接将cap * 2
	s = append(s, 4)
	fmt.Println(len(s), cap(s))

	// 当容量小于1024的时候是按照两倍的扩容,大于1024之后就不按照两倍容量扩容了

	slice := make([]int, 0, 1)
	oldCap := cap(s)
	for i := 0; i < 2000; i ++ {
		slice = append(slice, i) 
		newCap := cap(slice) 
		if oldCap < newCap {
			fmt.Printf("cap: %d ====> %d\n", oldCap, newCap)
			oldCap = newCap
		}
	}
}
