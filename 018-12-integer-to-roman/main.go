package main

import (
	"fmt"
	"strings"
)

func main() {
	num := 3749
	fmt.Println(intToRoman(num))
}

func intToRoman(num int) string {
	values := []int{
		1000, 900, 500, 400,
		100, 90, 50, 40,
		10, 9, 5, 4, 1,
	}

	romans := []string{
		"M", "CM", "D", "CD",
		"C", "XC", "L", "XL",
		"X", "IX", "V", "IV", "I",
	}
	var result strings.Builder
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			result.WriteString(romans[i])
			num -= values[i]
		}
	}
	return result.String()
}
