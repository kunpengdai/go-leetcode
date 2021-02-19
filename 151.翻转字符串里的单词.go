/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 翻转字符串里的单词
 */

// @lc code=start
func reverseWords(s string) string {
	words := strings.Split(s, " ")
	reversedWords := make([]string, 0, len(words))
	for i := len(words) - 1; i >= 0; i-- {
		if words[i] != "" {
			reversedWords = append(reversedWords, words[i])
		}
	}
	return strings.Join(reversedWords, " ")
}

// @lc code=end

