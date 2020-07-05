/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 */

// @lc code=start
func sortedSquares(A []int) []int {
	res := make([]int,len(A))
	left,right,idx:= 0,len(A)-1,len(A)-1
	for left<=right{
		var tmp int
		if abs(A[left])>abs(A[right]){
			tmp = A[left]*A[left]
			left++
		}else{
			tmp = A[right]*A[right]
			right--
		}
		res[idx] = tmp
		idx--
	}
	return res
}

func abs(num int)int{
	if num<0{
		return -num
	}
	return num
}
// @lc code=end
