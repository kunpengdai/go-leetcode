package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	start := 0
	res := 0
	for index, ch := range s {
		if _, ok := m[ch]; !ok || start > m[ch] {
			if index-start+1 > res {
				res = index - start + 1
			}
		} else {
			start = m[ch] + 1
		}
		m[ch] = index
	}
	return res
}
