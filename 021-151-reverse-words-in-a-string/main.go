package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "the sky is blue"
	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {
	var result strings.Builder
	i := len(s) - 1
	for i >= 0 {
		for i >= 0 && s[i] == ' ' {
			i--
		}
		if i < 0 {
			break
		}
		right := i
		for i >= 0 && s[i] != ' ' {
			i--
		}
		left := i + 1
		if result.Len() > 0 {
			result.WriteByte(' ')
		}
		result.WriteString(s[left : right+1])
	}

	return result.String()
}
