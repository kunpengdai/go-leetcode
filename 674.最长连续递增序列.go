/*
 * @lc app=leetcode.cn id=674 lang=golang
 *
 * [674] 最长连续递增序列
 */

// @lc code=start
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res, temp := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			temp++
			if temp > res {
				res = temp
			}
		} else {
			temp = 1
		}
	}

	return res
}

// @lc code=end

