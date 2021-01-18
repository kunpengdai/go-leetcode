/*
 * @lc app=leetcode.cn id=611 lang=golang
 *
 * [611] 有效三角形的个数
 */

// @lc code=start
func triangleNumber(nums []int) int {
	sort.Ints(nums)
	var res int
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			low, high := j+1, len(nums)-1
			for low < high {
				mid := (low + high) / 2
				switch {
				case nums[mid] >= nums[i]+nums[j]:
					high = mid

				default:
					low = mid + 1
				}
			}

			// 找到的是第一个可能的大于双边和的位置，如果确实大于，需要回退一个位置
			if nums[high] >= nums[i]+nums[j] {
				high--
			}

			res += (high - j)
		}
	}

	return res
}

// @lc code=end

