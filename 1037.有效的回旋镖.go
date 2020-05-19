/*
 * @lc app=leetcode.cn id=1037 lang=golang
 *
 * [1037] 有效的回旋镖
 */

// @lc code=start
func isBoomerang(points [][]int) bool {
	return  (points[0][0]-points[1][0])*(points[0][1]-points[2][1]) != 
		(points[0][0]-points[2][0])*(points[0][1]-points[1][1])
}
// @lc code=end

