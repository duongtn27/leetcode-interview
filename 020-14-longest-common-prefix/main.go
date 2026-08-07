package main

import "fmt"

func main() {
	strs := []string{"ab", "a"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for j := 0; j < min(len(prefix), len(strs[i])); j++ {
			if prefix[j] != strs[i][j] {
				prefix = prefix[:j]
				break
			}
		}
		if len(strs[i]) < len(prefix) {
			prefix = prefix[:len(strs[i])]
		}

		if prefix == "" {
			return ""
		}
	}
	return prefix
}
