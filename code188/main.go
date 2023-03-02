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

func main() {
	prices1 := []int{3, 2, 6, 5, 0, 3}
	ret1 := maxProfit(2, prices1)
	if ret1 != 7 {
		panic(fmt.Errorf("case188 error, not expect result 7: %d", ret1))
	}

}
