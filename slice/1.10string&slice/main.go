package main

import "fmt"

func main() {
	// str := "hello world"
	// // string的底层就是一个byte数组,所以也可以向数组那样对切片进行初始化
	// s1 := str[0:5]
	// fmt.Println(s1)
	// s2 := str[6:]
	// fmt.Println(s2)

	// string本身是不可变的,如果想要改变string中的字符,需要进行如下操作(英文字符串):
	str := "Hello World"
	s := []byte(str) // 中文需要开[]rune(str)这种字符切片
	for i, v := range s {
		fmt.Printf("idx : %v, val : %c\n", i, v)
	}
	s[6] = 'G'
	s = s[:8]
	s = append(s, '!')
	str = string(s)
	fmt.Println(str)
}
