/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
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

	for i, num := range nums {
		if !(*used)[i] {
			sl = append(sl, num)
			(*used)[i] = true
			backTracking(res, used, nums, sl)
			sl = sl[:len(sl)-1]
			(*used)[i] = false
		}
	}
}


func permute1(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return [][]int{
			{nums[0]},
		}
	}

	tmp := permute(nums[:len(nums)-1])
	res := make([][]int, len(nums)*len(tmp))
	for i, sl := range tmp {
		for j := 0; j < len(nums); j++ {
			res[i*(len(nums))+j] = insert(sl, nums[len(nums)-1], j)
		}
	}

	return res
}

func insert(nums []int, num, pos int) []int {
	res := make([]int, len(nums))
	copy(res, nums)
	switch pos {
	case 0:
		res = append([]int{num}, res...)
	case len(nums):
		res = append(res, num)
	default:
		res = append(res[:pos], append([]int{num}, res[pos:]...)...)
	}
	return res
}
// @lc code=end

