package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	max := 0
	h := 0
	for i < j {
		h = min(height[i], height[j]) * (j - i)
		if h > max {
			max = h
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}
