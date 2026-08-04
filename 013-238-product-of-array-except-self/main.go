package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, 0, n)
	left := 1
	for i := 0; i < n; i++ {
		result = append(result, left)
		left *= nums[i]
	}

	right := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= right
		right *= nums[i]
	}
	return result
}
