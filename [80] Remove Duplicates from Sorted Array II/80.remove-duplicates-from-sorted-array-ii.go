/*
 * @lc app=leetcode id=80 lang=golang
 *
 * [80] Remove Duplicates from Sorted Array II
 */
func removeDuplicates(nums []int) int {
	left,right,exist,len:=0,1,false,len(nums)
	
	for ;right<len;right++{
		if nums[left]==nums[right]{
			if !exist{
				left++
				nums[left],nums[right]=nums[right],nums[left]
				exist=true
			}
			continue
		}
		exist = false
		left++
		nums[left],nums[right]=nums[right],nums[left]
	}
	return left+1
}

