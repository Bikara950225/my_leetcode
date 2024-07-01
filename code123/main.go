package main

// dp模型
// buy1表示第一次买入的利润dp：buy1 = max(buy1, -price)
// sell1表示完成第一次交易的利润dp：sell1 = max(sell1, price+preBuy1); price+preBuy1意思是当前天数卖出股票的利润
// buy2表示第二次买入的利润dp：buy2 = max(buy2, preSell2-price); preSell2-price意思是计算当天买入股票的利润
// sell2表示完成第二次交易的利润dp：sell2 = max(sell2, price+preBuy2); price+preBuy2表示当天卖出股票的利润
func maxProfit(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for _, price := range prices[1:] {
		sell2 = max(sell2, price+buy2)
		buy2 = max(buy2, sell1-price)

		sell1 = max(sell1, price+buy1)
		buy1 = max(buy1, -price)
	}
	return sell2
}

func maxProfit2(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}

	minp := prices[0]
	// first profit dp
	fpDp := make([]int, len(prices))
	// second dp
	siDp, soDp := make([]int, len(prices)), make([]int, len(prices))
	for i, price := range prices[1:] {
		i++
		// 先算第一次的dp
		fpDp[i] = max(fpDp[i-1], price-minp)
		minp = min(minp, price)
		// 再算第二次
		if i > 1 {
			if i == 2 {
				siDp[i] = fpDp[i-1] - price
			} else {
				siDp[i] = max(siDp[i-1], fpDp[i-1]-price)
			}
		}
		if i > 2 {
			soDp[i] = max(soDp[i-1], siDp[i-1]+price)
		}
	}
	tail := len(prices) - 1
	return max(fpDp[tail], soDp[tail])
}
