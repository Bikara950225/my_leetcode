package main

import (
	"fmt"
	"my_leetcode/internal/tools"
)

// dp模型
// buy1表示第一次买入的利润dp：buy1 = max(buy1, -price)
// sell1表示完成第一次交易的利润dp：sell1 = max(sell1, price+preBuy1); price+preBuy1意思是当前天数卖出股票的利润
// buy2表示第二次买入的利润dp：buy2 = max(buy2, preSell2-price); preSell2-price意思是计算当天买入股票的利润
// sell2表示完成第二次交易的利润dp：sell2 = max(sell2, price+preBuy2); price+preBuy2表示当天卖出股票的利润
func maxProfit(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		sell2 = tools.IntMax(sell2, prices[i]+buy2)
		buy2 = tools.IntMax(buy2, sell1-prices[i])

		sell1 = tools.IntMax(sell1, prices[i]+buy1)
		buy1 = tools.IntMax(buy1, -prices[i])
	}
	return sell2
}

func maxProfit2(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}
	minDp := make([]int, len(prices))
	minDp[0] = prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minDp[i-1] {
			minDp[i] = prices[i]
		} else {
			minDp[i] = minDp[i-1]
		}
	}
	// 第一次卖出的Dp
	oDp1 := make([]int, len(prices))
	iDp2, oDp2 := make([]int, len(prices)), make([]int, len(prices))
	for i := 1; i < len(prices); i++ {
		oDp1[i] = max(oDp1[i-1], prices[i]-minDp[i])
		if i >= 2 {
			if i == 2 {
				iDp2[i] = oDp1[i] - prices[i]
			} else {
				iDp2[i] = max(iDp2[i-1], oDp1[i]-prices[i])
			}
		}
		if i >= 3 {
			if i == 3 {
				oDp2[i] = iDp2[i] + prices[i]
			} else {
				oDp2[i] = max(oDp2[i-1], iDp2[i]+prices[i])
			}
		}
	}
	e := len(prices) - 1
	rr := []int{oDp1[e], oDp2[e]}
	return max[int](1, rr...)
}

func main() {
	src1 := []int{3, 3, 5, 0, 0, 3, 1, 4}
	ret1 := maxProfit2(src1)
	if ret1 != 6 {
		panic(fmt.Errorf("code123 case1 error, not expect result 6: %d", ret1))
	}
}
