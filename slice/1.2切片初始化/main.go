package main

import (
	"fmt"
)

var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var slice0 []int = arr[2:8]
var slice1 []int = arr[0:6] // 可以简写 var slice []int = arr[:end]
var slice2 []int = arr[5:10] 
var slice3 []int = arr[0:len(arr)] // 可以简写为 var slice []int = arr[:]
var slice4 = arr[:len(arr) - 1] // 去掉切片的最后一个元素	

func main() {
	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("slice0: %v\n", slice0)
	fmt.Printf("slice1: %v\n", slice1)
	fmt.Printf("slice2: %v\n", slice2)
	fmt.Printf("slice3: %v\n", slice3)
	fmt.Printf("slice4: %v\n", slice4)
	
	fmt.Println("---------------------------------------")

	arr2 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	slice5 := arr2[2:8]
	slice6 := arr2[0:6]
	slice7 := arr2[5:10]
	slice8 := arr[0:len(arr2)]

	fmt.Printf("slice5: %v\n", slice5)
	fmt.Printf("slice6: %v\n", slice6)
	fmt.Printf("slice7: %v\n", slice7)
	fmt.Printf("slice8: %v\n", slice8)
}