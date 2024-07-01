package main

func maxProfit(prices []int) (ret int) {
	if len(prices) <= 0 {
		return
	}
	dp := prices[0]
	for _, price := range prices[1:] {
		profit := price - dp
		if profit > ret {
			ret = profit
		}
		dp = min(dp, price)
	}
	return
}
