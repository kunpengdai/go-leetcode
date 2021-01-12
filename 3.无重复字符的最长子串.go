/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	var res, temp int
	lastPos := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		if pos, ok := lastPos[s[i]]; !ok {
			temp++
			if temp > res {
				res = temp
			}
		} else {
			temp = i - pos
			lastPos = make(map[byte]int, temp)
			for j := pos + 1; j <= i; j++ {
				lastPos[s[j]] = j
			}
		}

		lastPos[s[i]] = i
	}

	return res
}

// @lc code=end

