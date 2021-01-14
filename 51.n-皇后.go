/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */

// @lc code=start

func conflict(res []int, row, pos int) bool {
	for i := 0; i < row; i++ {
		if pos == res[i] || (row-i) == abs(pos-res[i]) {
			return true
		}
	}

	return false
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func solveNQueens(n int) [][]string {
	init := make([]int, n)
	res := make([][]string, 0)
	dfs(&res, init, n, 0, 0)
	return res
}

func dfs(res *[][]string, poses []int, n, row, pos int) {

	if row == n {
		*res = append(*res, genRes(poses, n))
		return
	}

	if pos == n {
		return
	}

	if !conflict(poses, row, pos) {
		poses[row] = pos
		dfs(res, poses, n, row+1, 0)
	}

	dfs(res, poses, n, row, pos+1)
}

func genRes(poses []int, n int) []string {

	res := make([]string, n)
	for i := 0; i < n; i++ {
		var b strings.Builder
		for j := 0; j < n; j++ {
			if j != poses[i] {
				b.WriteByte('.')
			} else {
				b.WriteByte('Q')
			}
		}
		res[i] = b.String()
	}

	return res
}

// @lc code=end

