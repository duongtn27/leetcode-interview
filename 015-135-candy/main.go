package main

import "fmt"

func main() {
	nums := []int{5, 4, 3, 4, 5, 1}
	fmt.Println(candy(nums))
}

func candy(ratings []int) int {
	n := len(ratings)
	if n == 1 {
		return 1
	}
	candies := make([]int, len(ratings))
	for i := range candies {
		candies[i] = 1
	}
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	total := candies[n-1]
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
		total += candies[i]
	}
	return total
}
