/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
type numTimes struct {
	num   int
	times int
}

func topKFrequent(nums []int, k int) []int {
	num2Times := make(map[int]int, len(nums))
	for _, num := range nums {
		num2Times[num]++
	}

	nts := make([]numTimes, 0, len(num2Times))
	for num, times := range num2Times {
		nts = append(nts, numTimes{
			num:   num,
			times: times,
		})
	}

	sort.Slice(nts, func(i, j int) bool {
		return nts[i].times > nts[j].times
	})
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = nts[i].num
	}

	return res
}

// @lc code=end

