/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 */

// @lc code=start
func isUgly(num int) bool {

	if num <= 0 {
		return false
	}

	nums := []int{2, 3, 5}
	for _, n := range nums {
		for num%n == 0 {
			num /= n
		}
	}
	return num == 1
}

// @lc code=end

