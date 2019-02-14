package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("ab", "abeidboaooa"))
}

func check(s1, s2 string) bool {
	m1 := make(map[rune]int)
	m2 := make(map[rune]int)
	for _, r := range s1 {
		m1[r]++
	}
	for _, r := range s2 {
		m2[r]++
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}

func checkInclusion(s1 string, s2 string) bool {
	for i := 0; i <= len(s2)-len(s1); i++ {
		if check(s1, s2[i:i+len(s1)]) {
			return true
		}
	}
	return false
}
