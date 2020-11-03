/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	m := make(map[int]bool, len(s))
	startPos := make([]int, 0, len(s))
	startPos = append(startPos, 0)

	for len(startPos) != 0 {
		pos := startPos[0]
		startPos = startPos[1:]

		for _, word := range wordDict {
			if pos+len(word) <= len(s) && word == s[pos:pos+len(word)] {
				if pos+len(word) == len(s) {
					return true
				} else {
					if !m[pos+len(word)] {
						startPos = append(startPos, pos+len(word))
						m[pos+len(word)] = true
					}
				}
			}
		}
	}

	return false
}

// @lc code=end

