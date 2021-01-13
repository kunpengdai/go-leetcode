/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	allNums := make([]int, 0, len(nums1)+len(nums2))
	var i1, i2 int
	for i1 < len(nums1) && i2 < len(nums2) {
		if nums1[i1] < nums2[i2] {
			allNums = append(allNums, nums1[i1])
			i1++
		} else {
			allNums = append(allNums, nums2[i2])
			i2++
		}
	}

	if i1 < len(nums1) {
		allNums = append(allNums, nums1[i1:]...)
	}

	if i2 < len(nums2) {
		allNums = append(allNums, nums2[i2:]...)
	}

	if len(allNums)&1 == 1 {
		return float64(allNums[len(allNums)/2])
	}

	return (float64(allNums[len(allNums)/2-1]) + float64(allNums[len(allNums)/2])) / 2
}

// @lc code=end

