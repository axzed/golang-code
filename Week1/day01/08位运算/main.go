package main

import "fmt"

func main081() {
	// 按位与
	fmt.Println(12 & 10) // 8
	fmt.Printf("12 & 10的结果是: 十进制%d, 二进制%b\n", 12 & 10, 12 & 10)

	// 按位或
	fmt.Println(12 | 10) // 14
	fmt.Printf("12 | 10的结果是: 十进制%d, 二进制%b\n", 12 | 10, 12 | 10)

	// 按位异或(两个位不同就为真)
	fmt.Println(12 ^ 10)
	fmt.Printf("12 ^ 10的结果是: 十进制%d, 二进制%b\n", 12 ^ 10, 12 ^ 10)
}

func main() {
	// 移位运算
	fmt.Println(12 << 2) // 48
	fmt.Println(12 << 3) // 96
	fmt.Println(12 << 4) // 192
}