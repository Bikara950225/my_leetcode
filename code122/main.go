package main

func maxProfit(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}
	// dp1为持有股票的dp; dp2为没有股票的dp
	dp1, dp2 := -prices[0], 0
	for _, price := range prices[1:] {
		dp2 = max(dp2, price+dp1)
		dp1 = max(dp1, dp2-price)
	}
	return dp2
}

func maxProfit2(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}
	dp := [2]int{-prices[0], 0} // i/o
	for _, price := range prices[1:] {
		dp[1] = max(dp[1], dp[0]+price)
		dp[0] = max(dp[0], dp[1]-price)
	}
	return dp[1]
}
