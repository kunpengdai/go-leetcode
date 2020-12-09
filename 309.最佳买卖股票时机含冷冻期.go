/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */

// @lc code=start
func maxProfit(prices []int) int {

	if len(prices) <= 1 {
		return 0
	}

	dp1 := make([]int, len(prices)) // i时未持有股票后最多剩余的钱
	dp2 := make([]int, len(prices)) // i时持有股票最多剩余的钱

	dp2[0] -= prices[0]
	dp2[1] = max(-prices[0], -prices[1])
	dp1[1] = max(0, prices[1]-prices[0])
	for i := 2; i < len(prices); i++ {
		dp1[i] = max(dp1[i-1], dp2[i-1]+prices[i])
		dp2[i] = max(dp2[i-1], dp1[i-2]-prices[i])
	}

	return dp1[len(prices)-1]
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

// @lc code=end

