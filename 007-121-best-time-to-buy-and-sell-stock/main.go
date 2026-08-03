package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit(prices)
	fmt.Println(res)
}

func maxProfit(prices []int) int {
	minPrc := prices[0]
	maxPrf := 0

	for _, price := range prices {
		if price < minPrc {
			minPrc = price
			continue
		}
		if price-minPrc > maxPrf {
			maxPrf = price - minPrc
		}
	}
	return maxPrf
}
