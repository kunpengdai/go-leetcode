/*
 * @lc app=leetcode.cn id=228 lang=golang
 *
 * [228] 汇总区间
 */

// @lc code=start
type pair struct {
	start, end int
}

func summaryRanges(nums []int) []string {

	if len(nums) == 0 {
		return nil
	}

	pairs := make([]pair, 1)
	pairs[0] = pair{
		start: nums[0],
		end:   nums[0],
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] == pairs[len(pairs)-1].end+1 {
			pairs[len(pairs)-1].end++
		} else {
			pairs = append(pairs, pair{
				start: nums[i],
				end:   nums[i],
			})
		}
	}

	res := make([]string, len(pairs))
	for i, p := range pairs {
		if p.start == p.end {
			res[i] = fmt.Sprintf("%d", p.start)
		} else {
			res[i] = fmt.Sprintf("%d->%d", p.start, p.end)
		}
	}

	return res
}

// @lc code=end

