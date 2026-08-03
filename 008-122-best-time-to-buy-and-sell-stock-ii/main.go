package main

import "fmt"

func main() {
	prices := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	maxPrf := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxPrf += prices[i] - prices[i-1]
		}
	}
	return maxPrf
}
