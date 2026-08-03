package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	res := removeDuplicates(nums)
	fmt.Println(res, nums)
}

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	i, j := 2, 2
	for j < n {
		if nums[j] != nums[i-2] {
			nums[i] = nums[j]
			i++
		}
		j++
	}
	return i
}
