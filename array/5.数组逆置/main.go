package main

import "fmt"

func main() {
	intArr := [...]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(intArr)/2; i++ {
		intArr[i], intArr[len(intArr)-i-1] = intArr[len(intArr)-i-1], intArr[i]
	}
	fmt.Println(intArr)
}
