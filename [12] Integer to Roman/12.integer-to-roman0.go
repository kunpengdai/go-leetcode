/*
 * @lc app=leetcode id=12 lang=golang
 *
 * [12] Integer to Roman
 */
func intToRoman(num int) string {
	ints := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	strs := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	res := ""
	for num > 0 {
		for i, im := range ints {
			if num >= im {
				res += strs[i]
				num -= im
				break
			}
		}
	}
	return res
}

