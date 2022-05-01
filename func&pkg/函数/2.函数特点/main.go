package main

import (
	"fmt"
	"math"
)

// func add(a int, b int) int {
// 	return a + b
// }

func add(a, b int) (int, int) {
	return a + b, a - b
}

func hypot(x, y float64) float64 {
	return math.Sqrt(x * x + y * y)
}

func typedTwoValues() (int, int) {
	return 1, 2
}

func namedRetValues() (a, b int) {
	a = 1
	b = 2
	return 
}

func main() {
	// fmt.Println(hypot(3, 4))
	// a, b := typedTwoValues()
	a, b := namedRetValues()
	fmt.Println(a, b)

}
