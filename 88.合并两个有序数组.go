/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
	p1,p2,curr,tmp := m-1,n-1,len(nums1)-1,0
	for p1>=0 && p2>=0{
		if nums1[p1] > nums2[p2]{
			tmp = nums1[p1]
			p1--
		}else{
			tmp = nums2[p2]
			p2--
		}

		nums1[curr] = tmp
		curr--
	}

	if p2>=0{
		copy(nums1[:p2+1],nums2[:p2+1])
	}
}
// @lc code=end

