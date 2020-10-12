/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	end := nums[len(nums)-1]
	low, high := 0, len(nums)-1
	for low < high {
		mid := (low + high) / 2
		switch {
		case nums[mid] >= end:
			low = mid + 1
		case nums[mid] < end:
			high = mid
		}
	}

	return nums[low]
}
// @lc code=end

