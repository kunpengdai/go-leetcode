/*
 * @lc app=leetcode.cn id=1024 lang=golang
 *
 * [1024] 视频拼接
 */

// @lc code=start
func videoStitching(clips [][]int, T int) int {

	cur := 0
	steps := 0
	for cur < T{
		newPos := cur
		steps++ 
		for i:=0;i<len(clips);i++{
			if clips[i][0] <= cur{
				newPos = max(newPos,clips[i][1])
			}   
		}

		if newPos==cur{
			return -1
		} 

		cur = newPos 
	}

	return steps
}

func max(i,j int)int{
	if i>j{
		return i
	}
	return j
}
// @lc code=end

