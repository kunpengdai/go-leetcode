/*
 * @lc app=leetcode.cn id=424 lang=golang
 *
 * [424] 替换后的最长重复字符
 */

// @lc code=start
func characterReplacement(s string, k int) int {
	var l, maxChar int
	var m [26]int


	for r:=0;r<len(s);r++{
		m[int(s[r]-'A')]++
		if maxChar < m[int(s[r]-'A')]{
			maxChar =m[int(s[r]-'A')] 
		}

		if r-l+1>k+maxChar{
			m[int(s[l]-'A')]--	
			l++		
		}
	}

	return len(s)-l
}

// @lc code=end

