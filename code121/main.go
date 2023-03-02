package main

import (
	"fmt"
	"my_leetcode/internal/tools"
)

func maxProfit(prices []int) (ret int) {
	dp := prices[0]
	for i := 1; i < len(prices); i++ {
		profit := prices[i] - dp
		if profit > ret {
			ret = profit
		}
		dp = tools.IntMin(dp, prices[i])
	}
	return
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	ret := maxProfit(prices)
	expectRet := 5 // 第2天购入，第5天卖出的利润，6-1=5

	if expectRet != ret {
		panic(fmt.Errorf("code121 error, not expect result: %v", ret))
	}
}
