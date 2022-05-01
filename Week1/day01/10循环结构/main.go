package main

import (
	"fmt"
	"time"
)

// 有限次循环
func main101() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// 无限次循环
func main102() {
	for {
		fmt.Println("薛文朝")
		time.Sleep(1 * time.Second)
	}
}

// 退出循环
func main103() {
	cnt := 0
	for {
		fmt.Println("薛文朝")
		time.Sleep(500  * time.Microsecond)
		cnt ++
		if cnt == 10 {
			break
		}
	}
}

func main() {
	// main101()
	// main102()
	// main103()
}
