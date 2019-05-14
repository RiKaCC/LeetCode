package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	max := 0
	size := len(prices)

	for i := 0; i < size-1; i++ {
		if prices[i] < prices[i+1] {
			max += prices[i+1] - prices[i]
		}
	}
	return max
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
