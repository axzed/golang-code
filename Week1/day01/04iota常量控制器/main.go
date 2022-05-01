package main

import "fmt"

/*定义所有星期*/
const (
	_ = iota
	mondy
	tuesday
	wednesday
	thursday
	friday
	saturday
	sunday
)

func main() {
	fmt.Println(mondy, tuesday, wednesday, thursday, friday, saturday, sunday)

}
