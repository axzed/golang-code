package main

import "fmt"

func main() {
	var a [10]int
	for i := 0; i < len(a); i++ {
		a[i] = i + 1
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// for i, data := range a {
	// 	fmt.Printf("a[%d]=%d\t", i, data)
	// }

	strArr := [...]string{"tom", "jack", "mary"}
	fmt.Println(strArr)
	for i, data := range strArr {
		fmt.Println(i, data)
	}
}
