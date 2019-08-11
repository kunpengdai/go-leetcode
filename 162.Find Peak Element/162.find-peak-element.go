/*
 * @lc app=leetcode id=162 lang=golang
 *
 * [162] Find Peak Element
 */
func findPeakElement(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, beg, end int) int {
	mid := (beg + end) >> 1
	if mid == 0 {
		if nums[0] > nums[1] {
			return 0
		}
		return helper(nums, 1, end)
	}
	if mid == len(nums)-1 {
		if nums[mid] > nums[mid-1] {
			return len(nums) - 1
		}
		return helper(nums, beg, mid-1)
	}
	if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
		return mid
	}
	if nums[mid] < nums[mid+1] {
		return helper(nums, mid+1, end)
	}
	return helper(nums, beg, mid-1)

}

