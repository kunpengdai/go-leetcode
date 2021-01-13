/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

// @lc code=start
func multiply(num1 string, num2 string) string {

	if num1 == "0" || num2 == "0" {
		return "0"
	}

	res := make([]int, len(num1)+len(num2)-1)
	for i := 0; i < len(num1); i++ {
		for j := 0; j < len(num2); j++ {
			res[i+j] += int(num1[len(num1)-1-i]-'0') * int(num2[len(num2)-1-j]-'0')
		}
	}

	for i := 0; i < len(res)-1; i++ {
		num := res[i]
		res[i] %= 10
		res[i+1] += (num / 10)
	}

	if res[len(res)-1] >= 10 {
		num := res[len(res)-1]
		res[len(res)-1] = num % 10
		res = append(res, num/10)
	}

	bs := make([]byte, len(res))
	for i := 0; i < len(res); i++ {
		bs[len(bs)-1-i] = '0' + byte(res[i])
	}

	return string(bs)
}

// @lc code=end

