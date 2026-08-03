package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums))
}

func jump(nums []int) int {
	jumps := 0
	currentEnd := 0
	furthest := 0

	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > furthest {
			furthest = i + nums[i]
		}

		if i == currentEnd {
			jumps++
			currentEnd = furthest
		}
	}

	return jumps
}
