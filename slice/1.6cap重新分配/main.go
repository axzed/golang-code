package main

import "fmt"

func main() {
	s := make([]int, 0, 1)
	c := cap(s)

	// 成倍数扩容
	for i := 0; i < 50; i++ {
		s = append(s, i)
        n := cap(s);
		if n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
		}
	}
}
