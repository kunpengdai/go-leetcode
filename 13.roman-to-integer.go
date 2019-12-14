/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	rs := []struct{
		roman string
		value int
	}{
		{"M",1000},
		{"CM",900},
		{"D",500},
		{"CD",400},
		{"C",100},
		{"XC",90},
		{"L",50},
		{"XL",40},
		{"X",10},
		{"IX",9},
		{"V",5},
		{"IV",4},
		{"I",1},
	}
	res := 0
	for len(s)>0{
		for i := range rs{
			if len(s)>=len(rs[i].roman) && s[0:len(rs[i].roman)]==rs[i].roman{
				res += rs[i].value
				s=s[len(rs[i].roman):]
			}
		}
	}
	return res
}
// @lc code=end

