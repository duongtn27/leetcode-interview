package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	i := 0
	j := n - 1
	for i < j {
		total := numbers[i] + numbers[j]
		if total == target {
			return []int{i + 1, j + 1}
		}
		if total < target {
			i++
		} else {
			j--
		}
	}
	return nil
}
