package main

import "fmt"

// Person 结构体
type Person struct {
	name string
	age  int8
}

// NewPerson 构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// Dream Person做梦的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学习变好", p.name)
}

func main() {
	p1 := NewPerson("xuewenchao", 18)
	p1.Dream()
}
