/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start
func numDecodings(s string) int {

	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1
	dp[1] = 1

	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			if s[i-1] < '1' || s[i-1] > '2' {
				return 0
			}
			dp[i+1] = dp[i-1]
			continue
		}

		if s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6') {
			dp[i+1] = dp[i] + dp[i-1]
		} else {
			dp[i+1] = dp[i]
		}
	}

	return dp[len(s)]
}

// @lc code=end

