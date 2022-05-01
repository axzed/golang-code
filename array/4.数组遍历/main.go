package main

import "fmt"

func main() {
	// for-range遍历
	heroes := [...]string{"薛文朝", "超人", "蜘蛛侠"}
	chars := [...]rune{'薛', 'b', 'c'}

	fmt.Println(chars)
	fmt.Println(len(chars))
	chars[0] = 'a'
	for _, v := range chars {
		fmt.Printf("%c", v)
	}
	
	fmt.Println()

	fmt.Println(len(heroes))

	for i, v := range heroes {
		fmt.Println(i, v)
	}

	for _, v := range heroes {
		fmt.Println(v)
	}
}
