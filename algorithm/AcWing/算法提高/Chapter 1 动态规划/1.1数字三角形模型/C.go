package main

import "fmt"

const N int = 110
const INF int = 0x3f3f3f3f
var a, dp [N][N]int

func min(x, y int) int {
	if x < y {
		return x
	} else {
		 return y
	}
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	for i := 1; i <= n; i ++ {
		for j := 1; j <= n; j ++ {
			fmt.Scanf("%d", &a[i][j])
		}
	}

	for i := 1; i <= n; i ++ {
		for j := 1; j <= n; j ++ {
			dp[i][j] = INF
			if (i == 1 && j == 1) {
				dp[i][j] = a[i][j]
			}
			if i > 1 {
				dp[i][j] = min(dp[i][j], dp[i - 1][j] + a[i][j]) 
			}
			if j > 1 {
				dp[i][j] = min(dp[i][j], dp[i][j - 1] + a[i][j]) 
			}
		}
	}

	fmt.Printf("%d", dp[n][n])

}
