package main

import "fmt"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
	}
	for am := 1; am <= amount; am++ {
		for _, coin := range coins {
			if am-coin < 0 || dp[am-coin] == -1 {
				continue
			}
			if am-coin == 0 {
				dp[am] = 1
				break
			}
			if dp[am] == -1 || dp[am-coin]+1 < dp[am] {
				dp[am] = dp[am-coin] + 1
			}
		}
	}
	return dp[amount]
}

func main() {
	// 11 = 5 + 5 + 1
	ret1 := coinChange([]int{1, 2, 5}, 11)
	if ret1 != 3 {
		panic(fmt.Errorf("code322 error, not expect: %d", ret1))
	}

	ret2 := coinChange([]int{3}, 2)
	if ret2 != -1 {
		panic(fmt.Errorf("code322 error, not expect: %d", ret2))
	}

	ret3 := coinChange([]int{3}, 0)
	if ret3 != 0 {
		panic(fmt.Errorf("code322 error, not expect: %d", ret3))
	}
}
