/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) / 2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] > target:
			high = mid - 1
		case nums[mid] < target:
			low = mid + 1
		}
	}

	return -1
}
// @lc code=end

