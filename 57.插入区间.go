/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	return merge(intervals)
}

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

