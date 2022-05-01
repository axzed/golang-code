package main

import "fmt"

func main() {
	str := "abcdefg"
	// string不能够直接通过索引改变值
	// 曲线救国, 把string变成一个字符切片
	s1 := []byte(str)
	s1[1] = 'g'
	for _, v := range s1 {
		fmt.Printf("%c", v)
	}
}

