package main

import "fmt"

func main() {
	data := [...]int{0, 1, 2, 3, 4, 10: 0}
	s := data[:2:3]
	s = append(s, 100, 200)
	fmt.Println(s, data)
	fmt.Println(&s[0], &data[0])
}
