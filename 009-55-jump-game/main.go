package main

import "fmt"

func main() {
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
}

func canJump(nums []int) bool {
	furthest := 0
	for i := range nums {
		if i > furthest {
			return false
		}
		furthest = max(furthest, i+nums[i])
		if furthest >= len(nums)-1 {
			return true
		}
	}
	return true
}
