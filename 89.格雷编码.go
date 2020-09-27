/*
 * @lc app=leetcode.cn id=89 lang=golang
 *
 * [89] 格雷编码
 */

// @lc code=start
func grayCode(n int) []int {
	if n<=0{
		return []int{0}
	}

	if n == 1 {
		return []int{0, 1}
	}

	tmp := grayCode(n - 1)
	res := append(tmp, reverseAndAdd(tmp, 1<<uint32(n-1))...)
	return res
}

func reverseAndAdd(nums []int, plus int) []int {
	res := make([]int, len(nums))
	for i := range nums {
		res[i] = nums[len(nums)-i-1] + plus
	}
	return res
}
// @lc code=end

