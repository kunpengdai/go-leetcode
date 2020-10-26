/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长上升子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {

	dp := make([]int, len(nums))
	for i := range nums {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}

	var res int
	for _, v := range dp {
		if v > res {
			res = v
		}
	}

	return res
}

// @lc code=end

