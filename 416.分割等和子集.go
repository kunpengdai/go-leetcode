/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

// @lc code=start
func canPartition(nums []int) bool {

	if len(nums) == 0 {
		return false
	}

	var sum int
	for _, num := range nums {
		sum += num
	}

	if sum&1 == 1 {
		return false
	}

	target := sum >> 1
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	if nums[0] > target {
		return false
	}

	if nums[0] == target {
		return true
	}

	dp := make([][]bool, len(nums))
	for i := range dp {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}
	dp[0][nums[0]] = true
	for i := 1; i < len(nums); i++ {
		dp[i][nums[i]] = true
		copy(dp[i], dp[i-1])
		for j := 0; j <= target; j++ {
			if dp[i-1][j] && j+nums[i] <= target {
				dp[i][j+nums[i]] = true
			}
		}

		if dp[i][target] {
			return true
		}
	}

	return false
}

// @lc code=end

