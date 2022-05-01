package main

import "fmt"

const N, INF int = 510, 1e9

var a, dp [N][N]int

func main() {
	var n int
	fmt.Scanf("%d", &n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Scanf("%d", &a[i][j])
		}
	}

	for i := n; i >= 1; i-- {
		for j := i; j >= 1; j-- {
			dp[i][j] = max(dp[i+1][j], dp[i+1][j+1]) + a[i][j]
		}
	}

	fmt.Printf("%d", dp[1][1])
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
