/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start
func isHappy(n int) bool {

	if n == 0 {
		return false
	}

	nums := make(map[int]bool)
	for n > 1 {
		square := getSquare(n)
		if nums[square]{
			return false
		}

		nums[square] = true
		n = square
	}

	return true
}

func getSquare(n int) int {
	res := 0
	for n>0 {
		tmp := n%10
		n /= 10
		res += tmp*tmp
	}
	return res
}
// @lc code=end

