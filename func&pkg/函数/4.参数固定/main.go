package main

import "fmt"

func Test(a int, b int) {
	fmt.Printf("a = %d, b = %d", a, b)
}

func main() {
	Test(3, 5)
}
