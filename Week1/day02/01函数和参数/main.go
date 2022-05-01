package main

import "fmt"

func main() {
	sayHello()
	saySombody("xuewenchao")
	sayMany("xuewenchao", 10)
	sayLoveToThem("xuewenchao", "hh", "hehe")
}

func sayHello() {
	fmt.Println("hello world")
}

func saySombody(name string) {
	fmt.Printf("hello %s", name)
}

func sayMany(name string, n int) {
	fmt.Println("hello")
	for i := 0; i < n; i++ {
		fmt.Printf("hello %s\n", name)
	}
}

// 不定长参数 切片 类似于c++的vector
func sayLoveToThem(names ...string) {
	fmt.Println("i love you")
	fmt.Printf("%s \n", names)
	for i, name := range(names) {
		fmt.Printf("我爱No.%d %s\n", i, name)
	}
}
