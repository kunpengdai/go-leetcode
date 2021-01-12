/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 */

// @lc code=start
func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	for i := range ss {
		ss[i] = reverseString(ss[i])
	}

	return strings.Join(ss, " ")
}

func reverseString(s string) string {
	rs := []rune(s)
	for i := 0; i < len(rs)/2; i++ {
		rs[i], rs[len(rs)-i-1] = rs[len(rs)-i-1], rs[i]
	}
	return string(rs)
}

// @lc code=end

