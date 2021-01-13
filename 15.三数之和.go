/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	allRes := make([][]int, 0)
	for start := 0; start < len(nums)-2; start = findNextDifferentNum(nums, start) {
		l, r := start+1, len(nums)-1
		for l < r {
			sum := nums[start] + nums[l] + nums[r]
			switch {
			case sum == 0:
				allRes = append(allRes, []int{nums[start], nums[l], nums[r]})
				l = findNextDifferentNum(nums, l)
				r = findPrevDifferentNum(nums, r)
			case sum < 0:
				l = findNextDifferentNum(nums, l)
			case sum > 0:
				r = findPrevDifferentNum(nums, r)
			}
		}
	}

	return allRes
}

func findNextDifferentNum(nums []int, pos int) int {
	for i := pos + 1; i < len(nums); i++ {
		if nums[i] != nums[pos] {
			return i
		}
	}

	return len(nums)
}

func findPrevDifferentNum(nums []int, pos int) int {

	for i := pos - 1; i >= 0; i-- {
		if nums[i] != nums[pos] {
			return i
		}
	}

	return -1
}

// @lc code=end

