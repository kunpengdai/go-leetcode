/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	res := make([][]int, 0, len(nums))
	used := make([]bool, len(nums))
	sl := make([]int, 0, len(nums))
	backTracking(&res, &used, nums, sl)
	return res
}

func backTracking(res *[][]int, used *[]bool, nums, sl []int) {
	if len(sl) == len(nums) {
		*res = append(*res, append([]int{}, sl...))
		return
	}

	permuteJustNow := -1
	for i, num := range nums {
		if !(*used)[i] {

			if permuteJustNow != -1 && nums[i] == nums[permuteJustNow] {
				continue
			}

			permuteJustNow = i
			sl = append(sl, num)
			(*used)[i] = true
			backTracking(res, used, nums, sl)
			sl = sl[:len(sl)-1]
			(*used)[i] = false
		}
	}
}
// @lc code=end

