/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	res := nums[0]
	var sum int
	for _, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}

		if sum > res {
			res = sum
		}
	}

	return res
}

// @lc code=end

