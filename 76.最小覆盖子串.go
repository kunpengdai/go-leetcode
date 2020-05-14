/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	ms := make(map[byte]int)
	mt := make(map[byte]int)

	for i := range t{
		mt[t[i]]++
	}

	var p1 int
	for p1 < len(s){
		if mt[s[p1]]==0{
			p1++
		}else{
			break
		}
	}

	if p1 == len(s){
		return ""
	}

	var res string
	for p2:=p1;p2<len(s);p2++{
		if mt[s[p2]]>0{
			ms[s[p2]]++

			if satisfy(mt,ms) {
				if len(res)>p2-p1+1 || res==""{
					res = s[p1:p2+1]
				}

				np := p1
				for ;np<len(s);np++{

					if mt[s[np]]==0{
						continue
					}

					if ms[s[np]]>mt[s[np]]{
						ms[s[np]]--
					}else{
						break
					}
				}
				p1 = np

				if len(res)>p2+1-p1 || res==""{
					res = s[p1:p2+1]
				}
			}
		}	
	}

	return res
}

func satisfy(mt,ms map[byte]int)bool{

	if len(mt) > len(ms){
		return false
	}

	for k,v := range mt{
		if ms[k]<v{
			return false
		}
	}

	return true
}

// @lc code=end

