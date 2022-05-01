package main

import "fmt"

// 结构体定义
type person struct {
	name string
	city string
	age  int8
}

func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

func main() {
	var p1 person
	p1.name = "xuewenchao"
	p1.city = "guangzhou"
	p1.age = 18
	fmt.Printf("p1 = %v\n", p1)
	fmt.Printf("p1 = %#v\n", p1)

	// 匿名结构体
	var user struct {
		Name string
		Age  int
	}
	user.Name = "xiaowen"
	user.Age = 18
	fmt.Printf("%#v\n", user)

	// 通过new对结构体进行实例化
	p2 := new(person)
	fmt.Printf("%T\n", p2) // *main.person
	fmt.Printf("p2 = %#v\n", p2)

	p3 := &person{}
	fmt.Printf("%T\n", p3)
	fmt.Printf("p3 = %#v\n", p3)
	// 语法糖p3.name = "xuewenchao"  ---> (*p3).name = "xuewenchao"

	var p4 person
	fmt.Printf("p4 = %#v\n", p4)

	p5 := newPerson("xuewenchao", "guangzhou", 18)
	fmt.Printf("p5 = %#v\n", p5)
}
