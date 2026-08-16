package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	fmt.Println(minSubArrayLen(target, nums))
}

func minSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	minLen := 0

	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			curLen := right - left + 1
			if minLen == 0 || curLen < minLen {
				minLen = curLen
			}
			sum -= nums[left]
			left++
		}
	}

	return minLen
}
