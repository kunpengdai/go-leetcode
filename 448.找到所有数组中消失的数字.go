/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {

	hasSeen := make(map[int]bool, len(nums))
	for _, num := range nums {
		hasSeen[num] = true
	}

	var res []int
	for i := range nums {
		if !hasSeen[i+1] {
			res = append(res, i+1)
		}
	}

	return res
}

// @lc code=end

