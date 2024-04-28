package main

import (
	"fmt"
	"my_leetcode/internal/tools"
)

type dp struct {
	finishedCount int

	buy  int
	sell int
}

func maxProfit(k int, prices []int) int {
	dpList := make([]*dp, k)
	for i := range dpList {
		dpList[i] = &dp{
			i, -prices[0], 0,
		}
	}
	for i := 1; i < len(prices); i++ {
		price := prices[i]
		for kk := k - 1; kk >= 0; kk-- {
			var preSell int
			if kk > 0 {
				preSell = dpList[kk-1].sell
			}
			dpList[kk].sell = tools.IntMax(dpList[kk].sell, price+dpList[kk].buy)
			dpList[kk].buy = tools.IntMax(dpList[kk].buy, preSell-price)
		}
	}
	return dpList[k-1].sell
}

func maxProfit2(k int, prices []int) int {
	if k == 0 || len(prices) <= 0 {
		return 0
	}

	// 最小dp，用于计算”1次售卖的dp利润“
	minDp := make([]int, len(prices))
	minDp[0] = prices[0]
	for i := 1; i < len(prices); i++ {
		minDp[i] = min(minDp[i-1], prices[i])
	}

	// 第1级: 表示第n次交易
	// 第2级: 0表示购入dp，1表示出售dp
	// 第3级: 利润dp
	ioDpList := make([][2][]int, k)
	for i := range ioDpList {
		ioDpList[i][0] = make([]int, len(prices))
		ioDpList[i][1] = make([]int, len(prices))
	}

	for i := 1; i < len(prices); i++ {
		for j := 0; j < k; j++ {
			if j == 0 {
				// 计算第一次的利润
				ioDpList[0][1][i] = max(ioDpList[0][1][i-1], prices[i]-minDp[i])
				continue
			}

			// 非第一次的
			doubleJ := 2 * (j - 1)
			if i >= doubleJ+2 {
				// 购入
				if i == doubleJ+2 {
					ioDpList[j][0][i] = ioDpList[j-1][1][i] - prices[i]
				} else {
					ioDpList[j][0][i] = max(ioDpList[j][0][i-1], ioDpList[j-1][1][i]-prices[i])
				}
			}
			if i >= doubleJ+3 {
				// 出售
				if i == doubleJ+3 {
					ioDpList[j][1][i] = ioDpList[j-1][0][i] + prices[i]
				} else {
					ioDpList[j][1][i] = max(ioDpList[j][1][i-1], ioDpList[j-1][0][i]+prices[i])
				}
			}
		}
	}

	e := len(prices) - 1
	ret := ioDpList[0][1][e]
	for i := 1; i < k; i++ {
		ret = max(ret, ioDpList[i][1][e])
	}
	return ret
}

func main() {
	prices1 := []int{3, 2, 6, 5, 0, 3}
	ret1 := maxProfit2(2, prices1)
	if ret1 != 7 {
		panic(fmt.Errorf("case188 error, not expect result 7: %d", ret1))
	}

}
