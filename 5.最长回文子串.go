/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {

	var maxLenStr string

	for i := 0; i < len(s); i++ {
		str := expandFromCenter(s, i, i)
		if len(str) > len(maxLenStr) {
			maxLenStr = str
		}

		str = expandFromCenter(s, i, i+1)
		if len(str) > len(maxLenStr) {
			maxLenStr = str
		}
	}

	return maxLenStr
}

func expandFromCenter(s string, a, b int) string {

	if b < len(s) && s[a] != s[b] {
		return ""
	}

	for a >= 0 && b < len(s) {

		if s[a] == s[b] {
			a--
			b++
			continue
		}

		return s[a+1 : b]
	}

	return s[a+1 : b]
}

// @lc code=end

