
/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	if len(haystack)<len(needle){
		return -1
	}
	if needle==""{
		return 0
	}
	for i := 0;i <= len(haystack)-len(needle);i++{
		j := 0
		for ;j<len(needle);j++{
			if haystack[i+j]!=needle[j]{
				break
			}
		}
		if j == len(needle){
			return i
		}
	}

	return -1
}

// @lc code=end