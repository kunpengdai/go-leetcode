/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int)  {

	p0,curr := 0,0
	p2 := len(nums)-1

	for curr<=p2{
		if nums[curr]==0{
			nums[curr],nums[p0] = nums[p0],nums[curr]
			p0++
			curr++
		}else if nums[curr]==2{
			nums[curr],nums[p2]=nums[p2],nums[curr]
			p2--
		}else{
			curr++
		}
	}
}
// @lc code=end

