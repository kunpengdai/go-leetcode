/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	return commenPrefix(strs[0], longestCommonPrefix(strs[1:]))
}

func commenPrefix(str1, str2 string) string {
	var i int
	for ; i < len(str1) && i < len(str2); i++ {
		if str1[i] != str2[i] {
			break
		}
	}

	return str1[:i]
}

// @lc code=end

