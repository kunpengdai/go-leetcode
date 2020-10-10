/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	return []int{findLowBound(nums, target), findHighBound(nums, target)}
}

func findLowBound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	low, high := 0, len(nums)-1
	for low < high {
		if nums[low] == target {
			return low
		}
		mid := (low + high) / 2
		switch {
		case nums[mid] >= target:
			high = mid
		case nums[mid] < target:
			low = mid + 1
		}
	}

	if nums[low] == target {
		return low
	}

	return -1
}

func findHighBound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	low, high := 0, len(nums)-1
	for low < high {
		if nums[high] == target {
			return high
		}
		mid := (low + high + 1) / 2
		switch {
		case nums[mid] > target:
			high = mid - 1
		case nums[mid] <= target:
			low = mid
		}
	}

	if nums[high] == target {
		return high
	}

	return -1
}
// @lc code=end

