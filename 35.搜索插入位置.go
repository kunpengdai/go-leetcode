/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	if nums[len(nums)-1] < target {
		return len(nums)
	}

	low, high := 0, len(nums)-1
	for low < high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}

// @lc code=end

