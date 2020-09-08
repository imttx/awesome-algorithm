package leetcode

import "math"

func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	maxProfit := 0

	for _, p := range prices {
		if p < minPrice {
			minPrice = p
		} else if p-minPrice > maxProfit {
			maxProfit = p - minPrice
		}
	}

	return maxProfit
}
