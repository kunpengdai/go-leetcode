/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 */

// @lc code=start
func findPeakElement(nums []int) int {
	return findPeakElementHelper(nums, 0, len(nums)-1)
}

func findPeakElementHelper(nums []int, from, end int) int {

	if from == end {
		return from
	}

	mid := (from + end) / 2
	if nums[mid] > nums[mid+1] {
		return findPeakElementHelper(nums, from, mid)
	}

	return findPeakElementHelper(nums, mid+1, end)
}

// @lc code=end

