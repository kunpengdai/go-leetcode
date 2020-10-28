/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

// @lc code=start
func maxProduct(nums []int) int {

	mins := make([]int, len(nums))
	maxs := make([]int, len(nums))

	mins[0] = nums[0]
	maxs[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		maxs[i] = max(maxs[i-1]*nums[i], nums[i], mins[i-1]*nums[i])
		mins[i] = min(maxs[i-1]*nums[i], nums[i], mins[i-1]*nums[i])
	}

	res := nums[0]
	for _, num := range maxs {
		if res < num {
			res = num
		}
	}

	return res
}

func min(nums ...int) int {
	res := nums[0]
	for _, num := range nums {
		if num < res {
			res = num
		}
	}

	return res
}

func max(nums ...int) int {
	res := nums[0]
	for _, num := range nums {
		if num > res {
			res = num
		}
	}

	return res
}

// @lc code=end

