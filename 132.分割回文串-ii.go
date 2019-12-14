/*
 * @lc app=leetcode.cn id=132 lang=golang
 *
 * [132] 分割回文串 II
 */

// @lc code=start
func minCut(s string) int {

	length := len(s)
	if length < 2 {
		return 0
	}

	valid := make([][]bool, len(s))
	for i := 0; i < length; i++ {
		valid[i] = make([]bool, length)
	}
	for i := 0; i < length; i++ {
		valid[i][i] = true
	}

	for i := length - 2; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			if s[i] == s[j] {
				if j-i == 1 || valid[i+1][j-1] {
					valid[i][j] = true
				}
			}
		}
	}

	dp := make([]int, length)
	for i := 0; i < length; i++ {
		dp[i] = length
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 0; i < length; i++ {
		if valid[0][i] {
			dp[i] = 0
			continue
		}
		for j := 0; j < i; j++ {
			if valid[j+1][i] {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}

	return dp[length-1]
}

// @lc code=end

