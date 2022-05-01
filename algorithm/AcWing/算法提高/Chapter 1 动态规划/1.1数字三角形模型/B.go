package main

import (
	"fmt"
)

const N int = 110

var a, dp [N][N]int

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func solve() {
	var n, m int
	fmt.Scanf("%d%d", &n, &m)

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Scanf("%d", &a[i][j])
		}
	}

	dp[1][1] = a[1][1]

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = max(dp[i - 1][j], dp[i][j - 1]) + a[i][j]
		}
	}

	fmt.Println(dp[n][m])
}

func main() {
	var T int
	fmt.Scanf("%d", &T)
	for {
		if T == 0 {
			break
		}
		solve()
		T--
	}
}
