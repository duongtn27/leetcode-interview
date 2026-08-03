package main

import "fmt"

func main() {
	nums1 := []int{4, 5}
	val := 5
	res := removeElement(nums1, val)
	fmt.Println(res, nums1)
}

func removeElement(nums []int, val int) int {
	i := 0
	j := len(nums) - 1
	for i <= j {
		if nums[i] == val {
			temp := nums[j]
			nums[j] = nums[i]
			nums[i] = temp
			j--
		} else {
			i++
		}
		fmt.Println(i, j, nums)
	}
	return i
}
