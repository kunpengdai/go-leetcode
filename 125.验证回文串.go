/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
func isPalindrome(s string) bool {
	s = strings.ToUpper(s)
	i,j := 0,len(s)-1
	for i<j{
		if !isLetterOrDigit(s[i]){
			i++
			continue
		}
		if !isLetterOrDigit((s[j])){
			j--
			continue
		}
		if s[i]!=s[j]{
			return false
		}
		i++
		j--
	}
	return true
}

func isLetterOrDigit(b byte)bool{
	return (b>='A'&& b<='Z') || (b>='0' && b<='9') 
}
// @lc code=end

