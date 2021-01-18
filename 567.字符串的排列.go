/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {

	if len(s1) > len(s2) {
		return false
	}

	m1 := make(map[byte]int, len(s1))
	m2 := make(map[byte]int, len(s1))
	for i := 0; i < len(s1); i++ {
		m1[s1[i]]++
		m2[s2[i]]++
	}

	if equals(m1, m2) {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		m2[s2[i-len(s1)]]--
		m2[s2[i]]++
		if m2[s2[i-len(s1)]] == 0 {
			delete(m2, s2[i-len(s1)])
		}

		if equals(m1, m2) {
			return true
		}
	}

	return false
}

func equals(m1, m2 map[byte]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k := range m1 {
		if m1[k] != m2[k] {
			return false
		}
	}

	return true
}

// @lc code=end

