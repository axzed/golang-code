package main

import "fmt"

func main() {
	str := "你好,世界! hello world"
	s := []rune(str)
	s[3] = '够'
	s[4] = '浪'
	s[12] = 'g'
	s = s[:14]
	str = string(s)
	fmt.Println(str)
}
