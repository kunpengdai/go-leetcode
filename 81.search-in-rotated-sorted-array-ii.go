/*
 * @lc app=leetcode id=81 lang=golang
 *
 * [81] Search in Rotated Sorted Array II
 *
 * https://leetcode.com/problems/search-in-rotated-sorted-array-ii/description/
 *
 * algorithms
 * Medium (32.72%)
 * Likes:    797
 * Dislikes: 351
 * Total Accepted:    188.9K
 * Total Submissions: 576.8K
 * Testcase Example:  '[2,5,6,0,0,1,2]\n0'
 *
 * Suppose an array sorted in ascending order is rotated at some pivot unknown
 * to you beforehand.
 * 
 * (i.e., [0,0,1,2,2,5,6] might become [2,5,6,0,0,1,2]).
 * 
 * You are given a target value to search. If found in the array return true,
 * otherwise return false.
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [2,5,6,0,0,1,2], target = 0
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [2,5,6,0,0,1,2], target = 3
 * Output: false
 * 
 * Follow up:
 * 
 * 
 * This is a follow up problem to Search in Rotated Sorted Array, where nums
 * may contain duplicates.
 * Would this affect the run-time complexity? How and why?
 * 
 * 
 */
func search(nums []int, target int) bool {
	return helper(nums, target,0,len(nums)-1)
}

func helper(nums []int, target int,beg,end int)bool {
	mid := (beg+end)/2
	if nums[mid] == target {
		return true
	}

	if nums[mid]<target {
		if nums[mid] > nums[end]{

		}else{
			
		}
	}

	if nums[mid]>target {

	}
}

