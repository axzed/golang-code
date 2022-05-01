package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("hello xwc", i)
	wg.Done() // 通知wg把计数器 - 1
}

func main() { // 开启了一个主goroutine去执行main这个函数
	wg.Add(10000) // 计数牌 + 1
	for i:= 0; i < 10000; i ++ {
		go hello(i) // 开启了一个goroutine去执行这个函数
	}
	fmt.Println("hello main")
	// time.Sleep(time.Second) // 让主goroutine休息1s
	wg.Wait() // 等所有小弟都干完活
}
