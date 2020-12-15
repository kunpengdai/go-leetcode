/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
 */

// @lc code=start
func maxProfit(prices []int, fee int) int {

	if len(prices) == 0 {
		return 0
	}

	dp1 := make([]int, len(prices)) // i时持有股票拥有的最大收益
	dp2 := make([]int, len(prices)) // i时不持有股票拥有的最大收益

	dp1[0] -= prices[0]
	for i := 1; i < len(prices); i++ {
		dp1[i] = max(dp1[i-1], dp2[i-1]-prices[i])
		dp2[i] = max(dp2[i-1], dp1[i-1]+prices[i]-fee)
	}

	return dp2[len(prices)-1]
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

// @lc code=end

