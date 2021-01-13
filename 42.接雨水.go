/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) (res int) {

	leftMaxHeight := make([]int, len(height))
	rightMaxHeight := make([]int, len(height))

	var maxHeight int
	for i := 0; i < len(height); i++ {
		if height[i] > maxHeight {
			maxHeight = height[i]
		}

		leftMaxHeight[i] = maxHeight
	}

	maxHeight = 0
	for i := len(height) - 1; i >= 0; i-- {
		if height[i] > maxHeight {
			maxHeight = height[i]
		}

		rightMaxHeight[i] = maxHeight
	}

	for i, v := range height {
		if ans := min(leftMaxHeight[i], rightMaxHeight[i]); ans > v {
			res += (ans - v)
		}
	}

	return 
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// @lc code=end

