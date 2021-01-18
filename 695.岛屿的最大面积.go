/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 */

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
	checked := make([][]bool, len(grid))
	for i := range checked {
		checked[i] = make([]bool, len(grid[0]))
	}

	var res int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			var area int
			dfs(grid, &checked, i, j, &area)
			if area > res {
				res = area
			}
		}
	}

	return res
}

func dfs(grid [][]int, checked *[][]bool, i, j int, area *int) {

	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return
	}

	if (*checked)[i][j] {
		return
	}

	if grid[i][j] == 0 {
		return
	}

	(*area)++
	(*checked)[i][j] = true
	for _, direction := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		dfs(grid, checked, i+direction[0], j+direction[1], area)
	}
}

// @lc code=end

