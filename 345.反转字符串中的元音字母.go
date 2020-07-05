/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 */

// @lc code=start
func reverseVowels(s string) string {

	buf := []byte(s)
	for i,j:=0,len(s)-1;i<j;{
		if isVowel(s[i]) && isVowel(s[j]){
			buf[i],buf[j] = buf[j],buf[i]
			i++
			j--
		}else if isVowel(s[i]){
			j--
		}else if isVowel(s[j]){
			i++
		}else{
			i++
			j--
		}
	}

	return string(buf)
}

func isVowel(b byte)bool{
	return b == 'a' || b=='A' ||
	 b=='e' || b=='E' ||
	 b=='i' || b=='I' ||
	 b=='o' || b=='O' ||
	 b=='u' || b=='U' 
}
// @lc code=end

