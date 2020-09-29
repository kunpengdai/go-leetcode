/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */

// @lc code=start

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	numToGreater := make(map[int]int, len(nums2))
	stack := make([]int, 0, len(nums2)/2)
	for _, num := range nums2 {
		if len(stack) == 0 {
			stack = append(stack, num)
		} else {
			for len(stack) != 0 {
				lastNum := stack[len(stack)-1]
				if lastNum < num {
					stack = stack[:len(stack)-1]
					numToGreater[lastNum] = num
				} else {
					break
				}
			}
			stack = append(stack, num)
		}
	}
	for _, num := range stack {
		numToGreater[num] = -1
	}

	res := make([]int, len(nums1))
	for i, num := range nums1 {
		res[i] = numToGreater[num]
	}

	return res
}

// @lc code=end

