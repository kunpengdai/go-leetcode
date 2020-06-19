/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除排序数组中的重复项 II
 */

// @lc code=start
func removeDuplicates(nums []int) int {

	i := 2
	length := len(nums)
	for ;i< length;i++{
		if nums[i]==nums[i-2]{
			nums = append(nums[:i],nums[i+1:]...)
			length--
			i--
		}
	}

	return length
}
// @lc code=end

