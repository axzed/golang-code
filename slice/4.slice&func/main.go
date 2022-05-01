package main

import "fmt"

func InitData(num []int) {
	for i := 0; i < len(num); i++ {
		num[i] = i
	}
}

func InitDataArr(num [5]int) {
	for i := 0; i < 5; i ++ {
		num[i] = i
	}
}

func main() {
	arr := [5]int{0, 0, 0, 0, 0}
	InitDataArr(arr)
	for _, v := range arr {
		fmt.Println(v)
	}
	fmt.Println("-------------------------------")
	s := make([]int, 10)
	InitData(s)
	for _, v := range s {
		fmt.Println(v)
	}
}
