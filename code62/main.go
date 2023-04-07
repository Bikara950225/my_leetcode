package main

import "fmt"

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

// 用深度优先来实现，会超时，写来找感觉
func uniquePathsWithDfs(m int, n int) int {
	return 0
}

func main() {
	ret1 := uniquePaths(3, 7)
	if ret1 != 28 {
		panic(fmt.Errorf("code62 error, ret1 != 28: %d", ret1))
	}
}
