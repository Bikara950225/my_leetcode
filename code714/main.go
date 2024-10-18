package main

func maxProfit(prices []int, fee int) int {
	idp, odp := make([]int, len(prices)), make([]int, len(prices))
	idp[0] = -prices[0]

	for i := 1; i < len(prices); i++ {
		price := prices[i]

		odp[i] = max(odp[i-1], idp[i-1]+price-fee)
		idp[i] = max(idp[i-1], odp[i-1]-price)
	}
	e := len(prices) - 1
	return max(idp[e], odp[e])
}
