/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lc code=start
func maxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	maxProfitToP := make([]int, len(prices))
	maxProfitToAfter := make([]int, len(prices))

	minV := prices[0]
	for i := 1; i < len(prices); i++ {
		maxProfitToP[i] = max(prices[i]-minV, maxProfitToP[i-1])
		minV = min(minV, prices[i])
	}

	maxV := prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		maxProfitToAfter[i] = max(maxProfitToAfter[i+1], maxV-prices[i])
		maxV = max(maxV, prices[i])
	}

	res := maxProfitToP[len(prices)-1]
	for i := 0; i < len(prices)-1; i++ {
		res = max(res, maxProfitToP[i]+maxProfitToAfter[i+1])
	}

	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// @lc code=end

