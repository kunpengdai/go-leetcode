/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] 阶乘后的零
 */

// @lc code=start
func trailingZeroes(n int) int {
	var res int
	for n >= 5 {
		n /= 5
		res += n
	}
	return res
}

// @lc code=end

