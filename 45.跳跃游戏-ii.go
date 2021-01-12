/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {

	if len(nums) <= 1 {
		return 0
	}

	steps := make([]int, 1, len(nums))
	steps[0] = nums[0]
	for steps[len(steps)-1] < len(nums)-1 {
		lastMaxPos := steps[len(steps)-1]
		maxPos := lastMaxPos
		var preMaxPos int
		if len(steps) > 1 {
			preMaxPos = steps[len(steps)-2]
		}
		for i := preMaxPos; i <= lastMaxPos; i++ {
			maxPos = max(maxPos, i+nums[i])
		}

		steps = append(steps, maxPos)
	}

	return len(steps)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

