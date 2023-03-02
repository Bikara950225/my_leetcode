package main

import (
	"fmt"
	"my_leetcode/internal/tools"
)

func maxProfit(prices []int) int {
	// dp1为持有股票的dp; dp2为没有股票的dp
	dp1, dp2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		// 没有股票的dp模型： dp2 = max(preDp2, price+dp1), 前一个dp2和当前卖出股票的利润比对，取最大值
		dp2 = tools.IntMax(dp2, prices[i]+dp1)
		// 持有股票的dp模型： dp1 = max(preDp1, dp2-price), 前一个dp1和当前
		dp1 = tools.IntMax(dp1, dp2-prices[i])
	}
	return tools.IntMax(dp1, dp2)
}

func main() {
	ret := maxProfit([]int{7, 1, 5, 3, 6, 4})
	// 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4 。
	// 随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6 - 3 = 3 。
	// 总利润为 4 + 3 = 7 。
	expectRet := 7
	if ret != expectRet {
		panic(fmt.Errorf("code122 error: not expect result: %v", ret))
	}
}
