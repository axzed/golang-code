package main

import "fmt"

func main() {
	// 引用的是同一个底层的数组,在c根据b的扩容展开下后面的元素又出来了
	var a = []int{1, 2, 3, 4}
	fmt.Printf("slice a : %v, len(a) : %v\n", a, len(a))
	b := a[1:2]
	fmt.Printf("slice b : %v, len(b) : %v\n", b, len(b))
	c := b[0:3]
	fmt.Printf("slice c : %v, len(c) : %v\n", c, len(c))
}
