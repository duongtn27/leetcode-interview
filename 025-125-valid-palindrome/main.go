package main

import (
	"fmt"
)

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		for i < j && !isAlphaNum(s[i]) {
			i++
		}

		for i < j && !isAlphaNum(s[j]) {
			j--
		}

		if toLower(s[i]) != toLower(s[j]) {
			return false
		}

		i++
		j--
	}
	return true
}

func isAlphaNum(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}
