package main

import "fmt"

func main() {
	haystack := "sabutsad"
	needle := "sad"
	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	n := len(needle)
	for i := 0; i+n <= len(haystack); i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}
	return -1
}
