/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */

// @lc code=start
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := 1; i <= amount; i++ {
			if i-coin >= 0 && dp[i-coin] > 0 {
				dp[i] += dp[i-coin]
			}
		}
	}

	return dp[amount]
}

// @lc code=end

