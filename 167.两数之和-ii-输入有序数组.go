/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	i,j:=0,len(numbers)-1
	for i<j{
		res := numbers[i]+numbers[j]
		if res==target{
			return []int{i+1,j+1}
		}else if res<target{
			i++
		}else{
			j--
		}
	}
	return nil
}
// @lc code=end

