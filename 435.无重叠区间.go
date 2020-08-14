/*
 * @lc app=leetcode.cn id=435 lang=golang
 *
 * [435] 无重叠区间
 */

// @lc code=start
func eraseOverlapIntervals(intervals [][]int) int {

	sort.Slice(intervals, func(i int, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})

	var res int
	for i := 0; i < len(intervals)-1; i++ {

		if intervals[i+1][0] >= intervals[i][1] {
			continue
		}

		res++
		if intervals[i][1] < intervals[i+1][1] {
			intervals[i+1] = intervals[i]
		}
	}

	return res
}
// @lc code=end

