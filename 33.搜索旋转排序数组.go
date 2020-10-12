/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {

	if len(nums) == 0 {
		return -1
	}

	var low, high int
	minIndex := findMinIndex(nums)
	switch {
	case nums[minIndex] == target:
		return minIndex
	case nums[minIndex] < target:
		if nums[len(nums)-1] < target {
			high = minIndex - 1
		} else if nums[len(nums)-1] >= target {
			low = minIndex + 1
			high = len(nums) - 1
		}
	case nums[minIndex] > target:
		return -1
	}

	for low < high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if nums[low] == target {
		return low
	}

	return -1
}

func findMinIndex(nums []int) int {
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

	return low
}

// @lc code=end

