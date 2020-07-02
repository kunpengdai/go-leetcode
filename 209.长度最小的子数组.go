/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(s int, nums []int) int {
	if len(nums)==0{
		return 0
	}

	left,res,sum := 0,math.MaxInt32,0
	for i:=0;i<len(nums);i++{
		sum += nums[i]
		for sum >= s {
			if res > i-left+1{
				res = i-left+1
			}
			sum -= nums[left]
			left++
		}
	} 

	if res == math.MaxInt32{
		return 0
	}

	return res
}
// @lc code=end

