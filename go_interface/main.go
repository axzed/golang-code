package main

import "fmt"

// 接口
type Sayer interface {
	Say()
}

// 通用的函数接收Sayer类型
func MakeHungry(s Sayer) {
	s.Say()
}

type Cat struct{}

func (c Cat) Say() {
	fmt.Println("喵喵喵")
}

type Dog struct{}

func (d Dog) Say() {
	fmt.Println("汪汪汪")
}

func main() {
	var c Cat
	MakeHungry(c)
	var d Dog
	MakeHungry(d)
}
