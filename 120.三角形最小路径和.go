/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {

	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			switch {
			case j == 0:
				triangle[i][j] += triangle[i-1][j]

			case j == i:
				triangle[i][j] += triangle[i-1][j-1]

			default:
				triangle[i][j] += min(triangle[i-1][j-1], triangle[i-1][j])
			}
		}
	}

	res := math.MaxUint32
	for i := 0; i < len(triangle[len(triangle)-1]); i++ {
		res = min(res, triangle[len(triangle)-1][i])
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// @lc code=end

