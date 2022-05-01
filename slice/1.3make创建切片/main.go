package main

import "fmt"

var slice0 []int = make([]int, 10)
var slice1 = make([]int, 10)
var slice2 = make([]int, 10, 10)

func main() {
	fmt.Printf("make全局slice0: %v\n", slice0)
	fmt.Printf("make全局slice1: %v\n", slice1)
	fmt.Printf("make全局slice2: %v\n", slice2)

	fmt.Println("------------------------------------")

	slice3 := make([]int, 10)
	slice4 := make([]int, 10)
	slice5 := make([]int, 10, 10)

	fmt.Printf("make局部slice3: %v\n", slice3)
	fmt.Printf("make局部slice4: %v\n", slice4)
	fmt.Printf("make局部slice5: %v\n", slice5)

	// 可以直接创建slice对象,自动分配底层数组
	s1 := []int{0, 1, 2, 3, 4, 5} // 通过初始化表达式构造,可使用索引号
	fmt.Println(s1, len(s1), cap(s1))

	// 使用make来创建,指定len和cap值
	s2 := make([]int, 6, 8)
	fmt.Println(s2, len(s2), cap(s2))

	// 省略cap,相当于cap = len
	s3 := make([]int, 6)
	fmt.Println(s3, len(s3), cap(s3))

	// 使用make动态创建slice,避免了数组必须要用常量做长度的麻烦,还可以用指针直接访问底层数组,退化为普通的数组操作
	s := []int{0, 1, 2, 3}
	p := &s[2] // *int,获取底层数组元素指针
	*p += 100 // *取出地址中的值,然后进行修改
	fmt.Println(s) // 修改后切片指针指向的底层数组的值 0 1 102 3
}
