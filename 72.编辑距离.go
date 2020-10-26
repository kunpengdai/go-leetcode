/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {

	dp := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		dp[i] = make([]int, len(word2)+1)
	}

	for i := 1; i <= len(word1); i++ {
		dp[i][0] = i
	}

	for j := 1; j <= len(word2); j++ {
		dp[0][j] = j
	}

	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

func min(vs ...int) int {
	res := math.MaxInt32
	for _, num := range vs {
		if num < res {
			res = num
		}
	}

	return res
}

// @lc code=end

