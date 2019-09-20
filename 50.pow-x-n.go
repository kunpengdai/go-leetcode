/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 *
 * https://leetcode.com/problems/powx-n/description/
 *
 * algorithms
 * Medium (28.32%)
 * Likes:    962
 * Dislikes: 2360
 * Total Accepted:    359.4K
 * Total Submissions: 1.3M
 * Testcase Example:  '2.00000\n10'
 *
 * Implement pow(x, n), which calculates x raised to the power n (x^n).
 * 
 * Example 1:
 * 
 * 
 * Input: 2.00000, 10
 * Output: 1024.00000
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 2.10000, 3
 * Output: 9.26100
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: 2.00000, -2
 * Output: 0.25000
 * Explanation: 2^-2 = 1/2^2 = 1/4 = 0.25
 * 
 * 
 * Note:
 * 
 * 
 * -100.0 < x < 100.0
 * n is a 32-bit signed integer, within the range [−2^31, 2^31 − 1]
 * 
 * 
 */
func myPow(x float64, n int) float64 {
	if n==0{
		return 1.0
	}
	flag :=1
	if n< 0{
		flag=-1
		n=-n
	}
	cnt,nt := 0,n
	for nt>0{
		cnt++
		nt>>=1
	}
	fmt.Println("cnt:",cnt)
	fs := make([]float64, cnt)
	fs[0] = x
	for i := 1; i<len(fs);i++{
		fs[i] = fs[i-1]*fs[i-1]
	}
	res:= float64(1)
	
	for n >0{
		for j := len(fs)-1;j>=0;j--{
			if n >= (1<<uint(j)){
				res *= fs[j]
				n-=(1<<uint32(j))
			} 
		}
	}
	if flag == -1{
		res = 1/res
	}
	return res
}

