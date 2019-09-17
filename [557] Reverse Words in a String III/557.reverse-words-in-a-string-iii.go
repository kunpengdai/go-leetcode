/*
 * @lc app=leetcode id=557 lang=golang
 *
 * [557] Reverse Words in a String III
 */
func reverseWords(s string) string {
	ss := strings.Fields(s)
	for i := range ss{
		ss[i] = reverseWord(ss[i])
	}
	return strings.Join(ss," ")
}

func reverseWord(s string)string{
	bytes := []byte(s)
	for i:=0;i<len(bytes)/2;i++{
		bytes[i],bytes[len(s)-i-1] = bytes[len(s)-i-1],bytes[i]
	}
	return string(bytes)
}
