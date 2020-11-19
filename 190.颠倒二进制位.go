/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res <<= 1
		if num&(1<<i) > 0 {
			res += 1
		}
	}
	return res
}

// @lc code=end

