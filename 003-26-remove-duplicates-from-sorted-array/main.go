package main

import "fmt"

func main() {
	nums1 := []int{1, 1, 2}
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	res1 := removeDuplicates(nums1)
	res2 := removeDuplicates(nums2)
	fmt.Println(res1, nums1)
	fmt.Println(res2, nums2)
}

func removeDuplicates(nums []int) int {
	i := 1
	j := 1
	for j < len(nums) {
		if nums[j] != nums[j-1] {
			nums[i] = nums[j]
			i++
		}
		j++
	}

	return i
}
