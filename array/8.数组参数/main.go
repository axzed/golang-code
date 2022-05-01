package main

import (
	"fmt"
	"math/rand"
	"time"
)

func modify(a [5]int) {
	a[0] = 555
	fmt.Println("modify a = ", a)
}

func GetLongest(names [5]string) (max string) {
	max = names[0]
	for _, v := range names {
		fmt.Println(len(v))
		if len(v) > len(max) {
			max = v
		}
	}
	return
}

func GetAvg(nums [5]int) (avg float64) {
	var sum int
	for _, v := range nums {
		sum += v
	}
	avg = float64(sum / len(nums))
	return
}

// 设置随机数
func RandomFunc() {
	// 设置种子,只需一次
	rand.Seed(123)
	for i := 0; i < 5; i++ {
		fmt.Println("rand = ", rand.Int()) // 随机很大的数
	}
}

// 解决随机值固定的问题 rand.Intn()是个设置限定值的好函数
func RandomFunc02() {
	rand.Seed(time.Now().UnixNano()) // 用系统当前的实践作为种子参数
	for i := 0; i < 5; i ++ {
		fmt.Println("rand = ", rand.Intn(100)) // 限制在100以内的数
	}
}

func main() {
	// 1.
	a := [...]int{1, 2, 3, 4, 5}
	modify(a)                    // 传递数组
	fmt.Println("main: a = ", a) // 不会改变a[0],这里要与其它语言区分开来

	// 2.
	strs := [...]string{"马龙", "迈克尔乔丹", "雷吉米勒", "蒂姆邓肯", "科比布莱恩"}
	maxName := GetLongest(strs)
	fmt.Println(maxName)

	// 3.
	nums := [...]int{1, 33, 44, 555, 3324}
	avg := GetAvg(nums)
	fmt.Println(avg)

	// 4.
	RandomFunc()
	RandomFunc02()
}
