package main

import "fmt"

func main() {
	str := "hello world!!!"
	s := []rune(str)
	s[6] = 'g'
	s = s[:8]
	str = string(s)
	fmt.Printf("%#v\n", str)
}
