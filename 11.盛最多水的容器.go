/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	left, right, res := 0, len(height)-1, 0
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		if res < area {
			res = area
		}

		if height[left] < height[right] {
			tmp := left
			for left < right {
				left++
				if height[left] > height[tmp] {
					break
				}
			}
		} else {
			tmp := right
			for left < right {
				right--
				if height[right] > height[tmp] {
					break
				}
			}
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

// @lc code=end

