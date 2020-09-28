/*
 * @lc app=leetcode.cn id=784 lang=golang
 *
 * [784] 字母大小写全排列
 */

// @lc code=start
func letterCasePermutation(S string) []string {
	res := make([]string, 0, len(S))
	bytes := make([]byte, 0, len(S))
	backTracking(S, &res, 0, &bytes)
	return res
}

func backTracking(s string, res *[]string, pos int, bytes *[]byte) {

	if pos == len(s) {
		*res = append(*res, string(*bytes))
		return
	}
	if isLetter, b := shiftLetter(s[pos]); isLetter {
		*bytes = append(*bytes, s[pos])
		backTracking(s, res, pos+1, bytes)
		(*bytes)[len(*bytes)-1] = b
		backTracking(s, res, pos+1, bytes)
		*bytes = (*bytes)[:len(*bytes)-1]

	} else {
		*bytes = append(*bytes, s[pos])
		backTracking(s, res, pos+1, bytes)
		*bytes = (*bytes)[:len(*bytes)-1]
	}
}

func shiftLetter(b byte) (bool, byte) {
	if b >= 'a' && b <= 'z' {
		return true, b + 'A' - 'a'
	}
	if b >= 'A' && b <= 'Z' {
		return true, b + 'a' - 'A'
	}
	return false, b
}

// @lc code=end

