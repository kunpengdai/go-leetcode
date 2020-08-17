/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 缺失数字
 */

// @lc code=start
func missingNumber(nums []int) int {
	var sum, total int
	for i, num := range nums {
		if sum > i {
			sum += num - i
		} else {
			sum += num
			total += i
		}
	}
	total += len(nums)

	return total - sum
}
// @lc code=end

