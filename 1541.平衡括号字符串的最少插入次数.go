/*
 * @lc app=leetcode.cn id=1541 lang=golang
 *
 * [1541] 平衡括号字符串的最少插入次数
 */

// @lc code=start

func minInsertions(s string) int {
	var res, left int
	for i := 0; i < len(s); i++ {
		b := s[i]
		switch b {
		case '(':
			left++
		case ')':
			if left == 0 {
				res++
				left++
			}
			if i == len(s)-1 || s[i+1] == '(' {
				left--
				res++
			} else {
				i++
				left--
			}
		}
	}

	res += left * 2

	return res
}
// @lc code=end

