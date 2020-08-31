/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] 数组中重复的数据
 */

// @lc code=start
func findDuplicates(nums []int) []int {
	res := make([]int,0)
	m := make(map[int]bool, len(nums))
	for _,num := range nums {
		if m[num] {
			res = append(res,num)
		}
		m[num] = true
	}
	return res
}
// @lc code=end

