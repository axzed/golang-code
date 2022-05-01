package main

import "fmt"

// 常量定义(迭代定义)
const (
	// 定义每分钟的秒数
	SecondsPerMinute = 60
	// 定义每小时的秒数
	SecondsPerHour = SecondsPerMinute * 60
	// 定义每天的的秒数
	SecondPerDay = SecondsPerHour * 24
)

func resolveTime(seconds int) (day, hour, minute int) {
	day = seconds / SecondPerDay
	hour = seconds / SecondsPerHour
	minute = seconds / SecondsPerMinute
	return
}

func main() {
	// 将返回值作为打印参数
	fmt.Println(resolveTime(1000))
	// 只获取小时和分钟
	_, hour, minute := resolveTime(18000)
	fmt.Println(hour, minute)
	// 只获取天
	day, _, _ := resolveTime(90000)
	fmt.Println(day)
}
