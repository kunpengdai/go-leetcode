/*
 * @lc app=leetcode.cn id=986 lang=golang
 *
 * [986] 区间列表的交集
 */

// @lc code=start
func intervalIntersection(A [][]int, B [][]int) [][]int {

	var (
		i, j int
		res  [][]int
	)

	for i < len(A) && j < len(B) {
		if A[i][1] < B[j][0] {
			i++
			continue
		}

		if A[i][0] > B[j][1] {
			j++
			continue
		}

		res = append(res, []int{max(A[i][0], B[j][0]), min(A[i][1], B[j][1])})
		if A[i][1] > B[j][1] {
			j++
		} else {
			i++
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

