/*
 * @lc app=leetcode.cn id=540 lang=golang
 *
 * [540] 有序数组中的单一元素
 */

// @lc code=start
func singleNonDuplicate(nums []int) int {
	var res int
	for _,num := range nums {
		res ^= num
	}
	return res
}
// @lc code=end

