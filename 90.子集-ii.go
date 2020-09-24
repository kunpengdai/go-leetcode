/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	res := make([][]int, 0, 1<<uint64(len(nums)))
	dfs(nums, 0, &res, []int{})
	return res
}

func dfs(nums []int, pos int, res *[][]int, tmpSlice []int) {
	if pos == len(nums) {
		*res = append(*res, append([]int{}, tmpSlice...))
		return
	}

	dfs(nums, pos+1, res, tmpSlice)

	if pos > 0 && nums[pos] == nums[pos-1] {
		var n1, n2 int
		for i := pos - 1; i >= 0; i-- {
			if nums[i] != nums[pos] {
				break
			}
			n1++
		}

		for i := len(tmpSlice) - 1; i >= 0; i-- {
			if tmpSlice[i] != nums[pos] {
				break
			}
			n2++
		}

		if n1 != n2 {
			return
		}
	}

	tmpSlice = append(tmpSlice, nums[pos])
	dfs(nums, pos+1, res, tmpSlice)
}

// @lc code=end

