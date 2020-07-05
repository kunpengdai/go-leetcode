/*
 * @lc app=leetcode.cn id=713 lang=golang
 *
 * [713] 乘积小于K的子数组
 */

// @lc code=start
func numSubarrayProductLessThanK(nums []int, k int) int {

	if k <=1{
		return 0
	}
	left,prod,res := 0,1,0
	for idx,v := range nums{
		prod*=v
		for prod>= k{
			prod/=nums[left]
			left++
		}

		res += idx-left+1
	}

	return res
}
// @lc code=end

