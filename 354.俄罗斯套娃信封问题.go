/*
 * @lc app=leetcode.cn id=354 lang=golang
 *
 * [354] 俄罗斯套娃信封问题
 */

// @lc code=start
func maxEnvelopes(envelopes [][]int) int {

	if len(envelopes) <= 1 {
		return len(envelopes)
	}

	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		}
		return envelopes[i][1] > envelopes[j][1]
	})

	dp := make([]int, len(envelopes))
	for i := range dp {
		dp[i] = 1
	}

	res := 1
	for i := 1; i < len(envelopes); i++ {
		for j := 0; j < i; j++ {
			if envelopes[i][1] > envelopes[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
				res = max(res, dp[i])
			}
		}
	}

	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

