/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原IP地址
 */

// @lc code=start
func restoreIpAddresses(s string) []string {

	var res []string
	for a := 1; a < 4; a++ {
		for b := 1; b < 4; b++ {
			for c := 1; c < 4; c++ {
				for d := 1; d < 4; d++ {

					if a+b+c+d != len(s) {
						continue
					}

					pa, _ := strconv.ParseInt(s[:a], 10, 32)
					pb, _ := strconv.ParseInt(s[a:a+b], 10, 32)
					pc, _ := strconv.ParseInt(s[a+b:a+b+c], 10, 32)
					pd, _ := strconv.ParseInt(s[a+b+c:a+b+c+d], 10, 32)
					if pa < 256 && pb < 256 && pc < 256 && pd < 256 &&
						len(fmt.Sprintf("%d.%d.%d.%d", pa, pb, pc, pd)) == len(s)+3 {
						res = append(res, fmt.Sprintf("%d.%d.%d.%d", pa, pb, pc, pd))
					}
				}
			}
		}
	}

	return res
}

// @lc code=end

