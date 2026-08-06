package main

import "fmt"

func main() {
	s := "MCMXCIV"
	fmt.Println(intToRoman(s))
}

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func intToRoman(s string) int {
	left := 0
	right := 0
	n := len(s)
	total := roman[string(s[n-1])]
	for i := n - 2; i >= 0; i-- {
		left = roman[string(s[i])]
		right = roman[string(s[i+1])]
		if left < right {
			total -= left
		} else {
			total += left
		}
	}
	return total
}
