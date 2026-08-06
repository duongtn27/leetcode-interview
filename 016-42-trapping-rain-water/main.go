package main

import "fmt"

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}

func trap(height []int) int {
	leftMax, left, rightMax, right := 0, 0, 0, len(height)-1
	total := 0
	for left < right {
		if height[left] < height[right] {
			leftMax = max(height[left], leftMax)
			total += leftMax - height[left]
			left++
		} else {
			rightMax = max(height[right], rightMax)
			total += rightMax - height[right]
			right--
		}
	}
	return total
}
