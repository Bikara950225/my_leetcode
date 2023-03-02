package main

import (
	"fmt"
)

var steps = [2]int{1, 2}

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for _, step := range steps {
			if i - step == 0 {
				dp[i] += 1
			} else if i - step > 0 {
				dp[i] += dp[i-step]
			}
		}
	}
	return dp[n]
}

func main() {
	ret1 := climbStairs(2)
	expectRet1 := 2
	if ret1 != expectRet1 {
		panic(fmt.Errorf("code70 error, not expect result: %d", ret1))
	}
}
