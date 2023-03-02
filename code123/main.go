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

func main() {
	src1 := []int{3, 3, 5, 0, 0, 3, 1, 4}
	ret1 := maxProfit(src1)
	if ret1 != 6 {
		panic(fmt.Errorf("code123 case1 error, not expect result 6: %d", ret1))
	}
}
