/*
 * @lc app=leetcode.cn id=403 lang=golang
 *
 * [403] 青蛙过河
 */

// @lc code=start
func canCross(stones []int) bool {

	if stones[1] != 1 {
		return false
	}

	stoneMap := make(map[int]bool)
	for _, stone := range stones {
		stoneMap[stone] = true
	}

	stoneStep := make(map[int][]int)
	stoneStep[1] = []int{1}
	for i := 1; i < len(stones); i++ {

		if len(stoneStep[stones[i]]) == 0 {
			continue
		}

		for _, step := range stoneStep[stones[i]] {
			for _, pstep := range []int{step - 1, step, step + 1} {
				if stoneMap[stones[i]+pstep] {
					stoneStep[stones[i]+pstep] = append(stoneStep[stones[i]+pstep], pstep)
				}
			}
		}

		if len(stoneStep[stones[len(stones)-1]]) != 0 {
			return true
		}
	}

	return false
}

// @lc code=end

