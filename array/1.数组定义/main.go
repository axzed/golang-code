package main

import "fmt"

func main() {
	var intArr [3]int // int占8个字节
	fmt.Println(intArr)
	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30
	fmt.Println(intArr)

	fmt.Printf("%p, %p, %p, %p", &intArr, &intArr[0], &intArr[1], &intArr[2])

}
