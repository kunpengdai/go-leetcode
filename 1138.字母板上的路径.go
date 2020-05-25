/*
 * @lc app=leetcode.cn id=1138 lang=golang
 *
 * [1138] 字母板上的路径
 */

// @lc code=start
func alphabetBoardPath(target string) string {
	var b1 strings.Builder
	var ix,iy byte
	for i := range target{
		tx,ty := (target[i]-'a')/5,(target[i]-'a')%5
		for iy>ty{
			iy--
			b1.Write([]byte{'L'})
		}
		for ix > tx{
			ix--
			b1.Write([]byte{'U'})
		}
		for iy<ty{
			iy++
			b1.Write([]byte{'R'})
		}
		for ix<tx{
			ix++
			b1.Write([]byte{'D'})
		}
		b1.Write([]byte{'!'})
	}  
	return b1.String()
}
// @lc code=end

