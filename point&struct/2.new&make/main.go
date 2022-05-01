package main

import "fmt"

func main() {
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	fmt.Println(*a)
	fmt.Println(*b)

	c := make(map[string]int, 10)
	c["薛文朝"] = 100
	fmt.Println(c)
}
