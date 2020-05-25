/*
 * @lc app=leetcode.cn id=1034 lang=golang
 *
 * [1034] 边框着色
 */

// @lc code=start
func colorBorder(grid [][]int, r0 int, c0 int, color int) [][]int {
	if len(grid) == 0 {
		return grid
	}

	corner := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		corner[i] = make([]bool, len(grid[0]))
	}

	colorGrid(grid, &corner, r0, c0, grid[r0][c0], color)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if corner[i][j] && (j == 0 || j == len(grid[0])-1 || !corner[i][j+1] || !corner[i][j-1]) {
				grid[i][j] = color
			}
		}
	}

	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); j++ {
			if corner[j][i] && (j == 0 || j == len(grid)-1 || !corner[j+1][i] || !corner[j-1][i]) {
				grid[j][i] = color
			}
		}
	}

	return grid
}

func colorGrid(grid [][]int, corner *[][]bool, i, j, originColor, color int) {

	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return
	}

	ds := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	if originColor != grid[i][j] || (*corner)[i][j] {
		return
	}

	(*corner)[i][j] = true

	for p := 0; p < 4; p++ {
		ix := i + ds[p][0]
		iy := j + ds[p][1]
		colorGrid(grid, corner, ix, iy, originColor, color)
	}
}

// @lc code=end