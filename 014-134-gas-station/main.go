package main

import "fmt"

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
}

func canCompleteCircuit(gas []int, cost []int) int {
	start := 0
	tank := 0
	totalDiff := 0
	n := len(gas)

	for i := 0; i < n; i++ {
		diff := gas[i] - cost[i]
		totalDiff += diff
		tank += diff
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}
	if totalDiff < 0 {
		return -1
	}
	return start
}
