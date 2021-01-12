/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	m := make(map[int]bool,len(nums))
	for _,num := range nums {
		if m[num] {
			return true
		}
		m[num] = true
	}

	return false
}
// @lc code=end

