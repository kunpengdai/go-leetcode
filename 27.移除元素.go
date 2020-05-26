/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	var j int
	for i :=0;i<len(nums);i++{
		if nums[i]!=val{
			nums[j] = nums[i]
			j++
		}
	}
	return j	
}
// @lc code=end

