/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {

	if amount == 0 {
		return 0
	}

	dp := make([]int, amount+1)
	for _, coin := range coins {
		if coin <= amount {
			dp[coin] = 1
		}
	}

	for i := 1; i <= amount; i++ {
		if dp[i] == 0 {
			for _, coin := range coins {
				if i-coin > 0 && dp[i-coin] != 0 {
					dp[i] = min(dp[i], dp[i-coin]+1)
				}
			}
		}
	}

	if dp[amount] == 0 {
		return -1
	}

	return dp[amount]
}

func min(a, b int) int {

	if a == 0 {
		return b
	}

	if a < b {
		return a
	}
	return b
}

// @lc code=end

