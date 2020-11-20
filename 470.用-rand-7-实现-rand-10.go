/*
 * @lc app=leetcode.cn id=470 lang=golang
 *
 * [470] 用 Rand7() 实现 Rand10()
 */

// @lc code=start
func rand10() int {
	for {
		col := rand7()
		row := rand7()
		res := col + (row-1)*7
		if res <= 40 {
			return 1 + (res-1)%10
		}
	}

	return -1
}

// @lc code=end

