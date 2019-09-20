/*
 * @lc app=leetcode id=12 lang=golang
 *
 * [12] Integer to Roman
 */
func intToRoman(num int) string {
	is := []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
	ss := []string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}
	res := ""
	for i := range is{
		for num>=is[i]{
			num-=is[i]
			res+=ss[i]
		}
	}
	return res
}

