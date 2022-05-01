package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(32)
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
