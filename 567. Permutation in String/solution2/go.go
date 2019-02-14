package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("ab", "acagbvacbcva"))
}

func checkInclusion(s1 string, s2 string) bool {
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)
	if len(s1) > len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		m1[s1[i]]++
	}
	for i := 0; i < len(s1); i++ {
		m2[s2[i]]++
	}
	if check(m1, m2) {
		return true
	}
	for i := len(s1); i < len(s2); i++ {
		m2[s2[i]]++
		m2[s2[i-len(s1)]]--
		if check(m1, m2) {
			return true
		}
	}
	return false
}

func check(m1, m2 map[byte]int) bool {
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}
