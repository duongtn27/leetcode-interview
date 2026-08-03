package main

import "fmt"

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	res := majorityElement(nums)
	fmt.Println(res, nums)
}

func majorityElement(nums []int) int {
	major := 0
	count := 0
	for _, num := range nums {
		if count == 0 {
			major = num
		}
		if num == major {
			count++
		} else {
			count--
		}
	}
	return major
}
