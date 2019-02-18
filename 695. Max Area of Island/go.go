package main

import "fmt"

func main() {
	grid := [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
	fmt.Println(maxAreaOfIsland(grid))
}

func maxAreaOfIsland(grid [][]int) int {
	used := make([][]bool, 0)
	for i := 0; i < len(grid); i++ {
		used = append(used, make([]bool, len(grid[i])))
	}
	maxArea := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if !used[i][j] && grid[i][j] == 1 {
				area := 0
				dfs(grid, used, i, j, &area)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}

func dfs(grid [][]int, used [][]bool, i, j int, area *int) {
	used[i][j] = true
	*area++

	if i > 0 && !used[i-1][j] && grid[i-1][j] == 1 {
		dfs(grid, used, i-1, j, area)
	}
	if j > 0 && !used[i][j-1] && grid[i][j-1] == 1 {
		dfs(grid, used, i, j-1, area)
	}
	if i < len(grid)-1 && !used[i+1][j] && grid[i+1][j] == 1 {
		dfs(grid, used, i+1, j, area)
	}
	if j < len(grid[i])-1 && !used[i][j+1] && grid[i][j+1] == 1 {
		dfs(grid, used, i, j+1, area)
	}
}
