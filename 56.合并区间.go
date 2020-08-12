/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {

	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})

	var res [][]int
	tmp := intervals[0]
	for _, item := range intervals {

		if item[0] <= tmp[1] {
			if item[1] > tmp[1] {
				tmp[1] = item[1]
			}
			continue
		}

		res = append(res, tmp)
		tmp = item
	}

	res = append(res, tmp)
	return res
}
// @lc code=end

