/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	min := prices[0]
	var res int
	for _, price := range prices {
		if res < price-min {
			res = price - min
		}

		if price < min {
			min = price
		}
	}

	return res
}

// @lc code=end

