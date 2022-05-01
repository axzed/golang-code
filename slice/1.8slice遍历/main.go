package main

import "fmt"

func main() {
	data := [...]int{0, 1, 2, 3, 4, 5}
	slice := data[:]
	for i, v := range slice {
		fmt.Printf("index : %v, value : %v\n", i, v)
	}
}
