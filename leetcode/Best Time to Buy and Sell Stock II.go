package main

// [7,1,5,3,6,4]

func maxProfit(prices []int) int {
	profit := 0
	if len(prices) < 2 {
		return profit
	}
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}
